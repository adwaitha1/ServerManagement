package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	dbs "servermanagement.com/connection"
)

type Asset struct {
	Asset_ID       int       `json:"Asset_Id"`
	Manufacturer   string    `json:"Manufacturer"`
	BMC_IP         string    `json:"BMC_IP"`
	BMC_user       string    `json:"BMC_USER"`
	Asset_location string    `json:"Asset_location"`
	Reserved       bool      `json:"Reserved"`
	Assigned_to    int       `json:"Assigned_to"`
	Assigned_from  time.Time `json:"Assigned_from"`
	Assigned_by    string    `json:"Assigned_by"`
	Created_on     time.Time `json:"Created_on"`
	Created_by     string    `json:"Created_by"`
	Updated_on     time.Time `json:"Updated_on"`
	Updated_by     string    `json:"Updated_by"`
	OS_IP          string    `json:"OS_IP"`
	OS_User        int       `json:"OS_User"`
	Purpose        string    `json:"Purpose"`
	Cluster_ID     string    `json:"Cluster_ID"`
}

// list of Reserved Assets

func Reserved(w http.ResponseWriter, r *http.Request) {
	DB_conn := dbs.DB // connecting to database
	str := "SELECT Asset_ID,Manufacturer,BMC_IP,BMC_User,Asset_location,Reserved,COALESCE( assigned_to, 0 ) as assigned_to," +
		"Assigned_from,Assigned_by,Created_on,Created_by,Updated_on,Updated_by,OS_IP , COALESCE( os_user, 0 ) as os_user," +
		"Purpose,Cluster_ID FROM Asset where Reserved=true" //
	rows, err := DB_conn.Query(str)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"Status Code": "400 Bad Request", "Message": err})
		return

	}
	result := []Asset{} // creating slice
	var Asset_ID, Assigned_to, OS_User int
	var Manufacturer, OS_IP, BMC_IP, BMC_user, Asset_location, Assigned_by, Created_by, Updated_by, Cluster_ID, Purpose string
	var Created_on, Updated_on, Assigned_from time.Time
	var Reserved bool
	for rows.Next() {
		err := rows.Scan(&Asset_ID, &Manufacturer, &BMC_IP, &BMC_user, &Asset_location, &Reserved, &Assigned_to, &Assigned_from, &Assigned_by, &Created_on, &Created_by, &Updated_on, &Updated_by, &OS_IP, &OS_User, &Purpose, &Cluster_ID)

		if err != nil {
			fmt.Println(err)
			log.Printf("Failed to build content from sql rows: %v\n", err)

		}
		result = append(result, Asset{Asset_ID: Asset_ID, Manufacturer: Manufacturer, BMC_IP: BMC_IP, BMC_user: BMC_user, Asset_location: Asset_location, Reserved: Reserved, Assigned_to: Assigned_to, Assigned_by: Assigned_by, Assigned_from: Assigned_from, Created_by: Created_by, Created_on: Created_on, Updated_by: Updated_by, Updated_on: Updated_on, OS_IP: OS_IP, OS_User: OS_User, Cluster_ID: Cluster_ID, Purpose: Purpose})
	} // appending deatils to the result
	json.NewEncoder(w).Encode(map[string]interface{}{"ListAsset": result, "Status Code": "200 OK", "Message": "Listing all assets"})

}

// list of pools Assets

func Pool(w http.ResponseWriter, r *http.Request) {
	DB_conn := dbs.DB // connecting to database
	str := ("SELECT Asset_ID,Manufacturer,BMC_IP,BMC_User,Asset_location,Reserved,COALESCE( assigned_to, 0 ) as assigned_to,Assigned_from,Assigned_by,Created_on,Created_by,Updated_on,Updated_by,OS_IP ,COALESCE( os_user, 0 ) as os_user,Purpose,Cluster_ID FROM Asset where Reserved=false")

	rows, err := DB_conn.Query(str)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"Status Code": "400 Bad Request", "Message": err})
		return
		//panic(err.Error())
	}
	result := []Asset{} // creating slice

	var Asset_ID, Assigned_to, OS_User int
	var Manufacturer, OS_IP, BMC_IP, BMC_user, Asset_location, Assigned_by, Created_by, Updated_by, Cluster_ID, Purpose string
	var Created_on, Updated_on, Assigned_from time.Time
	var Reserved bool
	for rows.Next() {
		err := rows.Scan(&Asset_ID, &Manufacturer, &BMC_IP, &BMC_user, &Asset_location, &Reserved, &Assigned_to, &Assigned_from, &Assigned_by, &Created_on, &Created_by, &Updated_on, &Updated_by, &OS_IP, &OS_User, &Purpose, &Cluster_ID)

		if err != nil {
			fmt.Println(err)
			log.Printf("Failed to build content from sql rows: %v\n", err)

		}
		result = append(result, Asset{Asset_ID: Asset_ID, Manufacturer: Manufacturer, BMC_IP: BMC_IP, BMC_user: BMC_user, Asset_location: Asset_location, Reserved: Reserved, Assigned_to: Assigned_to, Assigned_by: Assigned_by, Assigned_from: Assigned_from, Created_by: Created_by, Created_on: Created_on, Updated_by: Updated_by, Updated_on: Updated_on, OS_IP: OS_IP, OS_User: OS_User, Cluster_ID: Cluster_ID, Purpose: Purpose})
	} // appending deatils to the result
	json.NewEncoder(w).Encode(map[string]interface{}{"ListAsset": result, "Status Code": "200 OK", "Message": "Listing all assets"})

}
