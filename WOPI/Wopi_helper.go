package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
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

// func Obtain_url(w http.ResponseWriter, r *http.Request) {
// 	// Obtain url,find file name in url
// 	// url example :http://Ipserver/viewer/dcomp.php?fileidstr=378b25d2518e113887e.wmv&iswindows=0&optuser=test
// 	file_name := strings.Split(r.RequestURI, "=")[1]
// 	// file_postfix, judge file type, if file type != office, return error
// 	file_postfix := strings.Split(file_name, ".")[1]

// 	edit := Resolve_xml("edit")

// 	fmt.Println(edit)
// }

func Edit_xml() map[string]string {
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
	Edit_dict := make(map[string]string) //edit_url map
	for k, _ := range v.Net.App.Action {
		if strings.EqualFold(v.Net.App.Action[k].Type, "edit") {
			url := strings.Split(v.Net.App.Action[k].Url, "<")[0]
			file_type := v.Net.App.Action[k].Name
			Edit_dict[file_type] = url
		} else {
			continue
		}
	}
	return Edit_dict
}

func View_xml() map[string]string {
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
	View_dict := make(map[string]string) //view_url map
	for k, _ := range v.Net.App.Action {
		if strings.EqualFold(v.Net.App.Action[k].Type, "view") {
			url := strings.Split(v.Net.App.Action[k].Url, "<")[0]
			file_type := v.Net.App.Action[k].Name
			View_dict[file_type] = url
		} else {
			continue
		}
	}
	return View_dict
}

//func View_url() {}

func Edit_url() {
	Edit_dict := Edit_xml()
	file_url := "http://Ipserver/viewer/dcomp.php?fileidstr=378b25d2518e113887e.docx&iswindows=0&optuser=test"
	file_name := strings.Split(file_url, "=")[1]
	file_postfix := strings.Split(file_name, ".")[1]
	if _, ok := Edit_dict[file_postfix]; ok {
		fmt.Println(strings.Join([]string{Edit_dict[file_postfix], file_name}, ""))
	}
}

func View_url(w http.ResponseWriter, r *http.Request) {
	View_url := View_xml()
	file_name := strings.Split(r.RequestURI, "=")[1]
	file_postfix := strings.Split(file_name, ".")[1]
	wopi_host := "WOPISrc=http://10.0.9.127/api/wopi/files/"
	asseen_token := "06lhXK6zWTUi"
	if _, ok := View_url[file_postfix]; ok {
		fmt.Println(strings.Join([]string{View_url[file_postfix], wopi_host, file_name, asseen_token}, ""))
	} else {
		log.Println("Error type")
	}
}

func main() {
	Edit_url()
	router := mux.NewRouter()
	router.HandleFunc("/", View_url)
	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
