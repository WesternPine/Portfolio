package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/*Registers the requested-server-webpage to the actual file/webpage
Then starts the web server.*/
func main() {
	loadHandlers()
	err := http.ListenAndServeTLS(":443", "cert.pem", "privkey.pem", nil)
	if err != nil {
		err = http.ListenAndServe(":80", nil)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

/*Will take values from a configuration file, and load each webpage accordingly*/
func loadHandlers() {
	byteArray, err := ioutil.ReadFile("pages.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	handlers := new(PageHandlers)
	err = json.Unmarshal(byteArray, handlers)
	if err != nil {
		fmt.Println(err.Error())
	}

	for i := 0; i < len(handlers.Handlers); i++ {
		handler := handlers.Handlers[i]
		http.Handle(handler.URLPath, &handler)
	}
}

/*Page Handlers*/
type PageHandlers struct {
	Handlers []PageHandler `json:"pages"`
}

/*Page Handler Object*/
type PageHandler struct {
	URLPath  string `json:"url"`
	FilePath string `json:"file"`
}

/*PageHandler Constructor
The reason for this constructor is so that "DOMAIN.HERE" returns a webpage,
WITHOUT the file extension, AND...
To serve a file that may be in a different directory OR location!
(Look in main function)*/
func NewPageHandler(urlPath string, filePath string) *PageHandler {
	return &PageHandler{
		URLPath:  urlPath,
		FilePath: filePath,
	}
}

/*To send the proper files with the proper extensions.
The file can change depending on the path specified within the constructor,
Or if it is the first file or not.*/
func (this *PageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[1:]

	if !strings.Contains(filePath, ".") {
		filePath = this.FilePath
	}

	data, err := ioutil.ReadFile(filePath)
	if err == nil {
		var contentType string
		if strings.HasSuffix(filePath, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(filePath, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(filePath, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(filePath, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(filePath, ".jpg") {
			contentType = "image/jpg"
		} else if strings.HasSuffix(filePath, ".svg") {
			contentType = "image/svg+xml"
		} else if strings.HasSuffix(filePath, ".mp4") {
			contentType = "video/mp4"
		} else {
			contentType = "text/plain"
		}

		//Have seen both instances working (Added both to be safe)
		w.Header().Set("Content Type", contentType)
		w.Header().Set("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}
