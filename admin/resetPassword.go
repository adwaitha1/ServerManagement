//Admin reset the user's password
package admin

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	dbs "servermanagement.com/connection"
	"servermanagement.com/cors"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Users_DETAILS struct {
	User_ID  int    `json:"User_ID"`
	Password string `json:"Password"`
}

// Reset password function resets the user password based on user id.
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	DB_conn := dbs.DB
	cors.SetupCORS(&w, r)

	var p Users_DETAILS //declare a variable p for type Users_DETAILS
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), 14)

	id := strconv.Itoa(p.User_ID)
	row := DB_conn.QueryRow("SELECT User_ID from Users where User_ID =" + id)

	var UserId int
	err_scan := row.Scan(&UserId)
	if err_scan != nil {
		log.Fatal(err_scan)
	}
	if p.User_ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Record not found", "Status Code": "400 BadRequest"})

	} else {
		_, err1 := DB_conn.Exec("UPDATE Users SET Password=$2 WHERE User_ID=$1;", p.User_ID, string(hashedPassword))

		if err1 != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err1, "Status Code": "400 BadRequest"})
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Record_updated!", "Status Code": "200 OK"})

	}
}
