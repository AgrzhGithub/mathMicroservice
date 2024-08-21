package main

import (
	"github.com/gorilla/mux"
	"log"
	"mathMicroservice/pkg"
	"net/http"
)

var RobotHandler = pkg.Handler
var dotsDistanceHandler = pkg.DotsHandler

func main() {

	r := mux.NewRouter()

	//robot
	//r.HandleFunc("/handler/{mass:[0-9]+}/{velocity:[0-9]+}/", RobotHandler)

	//2dotsDistance
	r.HandleFunc("/handler/{x1:[0-9]+}{y1:[0-9]+}/{x2:[0-9]+}{y2:[0-9]+}", dotsDistanceHandler)

	log.Fatal(http.ListenAndServe(":8080", r))

}
