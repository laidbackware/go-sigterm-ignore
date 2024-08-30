// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Basic web app which will listen for a SIGTERM and ignore it
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Create channel to listen for signals.
var signalChan chan (os.Signal) = make(chan os.Signal, 1)

func main() {
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: http.HandlerFunc(handler),
	}

	// SIGINT handles Ctrl+C locally.
	// SIGTERM handles Cloud Run termination signal.
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Start HTTP server.
	go func() {
		log.Printf("listening on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Receive output from signalChan.
	sig := <-signalChan
	log.Printf("%s signal caught", sig)

	// Infinite loop to ensure SIGKILL is triggered
	for {
		log.Print("still alive!")
		time.Sleep(1 * time.Second) 
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!\n")
}