package callmanagers

import (
	"log"
	"net/http"
)

type FirstCall struct {
	CallBase
}

func (f *FirstCall) CallHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Child Method Call")
}
