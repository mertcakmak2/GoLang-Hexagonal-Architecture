package app

import (
	"banking/domain"
	"banking/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	customerRepository := domain.NewCustomerRepositoryDb()
	customerService := service.NewCustomerService(customerRepository)
	// customerHandler := CustomerHandlers{Service: customerService}

	customerHandlerV2 := NewCustomerHandlers(customerService)

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandlerV2.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id}", customerHandlerV2.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
