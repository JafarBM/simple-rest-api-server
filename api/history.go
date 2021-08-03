package api

import (
	"sort"
	"sync"
)

type CalculatorHistory struct {
	Size  int
	Items []*historyItem
	mutex sync.Mutex
}

type historyItem struct {
	Task    string      `json:"task"`
	Numbers []int       `json:"numbers"`
	Answer  interface{} `json:"answer"`
}

type HistoryResponse struct {
	Size    int            `json:"size"`
	History []*historyItem `json:"history"`
	Code    int            `json:"code"`
	Message string         `json:"message"`
}

type CalculatorResponse struct {
	Task    string      `json:"task"`
	Numbers []int       `json:"numbers"`
	Answer  interface{} `json:"answer"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

type CalculatorRequest struct {
	Task    string `json:"task"`
	Numbers []int  `json:"numbers"`
}

func (h *CalculatorHistory) GetHistoryResponse() *HistoryResponse {
	return &HistoryResponse{
		Size:    h.Size,
		History: h.Items,
		Code:    200,
		Message: "History sent successfully!‬",
	}
}

func (h *CalculatorHistory) addToHistory(calculatorResponse *CalculatorResponse) {
	h.Items = append(h.Items, &historyItem{
		Task:    calculatorResponse.Task,
		Numbers: calculatorResponse.Numbers,
		Answer:  calculatorResponse.Answer,
	})
	h.Size += 1
}

func (h *CalculatorHistory) TaskCalculator(req *CalculatorRequest) (*CalculatorResponse, error) {
	response := &CalculatorResponse{
		Numbers: req.Numbers,
		Task:    req.Task,
	}
	if len(req.Numbers) == 0 {
		response.Code = 400
		response.Message = "Provide Numbers."
	}

	if req.Task == "mean" {
		response.Answer = calculateMean(req.Numbers)
		response.Code = 200
		response.Message = "Task done successfully!"
	} else if req.Task == "sort" {
		response.Answer = calculateSort(req.Numbers)
		response.Code = 200
		response.Message = "Task done successfully!"
	} else {
		response.Code = 400
		response.Message = "Invalid Task."
	}

	if response.Code == 200 {
		h.mutex.Lock()
		h.addToHistory(response)
		h.mutex.Unlock()
	}

	return response, nil
}

func calculateMean(numbers []int) float32 {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return float32(sum) / float32(len(numbers))
}

func calculateSort(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}
