package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
)

func exit() {
	fmt.Println("Bye!")
	exec.Command("/usr/bin/reset")
	os.Exit(0)
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "domains", Description: "get the domains"},
		{Text: "records", Description: "get the records of a given domain"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func executor(s string) {
	iniFile := "secrets.ini"
	if len(os.Args) > 1 {
		iniFile = os.Args[1]
	}
	iniCredentials, err := getINICredentials(iniFile)
	if err != nil {
		os.Exit(1)
	}
	apikey := iniCredentials.APIKey
	secretkey := iniCredentials.SecretKey
	s = strings.TrimSpace(s)
	f := func(c rune) bool {
		return c == ' '
	}
	inputs := strings.FieldsFunc(s, f)
	switch inputs[0] {
	case "exit":
		exit()
		return
	case "quit":
		exit()
		return
	case "domains":
		domains := getDomains(apikey, secretkey)
		domainPPrinter(domains)
		return
	case "records":
		recordID, _ := strconv.Atoi(inputs[1])
		records := getDNSRecords(apikey, secretkey, recordID)
		recordPPrint(records)
		return
	default:
		return
	}
}
