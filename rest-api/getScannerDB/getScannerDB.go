package getScannerDB

import (
	"encoding/json"
	"fmt"
	"time"
)

var (
	err error
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

func GetScannerDB(respBody []byte) ([]Scanner, error) {
	var body Body

	err = json.Unmarshal(respBody, &body)
	if err != nil {
		return nil, fmt.Errorf("unmarshall error: %v", err)
	}

	return body.Scanners, nil
}
