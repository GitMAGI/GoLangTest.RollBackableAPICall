package apicallerproxies

/*
L'idea sarebbe di ridefinire per ogni gestore di chiamata API, il metodo Task e cambiare la definizione di Execute in
accordo alla struttura di Task
*/

import "log"

type BaseCall struct {
	baseCall
}

type baseCall struct {
	callID      string
	operationID string
	task        func(args ...interface{}) (interface{}, error)
}

func NewBaseCallProxy(CallID string, OperationID string, Task func(args ...interface{}) (interface{}, error)) *BaseCall {
	m := new(BaseCall)
	m.callID = CallID
	m.operationID = OperationID
	m.task = Task
	return m
}

func (p *baseCall) before() {
	log.Println("Before Starting ...")
	//Scrivere su DB o File che si è iniziata la procedura con IDCall, IDOp. Indicare anche una data
	log.Println("Before Completed")
}

func (p *baseCall) after() {
	log.Println("After Starting ...")
	//Scrivere su DB o File che si è conclusa con successo la procedura con IDCall, IDOp. Indicare anche una data
	log.Println("After Completed")
}

func (p *BaseCall) Execute(args ...interface{}) (interface{}, error) {
	p.before()

	log.Println("Around Starting ...")

	defer p.manageError()
	var input []interface{} = nil
	if args != nil {
		input = args
	}
	data, err := p.task(input)

	log.Println("Around Completed")

	p.after()

	return data, err
}

func (p *baseCall) manageError() {
	err := recover()
	if err != nil {
		log.Println("Errore rilevato")
		//Scrivere su DB o File che si è conclusa con errori la procedura con IDCall, IDOp. Indicare anche una data
		panic(err)
	}
}
