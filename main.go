package main

import (
	"fmt"
	"gitmagi/golangtest/rollbackapicall/apicallerproxies"
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

func main() {
	fmt.Println(fmt.Sprintf("%s - Starting", AppName))

	p := func(args ...interface{}) (interface{}, error) {
		var data []interface{} = args[0].([]interface{})
		var p1 int = data[0].(int)
		var p2 string = data[1].(string)
		var p3 time.Time = data[2].(time.Time)

		fmt.Println(p1)
		fmt.Println(p2)
		fmt.Println(p3)

		return "ciao", nil
	}

	var a = apicallerproxies.NewTemplateCallProxy("2323", "132312", p)
	d, err := a.Execute(-23, "strada", time.Now())
	_ = d
	_ = err

	/*
		initialize()

		defer errorManager()
		mainContainer()

	*/
	fmt.Println(fmt.Sprintf("%s - Execution Completed", AppName))
}

func initialize() {

}

func mainContainer() {
	log.Println("MainContainer Starting")

	log.Println(AppName, "Routing setting up")
	r := mux.NewRouter()
	r.HandleFunc("/test", callmangers.TestAPIManager).Methods("POST")
	log.Println(AppName, "Routes successfully configured")

	var addr = SrvIPAddr + ":" + strconv.Itoa(SrvPort)
	log.Println(AppName, "Server booting on", addr)
	srv := &http.Server{
		Addr: addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * SrvWriteTimeout,
		ReadTimeout:  time.Second * SrvReadTimeout,
		IdleTimeout:  time.Second * SrvIdleTimeout,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}
	log.Println(AppName, "Server is listening at", addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}

	log.Println("MainContainer Execution Completed")
}

func errorManager() {
	msg := recover()
	if msg != nil {
		log.Println(msg)
	}
}
