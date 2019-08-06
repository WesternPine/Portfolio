package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"	
	//"gopkg.in/gomail.v2"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	//start router
	r := mux.NewRouter()

	//Handlers
	r.HandleFunc("/forms", addWebForm).Methods("POST")

	log.Fatal(http.ListenAndServe(":81", r))
}

//Form info model
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
	var form Form
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		//bad json
		w.Write([]byte("false"))
	} else {
		err = addFormToSQL(form)
		//err = emailForm(form)

		if err != nil {
			//internal error
			w.Write([]byte("false"))
			fmt.Println(err.Error())
		} else {
			//message sent
			w.Write([]byte("true"))
		}
	}
}

//func emailForm(f Form) error {
//	m := gomail.NewMessage()
//	m.SetHeader("From", "no-reply-form-submission@westernpine.dev")
//	m.SetHeader("To", "tyler@westernpine.dev")
//	m.SetAddressHeader("Cc", f.Email, f.Name)
//	m.SetHeader("Subject", f.Subject)
//	m.SetBody("text/html", f.Message)
//
//	d := gomail.Dialer{Host: "localhost", Port: 587}
//	if err := d.DialAndSend(m); err != nil {
//    	panic(err)
//	}
//}

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

/*
reads config file
converts to json
opens database connection
creates table if not already existing
returns database and table name
*/
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
