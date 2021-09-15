package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"C"
	"github.com/google/uuid"
)

type ReturnValue struct {
	Data string
}
type InvokeResponse struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue interface{}
}

type InvokeResponseStringReturnValue struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue string
}

type InvokeRequest struct {
	Data     map[string]interface{}
	Metadata map[string]interface{}
}



func genUUID() *C.char {

	id := uuid.New()
	return C.CString(id.String())
}


func simpleHttpTriggerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		w.Write([]byte(C.GoString((*C.char)(genUUID()))))
	}
}

func main() {
	customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if !exists {
		customHandlerPort = "8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/SimpleHttpTriggerWithReturn", simpleHttpTriggerHandler)
	fmt.Println("Go server Listening...on FUNCTIONS_CUSTOMHANDLER_PORT:", customHandlerPort)
	log.Fatal(http.ListenAndServe(":"+customHandlerPort, mux))
}
