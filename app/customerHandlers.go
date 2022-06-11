package app

import (
	"banking/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	Service service.CustomerService
}

func NewCustomerHandlers(service service.CustomerService) *CustomerHandlers {
	return &CustomerHandlers{Service: service}
}

func (handler *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := handler.Service.GetAllCustomer()
	writeResponse(w, http.StatusOK, customers)
}

func (handler *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer, err := handler.Service.GetCustomer(vars["customer_id"])
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "post request received")
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}

}
