package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	dbs "servermanagement.com/connection"
	"servermanagement.com/cors"
	auth "servermanagement.com/login"
)

type changepwd struct {
	User_id     int    `json:"user_id"`
	Oldpassword string `json:"oldpassword"`
	Newpassword string `json:"newpassword"`
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	cors.SetupCORS(&w, r)
	DB_conn := dbs.DB
	//id := r.URL.Query().Get("id")
	//fmt.Println(id)
	var p changepwd

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if p.Oldpassword == "nil" || p.Newpassword == "nil" || p.User_id == 0 {
		json.NewEncoder(w).Encode(map[string]string{"status": "400", "Data": "Bad Request"})
		w.WriteHeader(http.StatusBadRequest)

	} else {
		//id := strconv.Itoa(p.User_id)
		row := DB_conn.QueryRow("SELECT User_ID,Password from Users where User_ID = $1", p.User_id)
		//fmt.Println(row)
		var db_user, Password string
		err_scan := row.Scan(&db_user, &Password)
		if err_scan != nil {
			//panic(err_scan.Error())
			fmt.Println(err_scan.Error())
		}
		if db_user == "" || Password == "" {
			json.NewEncoder(w).Encode(map[string]string{"status": "OK", "Data": "No record found"})
		} else {
			//if user is available in table and password you entered matches the old password,new password is updated on table.
			temp_pwd, _ := auth.GeneratehashPassword(p.Newpassword)
			fmt.Println(temp_pwd)
			if auth.CheckPasswordHash(p.Oldpassword, Password) {
				hash_pwd, err_h := auth.GeneratehashPassword(p.Newpassword)
				fmt.Println(hash_pwd)
				if err_h != nil {
					log.Fatal(err_h)
				}
				change, err := DB_conn.Exec("update Users set Password =$1 where User_ID=$2", hash_pwd, p.User_id)
				if err != nil {
					log.Fatal(err)
				}
				affectedRow, err := change.RowsAffected()
				json.NewEncoder(w).Encode(map[string]interface{}{"status": "200", "Data": "Password updated", "updated row": affectedRow})
			} else {
				json.NewEncoder(w).Encode(map[string]string{"status": "401", "Data": "Unauthorised password"})

			}
		}

	}

}
