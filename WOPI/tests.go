package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Domain struct {
	XMLName     xml.Name `xml:"DomainCheckResult"`
	Domain      string   `xml:"Domain,attr"`
	Available   string   `xml:",attr"`
	ErrorNo     string   `xml:",attr"`
	Description string   `xml:",attr"`
}

type Domains struct {
	XMLName xml.Name `xml:"CommandResponse"`
	Domains []Domain `xml:"DomainCheckResult"`
}

type ApiResponse struct {
	XMLName     xml.Name `xml:"ApiResponse"`
	ApiResponse Domains
}

func main() {

	xmlFile, err := os.Open("data.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer xmlFile.Close()
	XMLdata, _ := ioutil.ReadAll(xmlFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var d ApiResponse
	err = xml.Unmarshal(XMLdata, &d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)
	for k, _ := range d.ApiResponse.Domains {
		fmt.Printf("%s -> %s\n", d.ApiResponse.Domains[k].Domain, d.ApiResponse.Domains[k].Available)
	}
	// fmt.Printf(d.ApiResponse.Domains)
}
