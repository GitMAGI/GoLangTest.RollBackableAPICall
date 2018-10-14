package apicallerproxies

import (
	"log"
	"time"
)

type FirstCall struct {
	CallBase
}

func (f *FirstCall) CallHandler(p1 string, p2 int, p3 time.Time) map[string]interface{} {
	log.Println("Child Method Call")
	log.Println(p1)
	log.Println(p2)
	log.Println(p3)

	var result = make(map[string]interface{})
	result["p1"] = p1
	result["p2"] = p2
	result["p3"] = p3

	return result
}
