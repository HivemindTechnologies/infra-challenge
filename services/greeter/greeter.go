package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var client = &http.Client{}
var hostname = os.Getenv("HOSTNAME")
var greetingProviderUrl = os.Getenv("GREETING_PROVIDER_URL")

func main() {
	fmt.Println("Hivemind's Greeter")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	greetingStr, err := GetGreeting()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmtStr := fmt.Sprintf("%s! I'm %s.", greetingStr, hostname)
	fmt.Println(fmtStr)
	fmt.Fprintln(w, fmtStr)
}

func GetGreeting() (string, error) {
	req, err := http.NewRequest("GET", greetingProviderUrl, nil)
	if err != nil {
		return "", fmt.Errorf("Error creating HTTP request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body: %v", err)
	}

	return string(body), nil
}
