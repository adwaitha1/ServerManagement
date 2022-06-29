package admin

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	dbs "servermanagement.com/connection"
)

// struct data in JSON
type UserDetails struct {
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
	Delete     string `json:"Delete"`
}

// view users list
func View_Role(w http.ResponseWriter, r *http.Request) {
	DB_conn := dbs.DB                                                                                                                               //database connection using funcation
	rows, err := DB_conn.Query("SELECT User_ID,Email_ID,First_Name, Last_Name ,Created_on, Created_by,Updated_on,Updated_by,Role,Teams from USERS") // data selecting from user_table

	if err != nil {
		log.Printf("Failed to build content from sql rows: %v \n", err)
	}

	users := []UserDetails{}
	for rows.Next() {
		var User_ID int
		var Email_ID, First_Name, Last_Name, Created_on, Created_by, Updated_on, Updated_by, Role, Teams string

		err = rows.Scan(&User_ID, &Email_ID, &First_Name, &Last_Name, &Created_on, &Created_by, &Updated_on, &Updated_by, &Role, &Teams)

		if err != nil {
			log.Printf("Failed to build content from sql rows: %v \n", err)
		}
		users = append(users, UserDetails{User_ID: User_ID, Email_ID: Email_ID, First_Name: First_Name, Last_Name: Last_Name, Created_on: Created_on, Created_by: Created_by, Updated_on: Updated_on, Updated_by: Updated_by, Role: Role, Teams: Teams})

	}
	//json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]interface{}{"Listusers": users, "status": "Ok", "Data": "record found"})
}

// // sort users by users_id (optional)
// func Users_ID(w http.ResponseWriter, r *http.Request) {
// 	var db = dbs.Database()
// 	id := r.URL.Query().Get("id") //using GET method

// 	rows, err := db.Query("SELECT * FROM USERS WHERE User_ID =" + id)
// 	counter := 0

// 	if err != nil {
// 		//fmt.Println(err)
// 		panic(err.Error())
// 	}

// 	var users UserDetails
// 	for rows.Next() {
// 		err := rows.Scan(&users.User_ID, &users.Email_ID, &users.Password, &users.First_Name, &users.Last_Name, &users.Created_on, &users.Created_by, &users.Updated_on, &users.Updated_by, &users.Role, &users.Teams, &users.Delete)
// 		w.WriteHeader(http.StatusAccepted)
// 		json.NewEncoder(w).Encode(map[string]string{"status": "Ok", "Data": "record found"})
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		counter++
// 	}

// 	if counter == 0 {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(map[string]string{"status": "400", "Data": "No record found"})

// 	} else {
// 		fmt.Println(users)
// 		json.NewEncoder(w).Encode(users)
// 	}
// }

// func calling
