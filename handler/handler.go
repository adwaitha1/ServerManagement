package handler

import (
	"fmt"
	"log"
	"net/http"

	d "servermanagement.com/DASHBOARD"
	a "servermanagement.com/admin"
	login "servermanagement.com/login"
	s "servermanagement.com/server"
	usr "servermanagement.com/users"
)

func HandleFunc() {
	http.HandleFunc("/dashboard1", d.GetDashboard1)
	http.HandleFunc("/dashboard2", d.GetDashboard2)
	http.HandleFunc("/dashboard3", d.GetDashboard3)
	http.HandleFunc("/login", login.Login)
	fmt.Println("Server is hosted \n Server piechart : http://localhost:5000/dashboard1 \n Cluster piechart :  http://localhost:5000/dashboard2 \n Location bargraph :  http://localhost:5000/dashboard3 \n Login : http://localhost:5000/login")
	http.HandleFunc("/add_server", s.AddAsset)   //add new server details
	http.HandleFunc("/view_server", s.ListAsset) //view all servers
	http.HandleFunc("/assign_asset", s.Assign_asset)
	http.HandleFunc("/delete_asset", s.Delete_asset)

	fmt.Println("Assign Server http://localhost:5000/assign_asset \n Delete Server http://localhost:5000/delete_asset")

	http.HandleFunc("/release_server", s.Release)
	http.HandleFunc("/list_asset/Reserved", s.Reserved) //list all reserved assets
	http.HandleFunc("/list_asset/pool", s.Pool)         // list all pool assets
	//http.HandleFunc("/search", s.Search_server)                //view servers by id
	//http.HandleFunc("/search_server/bmc", s.Search_server_bmc) //need to give signature in single qoutes: ?ip='123.4.5.6'
	fmt.Println("\n Add_Server API : http://localhost:5000/add_server \n View_Server API :  http://localhost:5000/view_server \n Search_server API :  http://localhost:5000/search?id= \n Search by BMC_IP : http://localhost:5000/search_server/bmc")
	//http.HandleFunc("/view-user", a.ViewUsers)
	http.HandleFunc("/change_password", usr.ChangePassword)

	http.HandleFunc("/reset_password", a.ResetPassword)
	http.HandleFunc("/create_user", a.Create_User)
	http.HandleFunc("/View_Role", a.View_Role)
	http.HandleFunc("/list_users", a.View_user)
	http.HandleFunc("/user_id", a.Users_ID)

	fmt.Println("Change password : http://localhost:5000/change_password \n Reset password : http://localhost:5000/reset_password")

	fmt.Println(" \nCreate Users API : http://localhost:5000/create_user \n View Users API : http://localhost:5000/view-user")
	log.Fatal(http.ListenAndServe(":5000", nil))
	fmt.Println("Server is hosted")
}
