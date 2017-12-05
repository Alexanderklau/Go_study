package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
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

// BaseFileName: 文件名。
// OwnerId: 文件所有者的唯一编号。
// Size: 文件大小，以bytes为单位。
// SHA256: 文件的256位bit的SHA-2编码散列内容。（Wordweb app必有，excel和ppt可以为null）
// Version: 文件版本号，文件如果被编辑，版本号也要跟着改变。
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

	router := mux.NewRouter()
	router.HandleFunc("/api/wopi/files/{file_name}", GetFileInfo2).Methods(http.MethodGet)
	router.HandleFunc("/api/wopi/files/{file_name}/contents", GetFileContent).Methods(http.MethodGet)
	router.HandleFunc("/api/wopi/files/{file_name}/contents", PostFileContent).Methods(http.MethodPost)

	err := http.ListenAndServe(":80", router)
	log.Println(router)
	if err != nil {
		log.Println("http listen err: ", err)
	}
}
