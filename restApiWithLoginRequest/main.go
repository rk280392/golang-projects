package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/rk280392/restApiWithLoginRequest/doLoginRequest"
	"github.com/rk280392/restApiWithLoginRequest/getControllers"
	"github.com/rk280392/restApiWithLoginRequest/getResponseBody"
	"github.com/rk280392/restApiWithLoginRequest/getScannerDB"
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
	password   string
	username   string
)

func parseURL(requestURL string) (*url.URL, error) {
	var emptyURL *url.URL
	parsedURL, err := url.ParseRequestURI(requestURL)
	if err != nil {
		flag.Usage()
		return emptyURL, fmt.Errorf("validation error: URL is not valid: %s", err)
	}
	return parsedURL, nil
}

func main() {

	flag.StringVar(&requestURL, "url", "", "url to access")
	flag.StringVar(&username, "u", "", "enter the username for the apiserver")
	flag.StringVar(&password, "p", "", "enter the password for the apiserver")
	flag.Parse()

	parsedURL, err := parseURL(requestURL)
	if err != nil {
		fmt.Printf("An error happened while parsing the URL %v", err)
		os.Exit(1)
	}

	token := ""
	if password != "" && username != "" {
		token, err = doLoginRequest.GetToken(parsedURL.Scheme+"://"+parsedURL.Host, username, password)
		if err != nil {
			fmt.Printf("An error happened while login: %v", err)
			os.Exit(1)
		}
	}

	respBody, err := getResponseBody.GetResponseBody(parsedURL, token)
	if err != nil {
		fmt.Printf("An error happened while token %v", err)
	}

	switch {
	case strings.Contains(requestURL, "scanner"):
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
	case strings.Contains(requestURL, "controller"):

		resultControllers, err := getControllers.GetControllers(respBody)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if resultControllers == nil {
			fmt.Printf("No response\n")
			os.Exit(1)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
		fmt.Fprintln(w, "ControllerName\tConnectionStatus\t\tImage")
		for _, controllers := range resultControllers {
			image := controllers.Labels.NeuvectorImage + ":" + controllers.Version
			fmt.Fprintf(w, "%s\t%s\t%s\n", controllers.DisplayName, controllers.ConnStatus, image)
		}
		w.Flush()
	default:
		fmt.Printf("No handler found for URL: %s\n", requestURL)
	}

}
