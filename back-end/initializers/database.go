package initializers

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

// var DB *gorm.DB
var db *sql.DB

func ConnectToDB() {
	rootCertPool := x509.NewCertPool()
	pem, _ := os.ReadFile("C:/Users/zombi/OneDrive/go/src/github.com/mattwirn/Bear-ly-Breaking-Even/back-end/BaltimoreCyberTrustRoot.crt.pem")

	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}
	mysql.RegisterTLSConfig("custom", &tls.Config{RootCAs: rootCertPool})

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true&tls=custom", "bbeAdmin@bbe-mysql-server", "bigB3autifulBe4rs!", "bbe-mysql-server.mysql.database.azure.com", "quickstartdb")
	var err error
	db, err = sql.Open("mysql", connectionString)

	/*
		dsn := "bbeAdmin:bigB3autifulBe4rs!@tcp(bbe-mysql-server.mysql.database.azure.com:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	*/
	if err != nil {
		log.Fatal("Failed to connect to database")
	} else {
		fmt.Println("Connected to database")
	}
}
