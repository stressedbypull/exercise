package main

/*
import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Index
func getDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deviceRegistration)
}

//Show
func getDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range deviceRegistration {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//Create
func createDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newDevice DeviceRegistration
	json.NewDecoder(r.Body).Decode(&newDevice)
	newDevice.Id = strconv.Itoa(len(deviceRegistration) + 1)
	deviceRegistration = append(deviceRegistration, newDevice)
	json.NewEncoder(w).Encode(newDevice)
}

//Update
func updateDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range deviceRegistration {
		if item.Id == params["id"] {
			deviceRegistration = append(deviceRegistration[:i], deviceRegistration[i+1:]...)
			var newDevice DeviceRegistration
			json.NewDecoder(r.Body).Decode(&newDevice)
			newDevice.Id = params["id"]
			deviceRegistration = append(deviceRegistration, newDevice)
			json.NewEncoder(w).Encode(newDevice)
			return
		}
	}
}

//Delete
func deleteDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range deviceRegistration {
		if item.Id == params["id"] {
			deviceRegistration = append(deviceRegistration[:i], deviceRegistration[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(deviceRegistration)
}

func main() {

	deviceRegistration = append(deviceRegistration, DeviceRegistration{Id: "msaad", DeviceId: "deviceid", Code: "123456"})

	//inizialize router
	router := mux.NewRouter()

	//endpoints
	router.HandleFunc("/devices", getDevices).Methods("GET") //.Queries("bucket", "{bucket}", "key", "{key}") // percorso e metodo
	router.HandleFunc("/devices/{id}", getDevice).Methods("GET")
	router.HandleFunc("/devices", createDevice).Methods("POST")
	router.HandleFunc("/devices/{id}", updateDevice).Methods(http.MethodPut)
	router.HandleFunc("/devices/{id}", deleteDevice).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":5000", router))
}
*/
