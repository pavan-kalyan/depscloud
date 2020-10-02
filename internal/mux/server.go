package mux

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-prometheus"

	"github.com/mjpitz/go-gracefully/check"
	"github.com/mjpitz/go-gracefully/health"
	"github.com/mjpitz/go-gracefully/state"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/rs/cors"

	"github.com/sirupsen/logrus"

	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	std "github.com/slok/go-http-metrics/middleware/std"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"google.golang.org/grpc"
	grpchealth "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Version struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

type Config struct {
	context.Context

	BindAddressHTTP string
	BindAddressGRPC string

	Checks []check.Check

	TLSConfig *TLSConfig

	Version *Version
}

func DefaultServers() (*grpc.Server, *http.ServeMux) {
	grpcOpts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_prometheus.StreamServerInterceptor,
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
		)),
	}

	grpc_prometheus.EnableHandlingTimeHistogram()

	return grpc.NewServer(grpcOpts...), http.NewServeMux()
}

func registerHealth(grpcServer *grpc.Server, httpServer *http.ServeMux, config *Config) {
	monitor := health.NewMonitor(config.Checks...)
	reports, unsubscribe := monitor.Subscribe()
	stopCh := config.Context.Done()

	healthCheck := grpchealth.NewServer()

	go func() {
		defer unsubscribe()

		for {
			select {
			case <-stopCh:
				return
			case report := <-reports:
				if report.Check == nil {
					if report.Result.State == state.Outage {
						healthCheck.SetServingStatus("", healthpb.HealthCheckResponse_NOT_SERVING)
					} else {
						healthCheck.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
					}
				}
			}
		}
	}()

	handler := health.HandlerFunc(monitor)
	httpServer.HandleFunc("/healthz", handler)
	httpServer.HandleFunc("/health", handler)

	healthpb.RegisterHealthServer(grpcServer, healthCheck)
	_ = monitor.Start(config.Context)
}

func registerMetrics(httpServer *http.ServeMux) {
	httpServer.Handle("/metrics", promhttp.Handler())
}

func registerVersion(httpServer *http.ServeMux, config *Config) {
	httpServer.HandleFunc("/version", func(writer http.ResponseWriter, request *http.Request) {
		version, err := json.Marshal(config.Version)
		if err == nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write(version)
		} else {
			logrus.Error(err.Error())
			http.Error(writer, "Something went a bit wrong here!", 500)
		}

	})
}

func monitorHandler(httpServer http.Handler) http.Handler {
	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})
	return std.Handler("", mdlw, httpServer)
}

func Serve(grpcServer *grpc.Server, httpServer http.Handler, config *Config) error {
	// TODO setup proper shutdown handlers

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.ProtoMajor == 2 &&
			strings.HasPrefix(request.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(writer, request)
		} else {
			httpServer.ServeHTTP(writer, request)
		}
	})

	reflection.Register(grpcServer)
	registerHealth(grpcServer, httpMux, config)
	registerMetrics(httpMux)
	registerVersion(httpMux, config)

	grpc_prometheus.Register(grpcServer)

	corsMux := cors.Default().Handler(httpMux)
	h2cMux := h2c.NewHandler(corsMux, &http2.Server{})

	var grpcListener net.Listener
	var grpcErr error

	var httpListener net.Listener
	var httpErr error

	tlsConfig, err := LoadTLSConfig(config.TLSConfig)
	if err != nil {
		return err
	}

	if tlsConfig != nil {
		httpListener, httpErr = tls.Listen("tcp", config.BindAddressHTTP, tlsConfig)
		grpcListener, grpcErr = tls.Listen("tcp", config.BindAddressGRPC, tlsConfig)
	} else {
		httpListener, httpErr = net.Listen("tcp", config.BindAddressHTTP)
		grpcListener, grpcErr = net.Listen("tcp", config.BindAddressGRPC)
	}

	if httpErr != nil {
		return httpErr
	}
	if grpcErr != nil {
		return grpcErr
	}

	logrus.Infof("[runtime] starting http on %s", config.BindAddressHTTP)
	go http.Serve(httpListener, monitorHandler(h2cMux))

	logrus.Infof("[runtime] starting grpc on %s", config.BindAddressGRPC)
	return grpcServer.Serve(grpcListener)
}
