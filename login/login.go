package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	dbs "servermanagement.com/connection"
	"servermanagement.com/cors"
)

type loginDetails struct {
	Email_Id string `json:"emailId"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("InfobellItSolutions")

func Login(w http.ResponseWriter, r *http.Request) {
	cors.SetupCORS(&w, r)
	DB_conn := dbs.DB
	var l loginDetails

	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if l.Email_Id == "" || l.Password == "" {
		//json.NewEncoder(w).Encode(map[string]string{"status": "400", "Data": "Bad Request"})
		w.WriteHeader(http.StatusBadRequest)
		return

	} else {
		row := DB_conn.QueryRow("SELECT User_ID,Email_ID,Password,Role from Users where Email_ID = '" + l.Email_Id + "'")
		var EMAIL, PASSWORD, ROLE string
		var id int
		ID := strconv.Itoa(id)
		err_scan := row.Scan(&ID, &EMAIL, &PASSWORD, &ROLE)
		if err_scan != nil {
			panic(err_scan.Error())
		}
		fmt.Println("Compared result :", CheckPasswordHash(l.Password, PASSWORD))
		if ID == "" || EMAIL == "" || PASSWORD == "" || ROLE == "" {
			json.NewEncoder(w).Encode(map[string]string{"status": "OK", "Message": "No record found"})
		} else if CheckPasswordHash(l.Password, PASSWORD) {
			expirationTime := time.Now().Add(time.Minute * 5)
			claims := &Claims{
				Username: EMAIL,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(jwtKey)
			if err != nil {
				fmt.Println("Error in generating JWT Err : ", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
			})
			json.NewEncoder(w).Encode(map[string]string{"User Id": ID, "Role": ROLE, "status": "200 OK", "Message": "Successfully Logged In"})
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"status": "401", "Message": "Unauthorised password"})
		}

	}
}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))
}
