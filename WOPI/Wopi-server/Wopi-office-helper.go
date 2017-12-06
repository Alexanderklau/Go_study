package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type fileInfo struct {
	BaseFileName   string `json:"BaseFileName"`
	OwnerId        string `json:"OwnerId"`
	Size           int64  `json:"Size"`
	SHA256         string `json:"SHA256"`
	Version        string `json:"Version"`
	SupportsUpdate bool   `json:"SupportsUpdate,omitempty"`
	UserCanWrite   bool   `json:"UserCanWrite,omitempty"`
	SupportsLocks  bool   `json:"SupportsLocks,omitempty"`
}

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
			url_type := strings.Split(v.Net.App.Action[k].Url, "/")[3]
			if strings.EqualFold(url_type, "x") {
				url := strings.Split(v.Net.App.Action[k].Url, "<")[0]
				file_type := v.Net.App.Action[k].Name
				Edit_dict[file_type] = url
			} else {
				url := strings.Split(v.Net.App.Action[k].Url, "<")[0]
				proxy := strings.Replace(url, "10.0.7.96", "10.0.7.95:8090", -1)
				file_type := v.Net.App.Action[k].Name
				Edit_dict[file_type] = proxy
			}
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
			url_type := strings.Split(v.Net.App.Action[k].Url, "/")[3]
			if strings.EqualFold(url_type, "x") {
				url := strings.Split(v.Net.App.Action[k].Url, "<")[0]
				file_type := v.Net.App.Action[k].Name
				View_dict[file_type] = url
			} else {
				url := strings.Split(v.Net.App.Action[k].Url, "<")[0]
				proxy := strings.Replace(url, "10.0.7.96", "10.0.7.95:8090", -1)
				file_type := v.Net.App.Action[k].Name
				View_dict[file_type] = proxy
			}
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
	datrix_url := "http://10.0.9.139/viewer/dcomp.php?"
	download_url := strings.Join([]string{datrix_url, file}, "")
	file_name := strings.Split(file, "=")[1]
	file_postfix := strings.Split(file_name, ".")[1]
	file_postfixs := strings.Split(file_postfix, "&")[0]
	file_files := strings.Split(file_postfixs, "&")[0]
	res, err := http.Get(download_url)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(file_files)
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
	wopi_host := "WOPISrc=http://10.0.9.127/api/wopi/files/"
	access_token := "&access_token=06lhXK6zWTUi"
	var info urlinfo
	if _, ok := Edit_urls[file_files]; ok {
		view_url := (strings.Join([]string{Edit_urls[file_files], wopi_host, file_name, access_token}, ""))
		info.Url = view_url
		log.Println(view_url)
	} else {
		log.Println("Error type")
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	log.Println(info)
	json.NewEncoder(w).Encode(info)
	log.Println("Download_url done.....")
}

func GetFileContent(w http.ResponseWriter, r *http.Request) {
	log.Println("GetFileContent start.......")

	vals := r.URL.Query()
	tmp, ok := vals["access_token"]
	if !ok || len(tmp[0]) == 0 {
		log.Println("access_token not found!")
	}

	vars := mux.Vars(r)
	fileName := vars["file_name"]
	if len(fileName) == 0 {
		log.Println("file_name empty!")
	}
	log.Println("file_name: ", fileName)

	testFilePath := path.Join(".", fileName)

	data, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		log.Println("read file err: ", err)
		return
	}

	w.Header().Set("Content-type", "application/octet-stream")
	w.Write(data)
	//json.NewEncoder(w).Encode(info)
	log.Println("GetFileContent done !")
}

