package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

const apiUrl = "https://ipwho.is/"

type IPWhoisResponse struct {
	IP            string  `json:"ip"`
	Success       bool    `json:"success"`
	Message       string  `json:"message"`
	Type          string  `json:"type"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continent_code"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"country_code"`
	Region        string  `json:"region"`
	RegionCode    string  `json:"region_code"`
	City          string  `json:"city"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	IsEU          bool    `json:"is_eu"`
	Postal        string  `json:"postal"`
	CallingCode   string  `json:"calling_code"`
	Capital       string  `json:"capital"`
	Borders       string  `json:"borders"`
	Org           string  `json:"org"`
	ISP           string  `json:"isp"`
	//Timezone      int     `json:"timezone"`
}

func main() {
	// Get IP address from command-line argument
	ip := flag.String("ip", "", "IP address to look up")
	flag.Parse()

	if *ip == "" {
		fmt.Println("Usage: go run main.go -ip=<IP_ADDRESS>")
		os.Exit(1)
	}

	// Make a request to the ipwhois API
	resp, err := http.Get(apiUrl + *ip)
	if err != nil {
		fmt.Printf("Error fetching IP info: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		os.Exit(1)
	}

	// Parse the JSON response
	var result IPWhoisResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Display the result
	if result.Success {
		fmt.Printf("IP: %s\n", result.IP)
		fmt.Printf("Continent: %s\n", result.Continent)
		fmt.Printf("Country: %s\n", result.Country)
		fmt.Printf("City: %s\n", result.City)
		//fmt.Printf("Timezone: %s\n", result.Timezone)
		fmt.Printf("Latitude: %v, Longitude: %v\n", result.Latitude, result.Longitude)

	} else {
		fmt.Println("Failed to fetch IP details")
	}

}
