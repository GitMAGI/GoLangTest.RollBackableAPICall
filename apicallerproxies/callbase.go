package apicallerproxies

import (
	"fmt"
	"log"
	"time"

	"github.com/gogap/aop"
)

type CallBase struct{}

func (f *CallBase) CallHandler(p1 string, p2 int, p3 time.Time) map[string]interface{} {
	log.Println("Base Method Call")
	log.Println(p1)
	log.Println(p2)
	log.Println(p3)

	var result = make(map[string]interface{})
	result["p1"] = p1
	result["p2"] = p2
	result["p3"] = p3

	return result
}

func (c *CallBase) Before(jp aop.JoinPointer) {
	log.Println("Before Starting ...")

	fmt.Println(jp)

	log.Println("Before Completed")
}

func (c *CallBase) After(p1 string, p2 int, p3 time.Time) {
	log.Println("After Starting ...")
	log.Println(p1)
	log.Println(p2)
	log.Println(p3)
	log.Println("After Completed")
}

func (p *CallBase) Around(pjp aop.ProceedingJoinPointer) {
	log.Println("Around Starting ...")

	defer manageError()

	//log.Println(pjp.Args())

	var p1 string = pjp.Args()[0].(string)
	var p2 int = pjp.Args()[1].(int)
	var p3 time.Time = pjp.Args()[2].(time.Time)

	//log.Println(p1)
	//log.Println(p2)
	//log.Println(p3)

	pjp.Proceed(p1, p2, p3)

	log.Println("Around Completed")
}

func manageError() {
	err := recover()
	if err != nil {
		log.Println("Errore rilevato")
		panic(err)
	}
}