func Download(w http.ResponseWriter, r *http.Request) {
	View_urls := View_xml()
	file := strings.Split(r.RequestURI, "?")[1]
	log.Println(file)
	datrix_url := "http://10.0.9.139/viewer/dcomp.php?"
	download_url := strings.Join([]string{datrix_url, file}, "")
	file_name := strings.Split(file, "=")[1]
	file_postfix := strings.Split(file_name, ".")[1]
	file_postfixs := strings.Split(file_postfix, "&")[0]
	file_files := strings.Split(file_name, "&")[0]
	res, err := http.Get(download_url)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(file_files)
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
	wopi_host := "WOPISrc=http://10.0.9.127/api/wopi/files/"
	access_token := "&access_token=06lhXK6zWTUi"
	var info urlinfo
	if _, ok := View_urls[file_postfixs]; ok {
		view_url := (strings.Join([]string{View_urls[file_postfixs], wopi_host, file_name, access_token}, ""))
		info.Url = view_url
		log.Println(view_url)
	} else {
		log.Println("Error type")
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	log.Println(info)
	json.NewEncoder(w).Encode(info)
	log.Println("Download_url done.....")
}

func PostFileContent(w http.ResponseWriter, r *http.Request) {
	log.Println("PostFileContent start..........")

	vals := r.URL.Query()
	tmp, ok := vals["access_token"]
	if !ok || len(tmp[0]) == 0 {
		log.Println("access_token not found!")
	}

	vars := mux.Vars(r)
	fileName := vars["file_name"]
	filename := strings.Split(fileName, ".")[0]
	log.Println(filename)
	filetype := strings.Split(fileName, ".")[1]
	t := time.Now()
	timestamp := strconv.FormatInt(t.Unix(), 10)
	filenames := filename + timestamp + "." + filetype
	log.Println(filenames)

	if len(fileName) == 0 {
		log.Println("file_name empty!")
	}

	log.Println("file_name: ", filenames)

	testFilePath := path.Join(".", filenames)

	log.Println("PATH: ", testFilePath)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("body empty")
	}

	ioutil.WriteFile(testFilePath, body, os.ModeAppend)

	w.Header().Set("Content-type", "application/octet-stream")

	// get Path,get file

}

//获取单个文件的大小
func getSize(path string) int64 {
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	fileSize := fileInfo.Size()
	return fileSize
}

func getMD5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Println("Open", err)
		return "", nil
	}

	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		log.Println("Copy", err)
		return "", nil
	}

	return string(md5hash.Sum(nil)), nil
}

func SHA256File(path string) (string, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	h := sha256.Sum256(buf)
	return base64.StdEncoding.EncodeToString(h[:]), nil
}

// GetFileInfo get file info
// /api/wopi/files/{file_name_xxx}?access_token=xxxxxx
func GetFileInfo2(w http.ResponseWriter, r *http.Request) {
	log.Println("GetFileInfo222")

	vals := r.URL.Query()
	tmp, ok := vals["access_token"]
	if !ok || len(tmp[0]) == 0 {
		log.Println("access_token not found!")
	}

	vars := mux.Vars(r)
	fileName := vars["file_name"]
	if len(fileName) == 0 {
		log.Println("file_name empty!")
	}
	log.Println("file_name: ", fileName)

	testFilePath := path.Join(".", fileName)

	log.Println("PATH: ", testFilePath)

	var info fileInfo
	info.BaseFileName = fileName
	info.OwnerId = "admin"
	info.Size = getSize(testFilePath)
	info.SHA256, _ = SHA256File(testFilePath)
	log.Println("debug: sha256_b42: ", info.SHA256)
	// String value: eIMevgBhTd8Iqh1VjWbfWx7wd5vQvmDxlABMfz+pTiI=
	//info.SHA256 = "eIMevgBhTd8Iqh1VjWbfWx7wd5vQvmDxlABMfz+pTiI="

	info.Version = "2222"
	info.UserCanWrite = true
	info.SupportsLocks = true

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(info)

	log.Println("GetFileInfo222 done...")
}

func main() {
	rounter := mux.NewRouter()
	rounter.HandleFunc("/api/edit", Edit_url).Methods(http.MethodGet)
	rounter.HandleFunc("/api/download", Download).Methods(http.MethodGet)
	rounter.HandleFunc("/api/wopi/files/{file_name}", GetFileInfo2).Methods(http.MethodGet)
	rounter.HandleFunc("/api/wopi/files/{file_name}/contents", GetFileContent).Methods(http.MethodGet)
	rounter.HandleFunc("/api/wopi/files/{file_name}/contents", PostFileContent).Methods(http.MethodPost)
	rounter.PathPrefix("/").Handler(http.FileServer(assetFS()))

	err := http.ListenAndServe(":80", rounter)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
