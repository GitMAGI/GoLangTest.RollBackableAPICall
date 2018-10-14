package main

import (
	"fmt"
	"log"
	"rollbackapicall/callmanagers"

	"github.com/gogap/aop"
)

const AppName string = "RollBackAPICall"

var callBaseProxy *aop.Proxy

func main() {
	fmt.Println(fmt.Sprintf("%s - Starting", AppName))

	initialize()

	defer errorManager()
	mainContainer()

	//Test01()

	fmt.Println(fmt.Sprintf("%s - Execution Completed", AppName))
}

func initialize() {

	beanFactory := aop.NewClassicBeanFactory()
	beanFactory.RegisterBean("callbase", new(callmanagers.CallBase))
	aspect := aop.NewAspect("aspect_1", "callbase")
	aspect.SetBeanFactory(beanFactory)
	pointcut := aop.NewPointcut("pointcut_1").Execution(`Echo()`)
	//pointcut.Execution(`Echo()`)
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
}

func mainContainer() {
	log.Println("MainContainer Starting")

	//panic(errors.New("Errore fittizio"))
	res := callBaseProxy.Method(new(callmanagers.CallBase).Echo).(func(string) string)("Ciao Mondo")
	fmt.Println(res)

	log.Println("MainContainer Execution Completed")
}

func errorManager() {
	msg := recover()
	if msg != nil {
		log.Println(msg)
	}
}
