package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var registredServices []Service

func init() {
	file, err := ioutil.ReadFile("files/services.json")

	if err == nil {
		json.Unmarshal(file, &registredServices)
		return
	}
}

func GetService() []Service {
	return registredServices
}

func Register(writer http.ResponseWriter, request *http.Request) {
	var service Service
	err := readResponse(request, &service)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		http.Error(writer, err.Error(), 0600)
		return
	}

	if !verifyIfServiceIsRegistred(service) {
		registredServices = append(registredServices, service)
		writeToFile(registredServices)
	}
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
}

func readResponse(request *http.Request, value interface{}) error {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(value)

	if err != nil {
		panic(err)
	}

	defer request.Body.Close()
	return nil
}

func verifyIfServiceIsRegistred(service Service) bool {
	for _, s := range registredServices {
		if service.Group == s.Group && service.Name == s.Name {
			return true
		}
	}

	return false
}

func writeToFile(registredServices []Service) error {
	file, err := json.Marshal(registredServices)

	if err != nil {
		panic(err)
	}

	return ioutil.WriteFile("files/services.json", file, 0600)
}
