package main

import (
	"database/sql"
	"encoding/json"
	"errors"
  "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	loadHandlers()
  
	http.HandleFunc("/formsubmissionhandler", addWebForm)
  
	err := http.ListenAndServeTLS(":443", "cert.pem", "privkey.pem", nil)
	if err != nil {
		err = http.ListenAndServe(":80", nil)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

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

type PageHandlers struct {
	Handlers []PageHandler `json:"pages"`
}

type PageHandler struct {
	URLPath  string `json:"url"`
	FilePath string `json:"file"`
}

func NewPageHandler(urlPath string, filePath string) *PageHandler {
	return &PageHandler{
		URLPath:  urlPath,
		FilePath: filePath,
	}
}

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

		w.Header().Set("Content Type", contentType)
		w.Header().Set("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}

type Form struct {
	Name    string `json:"name"`
	EMail   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func NewForm() Form {
	return Form{"BLANK NAME", "BLANK EMAIL", "BLANK SUBJECT", "BLANK MESSAGE"}
}

func addWebForm(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "POST":
      var form Form
    	err := json.NewDecoder(r.Body).Decode(&form)
    	if err != nil {
    		w.Write([]byte("false"))
    	} else {
    		err = addFormToSQL(form)
    		if err != nil {
    			w.Write([]byte("false"))
    			fmt.Println(err.Error())
    		} else {
    			w.Write([]byte("true"))
    		}
    	}
    default:
        fmt.Fprintf(w, "Sorry, only the POST method is supported here for form submissions.")
  }
}

func addFormToSQL(f Form) error {
	var table string
	var db *sql.DB
	var err error
	table, db, err = openCon()
	if err != nil {
		return err
	}
	pre, err := db.Prepare("INSERT INTO " + table + "(id,name,email,subject,message) VALUES(?,?,?,?,?);")
	if err != nil {
		return err
	}
 _, err = pre.Exec("0", f.Name, f.EMail, f.Subject, f.Message)
 if err != nil {
	 return err
 }
	defer db.Close()
	return err
}

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IP       string `json:"ip"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Table    string `json:"table"`
}

func openCon() (string, *sql.DB, error) {
	byteArray, err := ioutil.ReadFile("contactConfig.json")
	if err != nil {
		fmt.Println(err.Error())
		return "", nil, errors.New("unable to read 'contactConfig.json'")
	}
	config := new(Config)
	err = json.Unmarshal(byteArray, config)
	if err != nil {
		fmt.Println(err.Error())
		return "", nil, errors.New("unable to convert json to memory object")
	}
	db, err := sql.Open("mysql", config.Username+":"+config.Password+"@tcp("+config.IP+":"+config.Port+")/"+config.Database)
	if err != nil {
		fmt.Println(err.Error())
		return config.Table, nil, errors.New("unable to open database connection - ensure correct credentials")
	}
	qry, err := db.Query("CREATE TABLE IF NOT EXISTS `" + config.Table + "` (`id` int(255) NOT NULL AUTO_INCREMENT, `name` text NOT NULL, `email` text NOT NULL, `subject` text NOT NULL, `message` text NOT NULL, PRIMARY KEY (`id`) ) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;")
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return config.Table, nil, errors.New("unable to create table in database")
	} else {
		qry.Close()
	}
	return config.Table, db, nil
}
