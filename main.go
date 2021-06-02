package main

import "fmt"
import "log"
import "net/http"
import "encoding/json"
import "github.com/gorilla/mux"

type Item struct {
  UID	string `json:"UID"`
  Name string `json:"Name"`
  Desc string `json:"Desc"`
  Price float64 `json:"Price"`
}

var investory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called: homePage()")
}

func getInvestories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(investory)
}

func getInvestory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// Remove old item
	item, err := _getItemAtUid(params["uid"])
	if err == true {
		json.NewEncoder(w).Encode(Item{})
	} else {
		json.NewEncoder(w).Encode(item)
	}
}

func createInvestory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	// Adding new item
	investory = append(investory, item)

	json.NewEncoder(w).Encode(item)
}


func updateInvestory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	params := mux.Vars(r)

	// Remove old item
	_deleteItemAtUid(params["uid"])
	// Adding new item
	investory = append(investory, item)

	json.NewEncoder(w).Encode(investory)
}

func deleteInvestory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	// Remove old item
	_deleteItemAtUid(params["uid"])

	json.NewEncoder(w).Encode(investory)
}

func _deleteItemAtUid(uid string) {
	for index, item := range investory {
		if item.UID == uid {
			// Delete item from slice...
			investory = append(investory[:index], investory[index+1:]...)
			break
		}
	}
}

func _getItemAtUid(uid string) (Item, bool) { // todo: We have to use `*Item` and not `Item`
	for _, item := range investory {
		if item.UID == uid {
			return item, false
		}
	}
	return Item{}, true
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/investories", getInvestories).Methods("GET")
	router.HandleFunc("/investory/{uid}", getInvestory).Methods("GET")
	router.HandleFunc("/investory", createInvestory).Methods("POST")
	router.HandleFunc("/investory/{uid}", deleteInvestory).Methods("DELETE")
	router.HandleFunc("/investory/{uid}", updateInvestory).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	investory = append(investory, Item {
		UID: "0",
		Name: "Cheese",
		Desc: "A fine block of cheese",
		Price: 4.99,
	})

	investory = append(investory, Item {
		UID: "1",
		Name: "Milk",
		Desc: "A jug of milk",
		Price: 3.25,
	})

	handleRequests()
}
