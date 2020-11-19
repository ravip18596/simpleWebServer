package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func addHandlerFunc(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	a,aErr := strconv.Atoi(vars["first"])
	b,bErr := strconv.Atoi(vars["second"])
	if aErr != nil || bErr != nil{
		log.Println(a,aErr,b,bErr)
		w.WriteHeader(http.StatusBadRequest)
		response := make(map[string]string)
		response["error"]="Incorrect input"
		_ = json.NewEncoder(w).Encode(response)
		return
	}
	sum := a+b
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(sumResponse{Sum: sum})
}

func heartbeatHandlerFunc(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_ = json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: http.StatusOK})
	return
}

func handler()  http.Handler{
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/",heartbeatHandlerFunc)
	router.HandleFunc("/add/a/{first}/b/{second}",addHandlerFunc)
	return router
}
//REST Web Server
func main(){
	fmt.Println("Hello World!")
	log.Println("starting http web server started at port 8000")
	err := http.ListenAndServe(":8000",handler())
	if err != nil{
		log.Fatal(err)
	}
}
