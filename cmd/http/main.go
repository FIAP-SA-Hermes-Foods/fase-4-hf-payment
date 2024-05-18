package main

import (
	l "fase-4-hf-payment/external/logger"
	httpH "fase-4-hf-payment/internal/handler/http"
	"log"
	"net/http"
	"os"

	"github.com/marcos-dev88/genv"
)

func init() {
	if err := genv.New(); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}
}

func main() {

	router := http.NewServeMux()

	h := httpH.NewHandler()

	router.Handle("/hermes_foods/health", http.StripPrefix("/", httpH.Middleware(h.HealthCheck)))
	router.Handle("/hermes_foods/mercado_pago_api/", http.StripPrefix("/", httpH.Middleware(h.Handler)))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_HTTP_PORT"), router))
}
