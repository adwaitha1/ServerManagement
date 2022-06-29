package admin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	dbs "servermanagement.com/connection"
	"servermanagement.com/cors"
)

// JSON File Formatter
type userDetails struct {
	User_ID    int    `json:"User_ID"`
	Email_ID   string `json:"Email_ID"`
	Password   string `json:"Password"`
	First_Name string `json:"First_Name"`
	Last_Name  string `json:"Last_Name"`
	Created_on string `json:"Created_on"`
	Created_by string `json:"Created_by"`
	Updated_on string `json:"Updated_on"`
	Updated_by string `json:"Updated_by"`
	Role       string `json:"Role"`
	Teams      string `json:"Teams"`
	Delete     int    `json:"Delete"`
	Token      string `json:"Token"`
}

//list all users
func View_user(w http.ResponseWriter, r *http.Request) {
	cors.SetupCORS(&w, r)
	DB_conn := dbs.DB //database connection using function
	rows, err := DB_conn.Query("select user_id,email_id,first_name,last_name, created_on,created_by," +
		"updated_on,updated_by,role,teams,delete from users") // data selecting from user_table
	if err != nil {
		log.Printf("Failed to build content from sql rows: %v \n", err)
	}

	users := []userDetails{}
	for rows.Next() {
		var User_ID int
		var Delete bool
		var Email_ID, First_Name, Last_Name, Created_on, Created_by, Updated_on, Updated_by, Role, Teams string

		err = rows.Scan(&User_ID, &Email_ID, &First_Name, &Last_Name, &Created_on, &Created_by, &Updated_on, &Updated_by, &Role, &Teams, &Delete)

		if err != nil {
			log.Printf("Failed to build content from sql rows: %v \n", err)
		}
		users = append(users, userDetails{User_ID: User_ID, Email_ID: Email_ID, First_Name: First_Name, Last_Name: Last_Name, Created_on: Created_on, Created_by: Created_by, Updated_on: Updated_on, Updated_by: Updated_by, Role: Role, Teams: Teams})

	}
	//json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]interface{}{"Listusers": users, "status": "Ok", "Data": "record found"})
}

// LisT users by users_id (optional)
func Users_ID(w http.ResponseWriter, r *http.Request) {
	cors.SetupCORS(&w, r)
	DB_conn := dbs.DB
	id := r.URL.Query().Get("id") //using GET method

	rows, err := DB_conn.Query("select user_id,email_id,password,first_name,last_name," +
		"created_on,created_by,updated_on,updated_by,role,teams," +
		"delete from users WHERE User_ID =" + id)
	counter := 0
	if err != nil {
		//fmt.Println(err)
		panic(err.Error())
	}
	var users userDetails
	for rows.Next() {
		err := rows.Scan(&users.User_ID, &users.Email_ID, &users.Password, &users.First_Name, &users.Last_Name, &users.Created_on, &users.Created_by, &users.Updated_on, &users.Updated_by, &users.Role, &users.Teams, &users.Delete)
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{"status": "Ok", "Data": "record found"})
		if err != nil {
			panic(err.Error())
		}
		counter++
	}
	if counter == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"status": "400", "Data": "No record found"})

	} else {
		fmt.Println(users)
		json.NewEncoder(w).Encode(users)
	}
}

// create user
func Create_User(w http.ResponseWriter, r *http.Request) {
	cors.SetupCORS(&w, r)
	DB_conn := dbs.DB
	user := &userDetails{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Invalid details", "status": "400 "})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if _, err = DB_conn.Query("insert into Users values ($1,$2,$3, $4, $5, $6, $7, $8, $9, $10, $11,$12,$13)", user.User_ID, user.Email_ID, string(hashedPassword), user.First_Name, user.Last_Name, user.Created_on, user.Created_by, user.Updated_on, user.Updated_by, user.Role, user.Teams, user.Delete, user.Token); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Invalid details", "status": "400 "})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "created user succesfully", "status": "200 OK"})
}
