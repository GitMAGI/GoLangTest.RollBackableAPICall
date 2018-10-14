package main

import (
	"fmt"
	"gitmagi/golangtest/rollbackapicall/callmanagers"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gogap/aop"
	"github.com/gorilla/mux"
)

const AppName string = "RollBackAPICall"

const SrvIPAddr string = ""
const SrvPort int = 8000
const SrvIdleTimeout time.Duration = 60
const SrvReadTimeout time.Duration = 15
const SrvWriteTimeout time.Duration = 15

var callBaseProxy *aop.Proxy
var firstCallProxy *aop.Proxy

func main() {
	fmt.Println(fmt.Sprintf("%s - Starting", AppName))

	initialize()

	defer errorManager()
	mainContainer()

	fmt.Println(fmt.Sprintf("%s - Execution Completed", AppName))
}

func initialize() {

	/*
		beanFactory := aop.NewClassicBeanFactory()
		beanFactory.RegisterBean("firstcall", new(callmanagers.FirstCall))
		aspect := aop.NewAspect("aspect_1", "firstcall")
		aspect.SetBeanFactory(beanFactory)
		pointcut := aop.NewPointcut("pointcut_1").Execution(`CallHandler()`)
		aspect.AddPointcut(pointcut)
		aspect.AddAdvice(&aop.Advice{Ordering: aop.Before, Method: "Before", PointcutRefID: "pointcut_1"})
		aspect.AddAdvice(&aop.Advice{Ordering: aop.After, Method: "After", PointcutRefID: "pointcut_1"})
		aspect.AddAdvice(&aop.Advice{Ordering: aop.Around, Method: "Around", PointcutRefID: "pointcut_1"})
		gogapAop := aop.NewAOP()
		gogapAop.SetBeanFactory(beanFactory)
		gogapAop.AddAspect(aspect)
		var err error
		firstCallProxy, err = gogapAop.GetProxy("firstcall")
	*/

	beanFactory := aop.NewClassicBeanFactory()
	beanFactory.RegisterBean("callbase", new(callmanagers.CallBase))
	aspect := aop.NewAspect("aspect_1", "callbase")
	aspect.SetBeanFactory(beanFactory)
	pointcut := aop.NewPointcut("pointcut_1").Execution(`CallHandler()`)
	aspect.AddPointcut(pointcut)
	aspect.AddAdvice(&aop.Advice{Ordering: aop.Before, Method: "Before", PointcutRefID: "pointcut_1"})
	aspect.AddAdvice(&aop.Advice{Ordering: aop.After, Method: "After", PointcutRefID: "pointcut_1"})
	aspect.AddAdvice(&aop.Advice{Ordering: aop.Around, Method: "Around", PointcutRefID: "pointcut_1"})
	gogapAop := aop.NewAOP()
	gogapAop.SetBeanFactory(beanFactory)
	gogapAop.AddAspect(aspect)
	var err error
	callBaseProxy, err = gogapAop.GetProxy("callbase")

	if err != nil {
		log.Fatalln(err)
	}

	_ = callBaseProxy
	_ = firstCallProxy
}

func mainContainer() {
	log.Println("MainContainer Starting")

	log.Println(AppName, "Routing setting up")
	r := mux.NewRouter()
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		//callBaseProxy.Method(new(callmanagers.CallBase).CallHandler).(func(http.ResponseWriter, *http.Request))(w, r)
		callBaseProxy.Invoke(new(callmanagers.CallBase).CallHandler, w, r)
	}).Methods("POST")
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
