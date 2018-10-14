package callmanagers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gogap/aop"
)

type CallBase struct{}

func (c *CallBase) CallHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Base Method Call")
}

func (c *CallBase) Before(jp aop.JoinPointer) {
	log.Println("Before Starting ...")

	fmt.Println(jp)

	log.Println("Before Completed")
}

func (c *CallBase) After() {
	log.Println("After Starting ...")

	log.Println("After Completed")
}

func (p *CallBase) Around(pjp aop.ProceedingJoinPointer) {
	log.Println("Around Starting ...")

	defer manageError()

	fmt.Println(pjp)

	//ret := pjp.Proceed(pjp.Args()[0], &pjp.Args()[1])
	ret := pjp.Proceed()
	_ = ret

	log.Println("Around Completed")
}

func manageError() {
	err := recover()
	if err != nil {
		log.Println("Errore rilevato")
		panic(err)
	}
}
