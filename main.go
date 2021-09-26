package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/inblack67/go-micro/shortner"
)

func main() {
	router := mux.NewRouter()
	getRouter := router.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", shortner.GetURLs)

	server := &http.Server{
		Handler:      router,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":5000",
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal)

	// signal.Notify(signalChannel, os.Interrupt)
	// signal.Notify(signalChannel, os.Kill)

	signalChannelRes := <-signalChannel

	fmt.Println("Recieved termination request, shutting down: reason =>", signalChannelRes)

	timerContext, cancel := context.WithTimeout(context.Background(), 120*time.Second)

	defer cancel()

	err := server.Shutdown(timerContext)

	if err != nil {
		fmt.Println("Error shutting down")
		log.Fatal(err)
	}

}
