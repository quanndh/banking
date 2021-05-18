package app

import (
	"github.com/gorilla/mux"
	"github.com/quannguyennn/banking/domain"
	"github.com/quannguyennn/banking/service"
	"log"
	"net/http"
)

func Start() {
	//mux := http.NewServeMux()

	router := mux.NewRouter()

	//wiring
	customerHandler := CustomerHandler{service : service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", customerHandler.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.getCustomerDetail).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
