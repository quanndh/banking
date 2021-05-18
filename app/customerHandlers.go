package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/quannguyennn/banking/service"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (handler *CustomerHandler) getAllCustomers(res http.ResponseWriter, req *http.Request) {
	status := req.URL.Query().Get("status")

	customers, err :=  handler.service.GetAllCustomer(status)

	if err != nil {
		writeResponse(res, err.Code, err.AsMessage())
	} else {
		writeResponse(res, http.StatusOK, customers)
	}
}

func (handler *CustomerHandler) getCustomerDetail(res http.ResponseWriter, req *http.Request)  {
	vars := mux.Vars(req)

	customer, err := handler.service.GetCustomerDetail(vars["customer_id"])

	if err != nil {
		writeResponse(res, err.Code, err.AsMessage())
	} else {
		writeResponse(res, http.StatusOK, customer)

	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

