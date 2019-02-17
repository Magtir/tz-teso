package main

import (
	"net/http"
	"os"
	_ "project/app/config/routes"
	"time"
	"log"
	"project/app/config/routes"
	"fmt"
)

func main() {
	ipServer := "127.0.0.1"
	portServer := "8000"

	if len(os.Args) > 1 {
		ipServer = os.Args[1]
	}

	if len(os.Args) > 1 {
		portServer = os.Args[2]
	}

	fmt.Println("START SERVER " + ipServer + ":" + portServer)

	srv := &http.Server{
		Handler:      routes.InitRoute(),
		Addr:         ipServer + ":" + portServer,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}