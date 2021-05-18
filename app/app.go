package app

import (
	"fmt"
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

	host := fmt.Sprintf("%s:%s", "localhost", "3000")

	log.Fatal(http.ListenAndServe(host, router))
}
