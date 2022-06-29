package connection

import (
	"database/sql"
	"fmt"

	// library for connect postgresql
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	// DB variable for connection DB postgresql
	DB *sql.DB
)

// Connects to database and creates schema, returns connection object
func Connect() *sql.DB {

	fmt.Println("Connecting to database ... ")
	viper.SetConfigName("configdb") // name of config file (without extension)
	viper.SetConfigType("json")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")        // Look for config in the working directory
	fmt.Println("Reading configuration from json file ... ")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error reading json config file: %w \n", err))
	}

	user := viper.GetString("user") // viper is to read both from configuration file and environment variables.
	password := viper.GetString("password")
	host := viper.GetString("host")
	port := viper.GetString("port")
	dbname := viper.GetString("dbname")

	psqlInfo := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " sslmode=disable database=" + dbname
	fmt.Printf(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("Check your config file, Database not connected ...", err)
		return nil
	}
	fmt.Printf("DB connected successfully !! \n")
	DB = db
	return DB
}
