package initializers

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"

	gMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB2 *gorm.DB

// var DB *sql.DB

func ConnectToDB_OLD() {

	rootCertPool := x509.NewCertPool()
	pem, _ := os.ReadFile("BaltimoreCyberTrustRoot.crt.pem")

	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}
	mysql.RegisterTLSConfig("custom", &tls.Config{RootCAs: rootCertPool})

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true&tls=custom", "bbeAdmin@bbe-mysql-server", "bigB3autifulBe4rs!", "bbe-mysql-server.mysql.database.azure.com", "quickstartdb")
	var err error
	// DB, err = sql.Open("mysql", connectionString)

	// dsn := "bbeAdmin:bigB3autifulBe4rs!@tcp(bbe-mysql-server.mysql.database.azure.com:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	DB2, err = gorm.Open(gMysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	} else {
		fmt.Println("Connected to database")
	}

}
