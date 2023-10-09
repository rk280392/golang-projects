package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/rk280392/rest-api/getScannerDB"
	"github.com/rk280392/rest-api/parseURL"
)

type Scanner struct {
	CveDBCreateTime   time.Time `json:"cvedb_create_time"`
	CveDBVersion      string    `json:"cvedb_version"`
	JoinedTimestamp   int64     `json:"joined_timestamp"`
	ScannedContainers int       `json:"scanned_containers"`
	ScannedHosts      int       `json:"scanned_hosts"`
	ScannedImages     int       `json:"scanned_images"`
	ScannedServerless int       `json:"scanned_serverless"`
	Server            string    `json:"server"`
}

type Body struct {
	Scanners []Scanner `json:"scanners"`
}

var (
	requestURL string
	apiKey     string
)

func main() {

	flag.StringVar(&requestURL, "url", "", "url to access")
	flag.StringVar(&apiKey, "apikey", "", "enter the apikey to access the apiserver")
	flag.Parse()

	respBody, err := parseURL.ParseURL(requestURL, apiKey)
	if err != nil {
		fmt.Printf("An error happened while parsing the URL %v", err)
	}

	resultScanDB, err := getScannerDB.GetScannerDB(respBody)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if resultScanDB == nil {
		fmt.Printf("No response\n")
		os.Exit(1)
	}
	for _, scanner := range resultScanDB {
		fmt.Printf("Json Parsed\nCveDBVersion: %v\nCveDBVersion: %s\n", scanner.CveDBCreateTime, scanner.CveDBVersion)
	}

	/*
		resultControllers, err := getControllers.GetControllers(respBody)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if resultScanDB == nil {
			fmt.Printf("No response\n")
			os.Exit(1)
		}
		for _, controllers := range resultControllers {
			fmt.Printf("Json Parsed\nControllerName: %v\nConnectionStatus: %s\n", controllers.Name, controllers.ConnStatus)
		}
	*/

}
