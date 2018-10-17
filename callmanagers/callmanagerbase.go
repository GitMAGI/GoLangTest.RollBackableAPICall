package callmanagers

import (
	"net/http"
	"log"

	uuid "github.com/satori/go.uuid"
)

type callManagerBase struct {
	C string
}

type CallManagerBase struct {
	callmanagerbase
}

func NewBaseCallManager() *CallManagerBase {
	m := new(CallManagerBase)
	u2, err := uuid.NewV4()
	if err != nil {
		log.Fatalln(err)
	}
	m.callID = u2.String()
	callID = u2.String()
	return m
}

func (p *callManagerBase)initialize(){

}

func (p *callManagerBase)after(){

}

func (p *callManagerBase)errorManager(){

}

func (p *CallManagerBase)Execute(w http.ResponseWriter, r &http.Request){
	defer p.errorManager()


	p.after()
}	