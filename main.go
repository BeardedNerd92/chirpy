package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	const filePath = "."

	apiCfg := apiConfig{
		fileServerHits: 0,
	}

	mux := http.NewServeMux()

	corsMux := middlewareCors(mux)

	mux.Handle("/app/", apiCfg.middlerwareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filePath)))))
	mux.HandleFunc("/healthz", handlerRediness)
	mux.HandleFunc("/metrics", apiCfg.handlerMetrics)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filePath, port)
	log.Fatal(srv.ListenAndServe())
}
