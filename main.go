package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"request_test/tamrin3/api"
)

func main() {

	_ = NewRESTService()

	if err := http.ListenAndServe("127.0.0.1:8000", nil); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func (data RESTService) handleCalculator(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/calculator" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method Not supported!", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	req := &api.CalculatorRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		panic(err)
	}
	resp, _ := data.history.TaskCalculator(req)

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Fprintf(w, string(jsonResponse))
}

func (data RESTService) handleHistory(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/history" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not supported!", http.StatusNotFound)
		return
	}
	jsonResponse, err := json.Marshal(data.history.GetHistoryResponse())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Fprintf(w, string(jsonResponse))
}

type RESTService struct {
	history *api.CalculatorHistory
}

func NewRESTService() *RESTService{
	service := &RESTService{
		history: &api.CalculatorHistory{},
	}
	http.HandleFunc("/calculator", service.handleCalculator)
	http.HandleFunc("/history", service.handleHistory)
	return service
}
