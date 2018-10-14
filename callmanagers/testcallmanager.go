package callmangers

import (
	"encoding/json"
	"fmt"
	"gitmagi/golangtest/rollbackapicall/apicallerproxies"
	"log"
	"net/http"
	"time"

	"github.com/gogap/aop"
)

var callBaseProxy *aop.Proxy

func initialize() {
	beanFactory := aop.NewClassicBeanFactory()
	beanFactory.RegisterBean("callbase", new(apicallerproxies.CallBase))
	aspect := aop.NewAspect("aspect_1", "callbase")
	aspect.SetBeanFactory(beanFactory)
	pointcut := aop.NewPointcut("pointcut_1").Execution(`.*?`)
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

func TestAPIManager(w http.ResponseWriter, r *http.Request) {
	initialize()

	result := callBaseProxy.Method(new(apicallerproxies.CallBase).CallHandler).(func(string, int, time.Time) map[string]interface{})("test", -23, time.Now())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func Test() {
	beanFactory := aop.NewClassicBeanFactory()
	beanFactory.RegisterBean("callbase", new(apicallerproxies.CallBase))
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
	callProxy, err := gogapAop.GetProxy("callbase")

	if err != nil {
		log.Fatalln(err)
	}

	//result := callProxy.Method(new(apicallerproxies.CallBase).CallHandler).(func(string, int, time.Time) map[string]interface{})("test", -23, time.Now())
	var result map[string]interface{}
	callProxy.Invoke(new(apicallerproxies.CallBase).CallHandler, "test", -23, time.Now()).End(
		func(data map[string]interface{}) {
			result = data
		})

	fmt.Println(result)
}
