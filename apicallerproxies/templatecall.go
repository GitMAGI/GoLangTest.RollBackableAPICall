package apicallerproxies

/*
L'idea sarebbe di ridefinire per ogni gestore di chiamata API, il metodo Task e cambiare la definizione di Execute in
accordo alla struttura di Task
*/

import "log"

type templateCallProxy struct {
	callID      string
	operationID string
	task        func(args ...interface{}) (interface{}, error)
}

func NewTemplateCallProxy(CallID, OperationID string, Task func(args ...interface{}) (interface{}, error)) *templateCallProxy {
	m := new(templateCallProxy)
	m.callID = CallID
	m.operationID = OperationID
	m.task = Task
	return m
}

func (p *templateCallProxy) before() {
	log.Println("Before Starting ...")
	//Scrivere su DB o File che si è iniziata la procedura con IDCall, IDOp. Indicare anche una data
	log.Println("Before Completed")
}

func (p *templateCallProxy) after() {
	log.Println("After Starting ...")
	//Scrivere su DB o File che si è conclusa con successo la procedura con IDCall, IDOp. Indicare anche una data
	log.Println("After Completed")
}

func (p *templateCallProxy) Execute(args ...interface{}) (interface{}, error) {
	log.Println("Around Starting ...")

	defer p.manageError()
	var input []interface{} = nil
	if args != nil {
		input = args
	}
	data, err := p.task(input)

	log.Println("Around Completed")

	return data, err
}

func (p *templateCallProxy) manageError() {
	err := recover()
	if err != nil {
		log.Println("Errore rilevato")
		//Scrivere su DB o File che si è conclusa con errori la procedura con IDCall, IDOp. Indicare anche una data
		panic(err)
	}
}
