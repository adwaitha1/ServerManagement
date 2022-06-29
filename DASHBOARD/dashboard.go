package DASHBOARD

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	dbs "servermanagement.com/connection"
	//"servermanagement.com/cors"
)

//--------------------------------------------Dashboard:number of reserved and vacant asset---------------------------------------------------------------------

func GetDashboard1(write http.ResponseWriter, request *http.Request) {
	DB_conn := dbs.DB
	var count, total int
	err := DB_conn.QueryRow("SELECT count(*) from asset WHERE reserved is not null").Scan(&count) // exporting table
	if err != nil {
		json.NewEncoder(write).Encode(map[string]interface{}{"Status Code": "400 Bad Request", "Message": err})
	}
	err1 := DB_conn.QueryRow("SELECT COUNT(*) FROM asset").Scan(&total) // exporting table
	if err1 != nil {
		json.NewEncoder(write).Encode(map[string]interface{}{"Status Code": "400 Bad Request", "Message": err1})
	}
	vacant := total - count
	json.NewEncoder(write).Encode(map[string]interface{}{"reserved": count, "vacant": vacant, "Status": "200 OK", "Message": "Updated Statistics"})

}

//--------------------------------------------Dashboard:number of reserved and vacant asset in cluster---------------------------------------------------------------------

func GetDashboard2(write http.ResponseWriter, request *http.Request) {
	DB_conn := dbs.DB
	var count, total int
	err := DB_conn.QueryRow("SELECT count(distinct cluster_id) from asset where reserved='t' group by cluster_id").Scan(&count) // exporting table
	if err != nil {
		log.Printf("Failed to open table: %v \n", err)
	}
	err1 := DB_conn.QueryRow("select count(distinct cluster_id)from asset").Scan(&total) // exporting table
	if err1 != nil {
		log.Printf("Failed to open table: %v \n", err1)
	}
	json.NewEncoder(write).Encode(map[string]interface{}{"reserved": count, "Total": total, "Status": "200 OK", "Message": "Updated Statistics"})

}

//--------------------------------------------Dashboard:number of reserved by location---------------------------------------------------------------------
func GetDashboard3(write http.ResponseWriter, request *http.Request) {
	DB_conn := dbs.DB
	var reserve, vacant int
	var asset_location string
	rows, err := DB_conn.Query(" SELECT asset_location, COUNT(CASE WHEN reserved='t' THEN 1 ELSE NULL END)AS reserved,  COUNT(CASE WHEN reserved='f' or reserved is null THEN 1 ELSE NULL END)AS vacant FROM asset group by asset_location")
	if err != nil {
		log.Fatal(err)
		json.NewEncoder(write).Encode(map[string]interface{}{"status": "400 Bad Request", "Message": err})
	}
	fmt.Printf("Location|Reserved|Vacant\n")
	for rows.Next() {
		err := rows.Scan(&asset_location, &reserve, &vacant)
		if err != nil {
			log.Fatal(err)
			json.NewEncoder(write).Encode(map[string]interface{}{"status": "400 Bad Request", "Message": err})
		}
		fmt.Printf("%v : %v: %v\n ", asset_location, reserve, vacant)
		lct := map[string]interface{}{"Location": asset_location, "Reserved": reserve, "Vacant": vacant}
		json.NewEncoder(write).Encode(map[string]interface{}{"Dashboard": lct, "Status": "200 OK", "Message": "Updated Statistics"})

	}
}
