package main

import (
	"encoding/json"
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

type urlinfo struct {
	Url string `json:"Url"`
}

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

func Edit_url(w http.ResponseWriter, r *http.Request) {
	Edit_urls := Edit_xml()
	file := strings.Split(r.RequestURI, "src=")[1]
	file_name := strings.Split(file, "=")[1]
	log.Println(file_name)
	file_postfix := strings.Split(file_name, ".")[1]
	log.Println(file_postfix)
	wopi_host := "WOPISrc=http://10.0.9.127/api/wopi/files/"
	asseen_token := "&assen_token=06lhXK6zWTUi"
	var info urlinfo
	if _, ok := Edit_urls[file_postfix]; ok {
		Edit_url := (strings.Join([]string{Edit_urls[file_postfix], wopi_host, file_name, asseen_token}, ""))
		info.Url = Edit_url
	} else {
		log.Println("Error type")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
	log.Println("Edit_urls done.....")
}

func View_url(w http.ResponseWriter, r *http.Request) {
	View_urls := View_xml()
	file := strings.Split(r.RequestURI, "src=")[1]
	file_name := strings.Split(file, "=")[1]
	file_postfix := strings.Split(file_name, ".")[1]
	log.Println(file_postfix)
	wopi_host := "WOPISrc=http://10.0.9.127/api/wopi/files/"
	access_token := "&access_token=06lhXK6zWTUi"
	var info urlinfo
	if _, ok := View_urls[file_postfix]; ok {
		view_url := (strings.Join([]string{View_urls[file_postfix], wopi_host, file_name, access_token}, ""))
		info.Url = view_url
	} else {
		log.Println(r.RequestURI)
		log.Println("Error type")
	}
	w.Header().Set("Content-Type", "application/json")
	log.Println(info)
	json.NewEncoder(w).Encode(info)
	log.Println("View_url done.....")
}

func Download(w http.ResponseWriter, r *http.Request) {
	View_urls := View_xml()
	file := strings.Split(r.RequestURI, "src=")[1]
	file_name := strings.Split(file, "=")[1]
	file_postfix := strings.Split(file_name, ".")[1]
	file_postfixs := strings.Split(file_postfix, "&")[0]
	file_user := strings.Split(file, "&")[1]
	log.Println(file)
	wopi_host := "WOPISrc=http://10.0.9.127/api/wopi/files/"
	access_token := "&access_token=06lhXK6zWTUi"
	var info urlinfo
	if _, ok := View_urls[file_postfixs]; ok {
		view_url := (strings.Join([]string{View_urls[file_postfixs], wopi_host, file_name, access_token}, ""))
		info.Url = view_url
	} else {
		log.Println("Error type")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
	log.Println("View_url done.....")
}

func main() {
	rounter := mux.NewRouter()
	rounter.HandleFunc("/api/view", View_url).Methods(http.MethodGet)
	rounter.HandleFunc("/api/edit", Edit_url).Methods(http.MethodGet)
	rounter.HandleFunc("/api/download", Download).Methods(http.MethodGet)

	err := http.ListenAndServe(":9090", rounter)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
