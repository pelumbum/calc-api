package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/pelumbum/calc-api/calculator"
)


type requestData struct {
	Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	requestData := requestData{}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "invalid JSON", http.StatusUnprocessableEntity)
		return
	}
	expression := requestData.Expression
	result, err := calculator.Calc(expression)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.Write([]byte(fmt.Sprintf(`{"result":%f}`, result)))
}

func main() {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("failed to start server: %v\n", err)
	}
}
