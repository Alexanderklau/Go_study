package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Wopi struct {
	Net     NetZone  `xml:"net-zone"`
	XMLName xml.Name `xml:wopi-discovery`
}

type NetZone struct {
	App App `xml:"app"`
}

type App struct {
	Action []Action `xml:"action"`
}

type Action struct {
	Type string `xml:"name,attr"`
	Name string `xml:"ext,attr"`
	Url  string `xml:"urlsrc,attr"`
}

func index() map[string]string {
	file, err := os.Open("discovery.xml")
	if err != nil {
		fmt.Println("error: %v", err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error: %v", err)
	}
	v := Wopi{}
	err = xml.Unmarshal(data, &v)
	dict := make(map[string]string)
	//	fmt.Println(v.Net.App.Action[1].Type)
	for k, _ := range v.Net.App.Action {
		if strings.EqualFold(v.Net.App.Action[k].Type, "edit") {
			a := strings.Split(v.Net.App.Action[k].Url, "<")[0]
			b := v.Net.App.Action[k].Name
			dict[b] = a
		}
	}
	return dict
}

func main() {
	fmt.Println(index())
}
