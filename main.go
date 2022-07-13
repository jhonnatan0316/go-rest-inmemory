package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type workingEcosystem struct {
	WeId          int    `json:"weId"`
	WeName        string `json:"weName"`
	WeDescription string `json:"weDescription"`
}

type allWorkingEcosystems []workingEcosystem

var workingEcosystems = allWorkingEcosystems{
	{
		WeId:          1,
		WeName:        "Golang",
		WeDescription: "Welcome to the Java Software Engineer Working Ecosystem! The Java Software Engineer career covers all the necessary resources, tools and knowledge necessary to develop an application from the start to taking it to scale, and also to componentize and ensure quality, maintenance and performance. Here you will find also the technical skills, talent manifesto skills and some delivery competences that will accompany you during the multiple stages of your development of the Java deployments environments, such as desktops, web servers and cloud installations.",
	},
	{
		WeId:          2,
		WeName:        "Java",
		WeDescription: "A Working Ecosystem refers to the suite of tools, processes, and expertise used to solve certain types of problems. If you, as a glober, have expertise in a position’s specialty, those skills you have will be related to this WE Some of those areas of knowledge and practices acquired towards a specialty in your field can cross to another ecosystem, allowing you to explore and develop them in the multidisciplinary environment of your Position.",
	},
}

func getWorkingEcosystems(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(workingEcosystems)
}

func createWorkingEcosystem(writer http.ResponseWriter, request *http.Request) {
	var newWe workingEcosystem
	requestBody, error := ioutil.ReadAll(request.Body)
	if error != nil {
		fmt.Fprintf(writer, "Insert a valid Working Ecosystem")
	}

	json.Unmarshal(requestBody, &newWe)

	newWe.WeId = len(workingEcosystems) + 1
	workingEcosystems = append(workingEcosystems, newWe)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(newWe)
}

func getWorkingEcosystem(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	weId, error := strconv.Atoi(vars["weId"])

	if error != nil {
		fmt.Fprintf(writer, "weId Invalid!")
		return
	}

	for _, workingEcosystem := range workingEcosystems {
		writer.Header().Set("Content-Type", "application/json")
		if workingEcosystem.WeId == weId {
			json.NewEncoder(writer).Encode(workingEcosystem)
		}
	}
}

func deletedTask(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	weId, error := strconv.Atoi(vars["weId"])

	if error != nil {
		fmt.Fprintf(writer, "weId Invalid!")
		return
	}

	for index, workingEcosystem := range workingEcosystems {
		writer.Header().Set("Content-Type", "application/json")
		if workingEcosystem.WeId == weId {
			workingEcosystems = append(workingEcosystems[:index], workingEcosystems[index+1:]...)
			fmt.Fprintf(writer, "The working ecosystem %v has been delete succesfully!", weId)
		}
	}
}

func updateWorkingEcosystem(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	var updatedWe workingEcosystem

	weId, error := strconv.Atoi(vars["weId"])
	if error != nil {
		fmt.Fprintf(writer, "weId Invalid!")
		return
	}

	requestBody, error := ioutil.ReadAll(request.Body)
	if error != nil {
		fmt.Fprintf(writer, "You should enter valid Information!")
	}

	json.Unmarshal(requestBody, &updatedWe)

	for index, workingEcosystem := range workingEcosystems {
		if workingEcosystem.WeId == weId {
			workingEcosystems = append(workingEcosystems[:index], workingEcosystems[index+1:]...)
			updatedWe.WeId = weId
			workingEcosystems = append(workingEcosystems, updatedWe)
			fmt.Fprintf(writer, "The working ecosystem with weId %v has been update succesfully!", weId)
		}
	}

}

func indexRoute(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to my in memory api")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/working-ecosystems", getWorkingEcosystems).Methods("GET")
	router.HandleFunc("/working-ecosystems", createWorkingEcosystem).Methods("POST")
	router.HandleFunc("/working-ecosystems/{weId}", getWorkingEcosystem).Methods("GET")
	router.HandleFunc("/working-ecosystems/{weId}", deletedTask).Methods("DELETE")
	router.HandleFunc("/working-ecosystems/{weId}", updateWorkingEcosystem).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))
}
