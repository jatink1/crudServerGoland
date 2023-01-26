package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main(){

	//this is to ensure .env is loaded properly, if it isn't then it show throw an error as db connectivity won't be possible
	if err:=godotenv.Load(); err!=nil{
		log.Fatalln(err)
	}
	fmt.Println(uuid.New())

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/products",createProduct).Methods("POST")

	// db := connect()
	// defer db.Close()

	// product :=  Product{ID: uuid.New(), Name: "My Product", Quantity: 10, Price: 45}
	// fmt.Println(product)

	// Starting Server
	// if err := http.ListenAndServe(":8080",r); err!=nil{
	// 	log.Fatalln(err)
	// 	fmt.Println("Error")
	// }else{
	// 	fmt.Println("Connected to server on port 8080")
	// }
	log.Fatalln(http.ListenAndServe(":8080",r))
}

func createProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")

	// Get db connection
	db := connect()
	defer db.Close()

	// Create Product Instance
	product := &Product{
		ID: uuid.New().String(),
	}

	_ = json.NewDecoder(r.Body).Decode(&product)

	// Inserting into database
	_, err := db.Model(product).Insert()
	if err!= nil{
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// Returning Product
	json.NewEncoder(w).Encode(product)
}