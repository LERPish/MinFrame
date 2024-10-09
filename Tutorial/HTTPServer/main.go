/*
As per this tutorial
https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
*/
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	//"os"
)

const keyServerAddress = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")

	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}

	fmt.Printf("%s: got / request. First(%t)=%s, Second(%t)=%s\n, Body:\n%s\n", ctx.Value(keyServerAddress), hasFirst, first, hasSecond, second, body)
	io.WriteString(w, "Fuck Off!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddress))

	myName := r.PostFormValue("myName")
	if myName == "" {
		w.Header().Set("x-missing-field", "myName")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	io.WriteString(w, fmt.Sprintf("No, %s, you Fuck Off!\n", myName))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	ctx := context.Background()
	server := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddress, l.Addr().String())
			return ctx
		},
	}

	//ctx, cancelCtx := context.WithCancel(context.Background())

	/*serverOne := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddress, l.Addr().String())
			return ctx
		},
	}

	serverTwo := &http.Server{
		Addr:    ":4444",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddress, l.Addr().String())
			return ctx
		},
	}

	go func() {
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server 1 Closed\n")
		} else if err != nil {
			fmt.Printf("Error listening for Server 1: %s\n", err)
		}

		cancelCtx()
	}()

	go func() {
		err := serverTwo.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server 2 Closed\n")
		} else if err != nil {
			fmt.Printf("Error listening for Server 2: %s\n", err)
		}
	}()*/

	//<-ctx.Done()

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server Closed\n")
	} else if err != nil {
		fmt.Printf("Error listening for server: %s\n", err)
	}
}
