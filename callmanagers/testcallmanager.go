package callmangers

import (
	"encoding/json"
	"gitmagi/golangtest/rollbackapicall/apicallerproxies"
	"gitmagi/golangtest/rollbackapicall/apiset/product1/resources1"
	"gitmagi/golangtest/rollbackapicall/apiset/product1/resources2"
	"gitmagi/golangtest/rollbackapicall/apiset/product1/resources3"
	"log"
	"net/http"
	"time"

	"github.com/satori/go.uuid"
)

var callID string

var operation1 *apicallerproxies.BaseCall
var operation2 *apicallerproxies.BaseCall
var operation3 *apicallerproxies.BaseCall

func initialize() {
	u2, err := uuid.NewV4()
	if err != nil {
		log.Fatalln(err)
	}
	callID = u2.String()

	operation1 = apicallerproxies.NewBaseCallProxy(callID, "id1", resources1.Test)
	operation2 = apicallerproxies.NewBaseCallProxy(callID, "id2", resources2.Test)
	operation3 = apicallerproxies.NewBaseCallProxy(callID, "id3", resources3.Test)
}

func TestAPIManager(w http.ResponseWriter, r *http.Request) {
	initialize()

	r1, err1 := operation1.Execute(-45, "ciao", time.Now(), .98)
	r2, err2 := operation2.Execute("salve", .98)
	r3, err3 := operation3.Execute("comeva")

	_ = err1
	_ = err2
	_ = err3

	var result map[string]interface{} = make(map[string]interface{})
	result["r1"] = r1
	result["r2"] = r2
	result["r3"] = r3

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
