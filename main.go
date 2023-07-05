package main

import (
	"fmt"

	// "io/ioutil"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

/*
	func hello(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello\n")
	}

	func headers(w http.ResponseWriter, req *http.Request) {
		for name, headers := range req.Header {
			for _, h := range headers {
				fmt.Fprintf(w, "%v: %v\n", name, h)
			}
		}
	}

	func userDetails(w http.ResponseWriter, req *http.Request) {
		jsonData, err := ioutil.ReadFile("storage/user_data.json")
		if err != nil {
			http.Error(w, "Failed to read JSON file", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonData)
		if err != nil {
			fmt.Println("Failed to write JSON response:", err)
		}
	}

	func userDetailsP(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		jsonData, err := ioutil.ReadFile("storage/user_data.json")
		if err != nil {
			http.Error(w, "Failed to read JSON file", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonData)
		if err != nil {
			fmt.Println("Failed to write JSON response:", err)
		}
	}
*/
type Customer struct {
	ID          int    `json:"customer_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

func InsertCustomer(customer Customer) error {
	// Set up the database connection
	db, err := sql.Open("mysql", "root:Jaimatadi382#@tcp(Pradeeps-iMac.local:3306)/Shrimad_Bank")
	if err != nil {
		log.Println("DB could not be opened error is ", err)
		return err
	}
	defer db.Close()

	// Prepare the insert statement
	stmt, err := db.Prepare("INSERT INTO Customers (customer_id, first_name, last_name, date_of_birth, address, phone_number, email, username, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("prepare statement failed ", err)
		return err
	}
	defer stmt.Close()

	// Execute the insert statement
	_, err = stmt.Exec(customer.ID, customer.FirstName, customer.LastName, customer.DateOfBirth, customer.Address, customer.PhoneNumber, customer.Email, customer.Username, customer.Password)
	if err != nil {
		log.Println("execute statement failed ", err)
		return err
	}

	fmt.Println("Customer inserted successfully.")
	return nil
}

func main() {
	r := gin.Default()

	r.POST("/customers", func(c *gin.Context) {
		var customer Customer
		if err := c.ShouldBindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := InsertCustomer(customer)
		if err != nil {
			log.Println("Error inserting customer:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert customer"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Customer inserted successfully"})
	})

	if err := r.Run(":8090"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// func main() {
// 	http.HandleFunc("/hello", hello)
// 	http.HandleFunc("/headers", headers)
// 	http.HandleFunc("/getuserdetails", userDetails)
// 	http.HandleFunc("/getuserdetailsp", userDetailsP)

// 	http.ListenAndServe(":8090", nil)
// }
