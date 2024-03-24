package main

import (
	"fmt"
	"net"
	"net/http"
)

func validateMethod(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		handler(w, r)
	}
}

func checkStatusHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Dota is the best")
}

func bestSkinHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "PLEASE BUY ME AN ARCANA")
}

func dotaShopHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Thank you for shopping at dota")
}

func dotaPut(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "putting something brok")
}

func main() {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/dota/store", validateMethod(http.MethodGet, checkStatusHandler))

	serveMux.HandleFunc("/dota/skin", validateMethod(http.MethodGet, bestSkinHandler))

	serveMux.HandleFunc("/dota/buy-skin", validateMethod(http.MethodPost, dotaShopHandler))

	serveMux.HandleFunc("/dota/put-skin", validateMethod(http.MethodPut, dotaPut))

	httpServer := &http.Server{
		Addr:    "localhost:4321",
		Handler: serveMux,
	}

	listener, err := net.Listen("tcp", httpServer.Addr)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	err = httpServer.Serve(listener)

	if err != http.ErrServerClosed {
		fmt.Println(err)
		return
	}

}
