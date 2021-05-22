package utils

import (
	"TestScrapeCRUD/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// HandleSuccess ...
func HandleSuccess(resp http.ResponseWriter, status int, data interface{}) {
	responses := models.ResponseWrapper{
		Success: true,
		Message: "Success",
		Data:    data,
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)

	err := json.NewEncoder(resp).Encode(responses)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Oopss, something error"))
		fmt.Printf("[HandleSuccess] error when encode data with error : %v \n", err)
	}
}

//HandleError ...
func HandleError(resp http.ResponseWriter, status int, msg string) {
	responses := models.ResponseWrapper{
		Success: false,
		Message: msg,
		Data:    nil,
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)

	err := json.NewEncoder(resp).Encode(responses)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("ooppss, something error"))
		fmt.Printf("[HandleError] error when encode data with error : %v \n", err)
	}
}
