package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Offices struct {
	Name     string `json:"Name"`
	Location string `json:"Location"`
}
type officelist []Offices

type OfficeResponse struct {
	Offices []Office `json:"offices"`
}

type Office struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Floor    string `json:"floor"`
	Capacity int    `json:"capacity"`
}

type SingleOffice []Office

func main() {
	log.Println("Starting API server")

	router := mux.NewRouter()

	router.HandleFunc("/health_check", HealthCheck).Methods("GET")
	router.HandleFunc("/offices", OfficesEndpoint).Methods("GET")
	router.HandleFunc("/offices/{Name}", getOneOffice).Methods("GET")
	router.HandleFunc("/", homeLink)

	http.ListenAndServe(":8080", router)
}

var officesList = officelist{
	{
		Name:     "Headquarters",
		Location: "City Center",
	},
	{
		Name:     "Branch A",
		Location: "Suburb Area",
	},
	{
		Name:     "Branch B",
		Location: "Downtown",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Welcome to my page! I work in different offices. Here is the list of my offices.")
	json.NewEncoder(w).Encode(officesList)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering health check endpoint")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Description: This app provides information about offices. Welcome to my app! ;)\nAuthor: Merey")
}

func getOneOffice(w http.ResponseWriter, r *http.Request) {
	officeName := mux.Vars(r)["Name"]

	for _, singleOffice := range offices {
		if singleOffice.Name == officeName {
			json.NewEncoder(w).Encode(singleOffice)
		}
	}
}

func OfficesEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering offices endpoint")
	var response OfficeResponse
	offices := prepareOfficeResponse()

	response.Offices = offices

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

var offices = SingleOffice{
	{
		Id:       1,
		Name:     "Miras",
		Location: "Miras 25",
		Floor:    "3nd",
		Capacity: 500,
	},
	{
		Id:       2,
		Name:     "Branch A",
		Location: "Suburb Area",
		Floor:    "2nd",
		Capacity: 150,
	},
	{
		Id:       3,
		Name:     "Branch B",
		Location: "Downtown",
		Floor:    "5th",
		Capacity: 200,
	},
}

func prepareOfficeResponse() []Office {
	var offices []Office

	var office Office
	office.Id = 1
	office.Name = "Headquarters"
	office.Location = "City Center"
	office.Floor = "10th"
	office.Capacity = 500
	offices = append(offices, office)

	office.Id = 2
	office.Name = "Branch A"
	office.Location = "Suburb Area"
	office.Floor = "2nd"
	office.Capacity = 150
	offices = append(offices, office)

	office.Id = 3
	office.Name = "Branch B"
	office.Location = "Downtown"
	office.Floor = "5th"
	office.Capacity = 200
	offices = append(offices, office)

	return offices
}
