package callmanagers

import (
	"errors"
	"log"

	"github.com/gogap/aop"
)

type CallBase struct{}

func (c *CallBase) Echo(var1 string) string {
	//fmt.Println(" >>>", var1)

	panic(errors.New("Errore interno"))

	return var1
}

func (c *CallBase) Before(jp aop.JoinPointer) {
	log.Println("Before Starting ...")

	//fmt.Println(jp)

	log.Println("Before Completed")
}

func (c *CallBase) After() {
	log.Println("After Starting ...")

	log.Println("After Completed")
}

func (p *CallBase) Around(pjp aop.ProceedingJoinPointer) {
	log.Println("Around Starting ...")

	defer manageError()

	ret := pjp.Proceed(pjp.Args()[0])
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
