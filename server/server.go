//viewing server as infra admin
package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	dbs "servermanagement.com/connection"
	"servermanagement.com/cors"
)

type asset struct {
	Asset_Id       int       `json:"Asset_Id"`
	Manufacturer   string    `json:"Manufacturer"`
	BMC_IP         string    `json:"BMC_IP"`
	BMC_User       string    `json:"BMC_USER"`
	BMC_Password   string    `json:"BMC_Password"`
	Asset_location string    `json:"Asset_location"`
	Reserved       bool      `json:"Reserved"`
	Assigned_by    string    `json:"Assigned_by"`
	Assigned_to    int       `json:"Assigned_to"`
	Assigned_from  time.Time `json:"Assigned_from"`
	Created_on     time.Time `json:"Created_on"`
	Created_by     string    `json:"Created_by"`
	Updated_on     time.Time `json:"Updated_on"`
	Updated_by     string    `json:"Updated_by"`
	OS_IP          string    `json:"OS_IP"`
	OS_User        int       `json:"OS_User"`
	OS_Password    string    `json:"OS_Password"`
	Purpose        string    `json:"Purpose"`
	Cluster_Id     string    `json:"Cluster_ID"`
	Delete         int       `json:"Delete"`
	Status         bool      `json:"Status"`
}

//------------------------------------------------add asset(creating asset)---------------------------------------------------------------------

func AddAsset(write http.ResponseWriter, request *http.Request) {
	cors.SetupCORS(&write, request)
	DB_conn := dbs.DB
	var assets asset
	err := json.NewDecoder(request.Body).Decode(&assets)
	if err != nil {
		json.NewEncoder(write).Encode(map[string]interface{}{"status": "400 Bad Request", "Message": err})
		return
	}
	addStatement := `INSERT INTO asset (Asset_ID,Manufacturer,BMC_IP,BMC_User,BMC_Password,Asset_location,Created_on,Created_by,OS_IP ,OS_User ,OS_Password,Purpose,Cluster_ID,Delete) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,'1')`
	_, err = DB_conn.Exec(addStatement, assets.Asset_Id, assets.Manufacturer, assets.BMC_IP, assets.BMC_User, assets.BMC_Password, assets.Asset_location, assets.Created_on, assets.Created_by, assets.OS_IP, assets.OS_User, assets.OS_Password, assets.Purpose, assets.Cluster_Id)
	if err != nil {
		json.NewEncoder(write).Encode(map[string]interface{}{"status": "400 Bad Request", "Message": err})
		return
	}

	json.NewEncoder(write).Encode(map[string]interface{}{"Status Code": "200 OK", "Message": "Recorded sucessfully"})
}

// list Assets

