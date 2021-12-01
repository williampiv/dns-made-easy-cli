package main

import "fmt"

func domainPPrinter(d []domain) {
	for _, i := range d {
		fmt.Printf("Domain: %s\nID:%d\n\n", i.Name, i.ID)
	}
}

func recordPPrint(r []dnsRecord) {
	for _, i := range r {
		fmt.Printf("Name: %s\nValue: %s\nType: %s\n\n", i.Name, i.Value, i.Type)
	}
}
