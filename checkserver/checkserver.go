package main

import (
	"flag"
	"fmt"
	"log"

	//"io/ioutil"
	"net/http"

	//"encoding/json"
)

type checkDefinition struct {
	url string
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	method := req.Method // Get HTTP Method (string)
	req.ParseForm()      // Populates request.Form
	values := req.Form

	// shell> curl -H 'Content-Type: application/json' \
	//            -X POST http://127.0.0.1:5984/demo \
	//            -d '{"company": "Example, Inc."}'

	// decoder := json.NewDecoder(r.Body)
	// var check checkDefinition
	// err := decoder.Decode(&check)
	// if err != nil {
	// 	fmt.Println("ERROR", err)
	// 	panic(err)
	// }
	// log.Println("check", check)

	fmt.Println("EFA values", values)
	fmt.Println("EFA method", method)

	fmt.Fprint(w, "LALALA")
}

func main() {
	address := flag.String("address", "0.0.0.0", "listen address")
	port := flag.String("port", "8080", "listen port")
	flag.Parse()

	http.HandleFunc("/", rootHandler)

	addressAndPort := fmt.Sprintf("%s:%s", *address, *port)
	log.Println("Server running in", addressAndPort, " ...")
	log.Fatal(http.ListenAndServe(addressAndPort, nil))
}
