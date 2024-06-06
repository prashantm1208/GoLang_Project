package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"assesment.com/client"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
)

func PostCustomerDetailHandler(w http.ResponseWriter, r *http.Request) {
	var details client.CustomerDetails

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&details)
	if err != nil {
		log.Println("Error occured while decoding customer details :", err)
	}

	err2 := client.InsertCustomer(details)
	if err != nil {
		log.Fatal("Error occured while inserting customer details :", err2)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func GetCustomerHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	custid := params["customid"]
	customID, err := strconv.Atoi(custid)
	if err != nil {
		log.Fatal("Error occured while converting string to int :", err)
	}
	customer, err2 := client.GetCustomerByID(customID)
	if err2 != nil {
		log.Fatal("Error while fetching record: ", err2)
	}

	json.NewEncoder(w).Encode(customer)
}

func UpdateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var emp client.CustomerDetails
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	emp.ID = id
	if err := client.UpdateCustomer(emp); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(emp)

}

func DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	custid := params["customid"]
	customID, err := strconv.Atoi(custid)
	if err != nil {
		log.Fatal("Error occured while converting string to int :", err)
	}

	err2 := client.DeleteCustomerByID(customID)
	if err2 != nil {
		w.Write([]byte(err2.Error()))
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Paging
func ListCustomerHandler(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	options := options.Find()
	options.SetSkip(int64((page - 1) * pageSize))
	options.SetLimit(int64(pageSize))
	session := client.Connect()

	collection := session.Database("Customer").Collection("details")
	cur, err := collection.Find(ctx, bson.D{}, options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(ctx)

	var employees []client.CustomerDetails
	if err := cur.All(ctx, &employees); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(employees)
}
