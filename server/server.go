package server

import (
    "log"
    "net/http"
)

//Starts an HTTP/HTTPS server
func ListenAndServe(addr, certfile, keyFile string, usesTls bool, handler http.Handler) {
    if usesTls == true {
        log.Printf("Starting HTTP server on %s\n", addr)
        err := http.ListenAndServe(addr, handler)
        if err != nil {
            log.Fatalf("HTTP server failed %s", err)
        }
    } else {
        log.Printf("Starting HTTPS server on %s\n", addr)
        err := http.ListenAndServeTLS(addr, certfile, keyFile, handler)
        if err != nil {
            log.Fatalf("HTTPS server failed %s", err)
        }
    }
}
