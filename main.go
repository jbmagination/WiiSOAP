package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// The Check struct(ure) will attempt to retrieve all the namespace data.
// Assuming that namespaces that don't exist are given a "nil", the first result that  isn't a "nil" will be used as the template response.

type Check struct {
	// ECommerce Namespaces
	CDS  string `ecs:"CheckDeviceStatus"`
	LET  string `ecs:"ListETickets"`
	NETS string `ecs:"NotifyETicketsSynced"`
	PT   string `ecs:"PurchaseTitle"`

	// Identity Authentication Namespaces
	CR  string `ecs:"CheckRegistration"`
	GRI string `ecs:"GetRegistrationInfo"`
	REG string `ecs:"Register"`
	UNR string `ecs:"Unregister"`
}

func main() {
	file, err := ioutil.ReadFile("SOAP-WSC/ECS/checkDeviceStatus.xml")
	// ChRes is a variable that's in the form of JSON. This organises all the data into sub-variables like ChRes.CDS.
	// This is probably my favourite thing in GoLang to be honest.
	if err != nil {
		panic(err)
		return
	}
	ChRes := Check{
		CDS:  "",
		LET:  "",
		NETS: "",
		PT:   "",
		CR:   "",
		GRI:  "",
		REG:  "",
		UNR:  "",
	}
	err = xml.Unmarshal(file, &ChRes)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(ChRes)

	// http.ListenAndServe starts a HTTP server, which is important to take note of as we will be using this to deliver the SOAP.
	http.ListenAndServe(":80", nil)
}
