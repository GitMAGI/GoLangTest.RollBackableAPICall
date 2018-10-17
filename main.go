package main

import (
	"fmt"
	"gitmagi/golangtest/rollbackapicall/callmanagers"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const AppName string = "RollBackAPICall"

const SrvIPAddr string = ""
const SrvPort int = 8000
const SrvIdleTimeout time.Duration = 60
const SrvReadTimeout time.Duration = 15
const SrvWriteTimeout time.Duration = 15

var routeConfig *mux.Router

func main() {
	fmt.Println(fmt.Sprintf("%s - Starting", AppName))

	initialize()

	defer errorManager()
	mainContainer()

	fmt.Println(fmt.Sprintf("%s - Execution Completed", AppName))
}

func initialize() {
	log.Println(AppName, "Routing setting up")

	routeConfig = mux.NewRouter()
	routeConfig.HandleFunc("/test", callmangers.TestAPIManager).Methods("POST")

	log.Println(AppName, "Routes successfully configured")
}

func mainContainer() {
	var addr = SrvIPAddr + ":" + strconv.Itoa(SrvPort)
	log.Println(AppName, "Server booting on", addr)
	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * SrvWriteTimeout,
		ReadTimeout:  time.Second * SrvReadTimeout,
		IdleTimeout:  time.Second * SrvIdleTimeout,
		Handler:      routeConfig,
	}
	log.Println(AppName, "Server is listening at", addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func errorManager() {
	msg := recover()
	if msg != nil {
		log.Println(msg)
	}
}
