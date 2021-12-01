package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type domain struct {
	Created            int64    `json:"created"`
	FolderID           int      `json:"folderId"`
	GtdEnabled         bool     `json:"gtdEnabled"`
	Updated            int64    `json:"updated"`
	ProcessMulti       bool     `json:"processMulti"`
	ActiveThirdParties []string `json:"activeThirdParties"`
	PendingActionID    int      `json:"pendingActionId"`
	Name               string   `json:"name"`
	ID                 int      `json:"id"`
}

type domainResponse struct {
	TotalRecords int      `json:"totalRecords"`
	TotalPages   int      `json:"totalPages"`
	Data         []domain `json:"data"`
}

type dnsRecord struct {
	Source      int    `json:"source"`
	TTL         int    `json:"ttl"`
	GtdLocation string `json:"gtdLocation"`
	SourceID    int    `json:"sourceId"`
	Failover    bool   `json:"failover"`
	Monitor     bool   `json:"monitor"`
	HardLink    bool   `json:"hardLink"`
	DynamicDNS  bool   `json:"dynamicDns"`
	Failed      bool   `json:"failed"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	ID          int    `json:"id"`
	Type        string `json:"type"`
}

type dnsRecordResponse struct {
	TotalRecords int         `json:"totalRecords"`
	TotalPages   int         `json:"totalPages"`
	Data         []dnsRecord `json:"data"`
}

func addHeaders(apikey string, secretkey string, w *http.Request) {
	secretKeyHMAC, currentTime := generateHMAC(secretkey)
	headers := map[string]string{
		"x-dnsme-apiKey":      apikey,
		"x-dnsme-requestDate": currentTime,
		"x-dnsme-hmac":        secretKeyHMAC,
	}
	for header, value := range headers {
		w.Header.Add(header, value)
	}
}

func getDomains(apikey string, secretkey string) []domain {
	var domains domainResponse
	req, _ := http.NewRequest("GET", "https://api.dnsmadeeasy.com/V2.0/dns/managed/", nil)
	addHeaders(apikey, secretkey, req)
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return domains.Data
	}
	defer resp.Body.Close()
	responseBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(responseBytes, &domains)
	return domains.Data
}

func getDNSRecords(apikey string, secretkey string, recordID int) []dnsRecord {
	var records dnsRecordResponse
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://api.dnsmadeeasy.com/V2.0/dns/managed/%d/records", recordID), nil)
	addHeaders(apikey, secretkey, req)
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return records.Data
	}
	defer resp.Body.Close()
	responseBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(responseBytes, &records)
	return records.Data
}
