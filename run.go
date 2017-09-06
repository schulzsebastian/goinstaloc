package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/schulzsebastian/goinstaloc/api"
)

func runWebpack() {
	cmd := exec.Command("sh", "-c", "cd static && npm run dev")
	_ = cmd.Start()
	log.Print("Running Webpack on port :5001")
}

func startHTTPServer(prod bool) {
	router := mux.NewRouter()
	router.HandleFunc("/websocket", api.WebsocketHandler)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	if prod == true {
		log.Print("NPM building...")
		cmd := exec.Command("sh", "-c", "cd static && npm run build")
		_ = cmd.Start()
		cmd.Wait()
	}
	log.Print("Running HTTP server on port :5000")
	http.ListenAndServe(":5000", router)
}

func listenStop(close chan bool, mode bool) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for _ = range c {
		log.Print("Shutdowning HTTP server on port :5000")
		if mode == false {
			log.Print("Shutdowning Webpack on port :5001")
		}
		close <- true
	}
}

func main() {
	// Production mode
	prod := false
	// Listen for CTRL+C
	close := make(chan bool)
	go listenStop(close, prod)
	// Run prod or dev mode
	if prod == true {
		go startHTTPServer(true)
	} else {
		go startHTTPServer(false)
		go runWebpack()
	}
	// Wait for value from listenStop
	<-close
}