func ListAsset(w http.ResponseWriter, r *http.Request) {
	cors.SetupCORS(&w, r)
	DB_conn := dbs.DB // connecting to database
	{
		str := "SELECT asset_id,manufacturer,bmc_ip,bmc_user,bmc_password,asset_location,reserved," +
			"COALESCE( assigned_to, 0 ) as assigned_to, assigned_by,assigned_from , created_by , created_on,updated_by," +
			"updated_on,os_ip , COALESCE( os_user, 0 ) as os_user , os_password,cluster_id ,purpose  FROM Asset" //

		rows, err := DB_conn.Query(str)
		if err != nil {
			//fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Status Code": "400 Bad Request", "Message": err})
			return
			//panic(err.Error())
		}
		result := []asset{} // creating slice
		var Asset_ID, OS_User, Assigned_to int
		var Manufacturer, OS_IP, OS_Password, BMC_IP, BMC_user, BMC_password, Asset_location, Assigned_by, Created_by, Updated_by, Cluster_ID, Purpose string
		var Created_on, Updated_on, Assigned_from time.Time
		var Reserved bool
		for rows.Next() {
			err := rows.Scan(&Asset_ID, &Manufacturer, &BMC_IP, &BMC_user, &BMC_password, &Asset_location, &Reserved, &Assigned_to, &Assigned_by, &Assigned_from, &Created_by, &Created_on, &Updated_by, &Updated_on, &OS_IP, &OS_User, &OS_Password, &Cluster_ID, &Purpose)

			if err != nil {
				fmt.Println(err)
				log.Printf("Failed to build content from sql rows: %v\n", err)

			}
			result = append(result, asset{Asset_Id: Asset_ID, Manufacturer: Manufacturer, BMC_IP: BMC_IP, BMC_User: BMC_user, BMC_Password: BMC_password, Asset_location: Asset_location, Reserved: Reserved, Assigned_to: Assigned_to, Assigned_by: Assigned_by, Assigned_from: Assigned_from, Created_by: Created_by, Created_on: Created_on, Updated_by: Updated_by, Updated_on: Updated_on, OS_IP: OS_IP, OS_User: OS_User, OS_Password: OS_Password, Cluster_Id: Cluster_ID, Purpose: Purpose})
		} // appending details to the result

		json.NewEncoder(w).Encode(map[string]interface{}{"ListAsset": result, "Status Code": "200 OK", "Message": "Listing all assets"})

	}
}
func Assign_asset(w http.ResponseWriter, r *http.Request) {
	cors.SetupCORS(&w, r)
	DB_conn := dbs.DB // connecting to database
	var p asset
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 Bad Request"})
		return
	}
	_, err = DB_conn.Exec("UPDATE asset SET assigned_to=$1, reserved = 't' WHERE asset_id=$2;", p.Assigned_to, p.Asset_Id)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 Bad Request"})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Server Assigned!", "Status Code": "200 OK"})

	Reserved := p.Reserved
	if !Reserved {
		_, err := DB_conn.Query(`INSERT into Historic_details (Asset_ID,Assigned_to,Assigned_from,Updated_ON,Updated_by,Remarks)
		SELECT Asset_ID,Assigned_to,Assigned_from,Updated_ON,Updated_by,'Server_assigned' FROM Asset WHERE asset_id=$2`, p.Asset_Id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, "Record Updated!")
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "No update required"})

	}
}

//---------------Delete Server(Updating delete and reserved column in asset table)--------------
func Delete_asset(w http.ResponseWriter, r *http.Request) {
	cors.SetupCORS(&w, r)
	DB_conn := dbs.DB // connecting to database
	var p asset
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 Bad Request"})
		return
	}
	_, err = DB_conn.Exec("UPDATE asset SET Delete='1', Reserved = 'f' , Assigned_to = null WHERE Asset_Id=$1", p.Asset_Id)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 Bad Request"})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Deleted Server!", "Status Code": "200 OK"})
	row := DB_conn.QueryRow("SELECT Delete from asset where Asset_Id=$1;", p.Asset_Id)
	var del int
	err1 := row.Scan(&del)
	if err1 != nil {
		log.Fatal(err1)
	}
	if !p.Reserved && del == 1 {
		_, err := DB_conn.Query(`INSERT into Historic_details (Asset_ID,Assigned_to,Assigned_from,Updated_ON,Updated_by,Remarks) 
		SELECT Asset_ID,Assigned_to,Assigned_from,Updated_ON,Updated_by,'Server Deleted' FROM Asset `)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, "Record Updated!")
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "No update required"})
	}
}

//Release server (updating Reserve table)

func Release(w http.ResponseWriter, r *http.Request) {
	cors.SetupCORS(&w, r)
	DB_conn := dbs.DB // connecting to database

	var p asset //

	err := json.NewDecoder(r.Body).Decode(&p)
	//Asset_ID := 0
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"Status Code": "400 Bad Request", "Message": err})
		return
	}

	//Asset_ID++
	a := p.Asset_Id
	if p.Asset_Id == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{"Status Code": "400 Bad Request", "Message": err})
		return
	} else if a != p.Asset_Id {
		json.NewEncoder(w).Encode(map[string]interface{}{"Status Code": "400 Bad Request", "Message": err})
		return

	} else {

		_, err = DB_conn.Exec(`UPDATE Asset SET Reserved='false',Assigned_to=null where Asset_ID=$1;`, p.Asset_Id) // query for updating

		if err != nil {
			fmt.Println(err)
			// json.NewEncoder(w).Encode(map[string]interface{}{"Status Code": "400 Bad Request", "Message": err})
			// return
		}
		fmt.Fprintf(w, "Record Updated!")
	}

}
