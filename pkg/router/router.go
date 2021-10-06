package router

import (
	"context"
	"net/http"
	"restapi/pkg/controller"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)

var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "Status of HTTP response",
	},
	[]string{"status"},
)

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_response_time_seconds",
	Help: "Duration of HTTP requests. Custom",
}, []string{"path"})

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		rw := newResponseWriter(w)
		next.ServeHTTP(rw, r)

		statusCode := rw.statusCode

		responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
		totalRequests.WithLabelValues(path).Inc()

		timer.ObserveDuration()
	})
}

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})
	hdFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hd_errors_total",
			Help: "Number of hard-disk errors.",
		},
		[]string{"device"},
	)
)

func init() {
	prometheus.Register(totalRequests)
	prometheus.Register(responseStatus)
	prometheus.Register(httpDuration)
	prometheus.MustRegister(cpuTemp)
}

//Router for all requests
func Router(ctx context.Context) (err error) {

	cpuTemp.Set(65.3)
	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()

	// Initialization mux router
	r := mux.NewRouter()

	// Prometheus Middleware
	r.Use(prometheusMiddleware)

	// Prometheus endpoint
	r.Path("/prometheus").Handler(promhttp.Handler())

	//Events requests
	r.Handle("/api/events/filter", controller.IsAuthorized(controller.FilterEvents)).Methods("GET")
	r.Handle("/api/events/filter/between", controller.IsAuthorized(controller.BetweenFilterEvents)).Methods("GET")
	r.Handle("/api/events", controller.IsAuthorized(controller.GetEvents)).Methods("GET")
	r.Handle("/api/events/{id}", controller.IsAuthorized(controller.GetEvent)).Methods("GET")
	r.Handle("/api/events", controller.IsAuthorized(controller.CreateEvent)).Methods("POST")
	r.Handle("/api/events/{id}", controller.IsAuthorized(controller.UpdateEvent)).Methods("PUT")
	r.Handle("/api/events/{id}", controller.IsAuthorized(controller.DeleteEvent)).Methods("DELETE")

	//User requests
	r.Handle("/api/users", controller.IsAuthorized(controller.UpdateUser)).Methods("PUT")

	//Authorization requests
	r.HandleFunc("/api/login", controller.Login).Methods("POST")
	r.Handle("/api/logout", controller.IsAuthorized(controller.Logout)).Methods("POST")

	srv := &http.Server{
		Addr:    ":9009",
		Handler: r,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//log.Fatalf("listen:%+s\n", err)
			log.Fatal().Err(err).Msg("listen:%+s\n")
		}
	}()

	log.Info().Msg("server started")
	<-ctx.Done()

	log.Info().Msg("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatal().Err(err).Msg("server Shutdown Failed:%+s")
	}

	log.Info().Msg("server exited properly")
	if err == http.ErrServerClosed {
		err = nil
	}

	return

}
