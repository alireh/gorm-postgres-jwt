package model

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

var secretKey = envVariable("SECRET")

// ############ Application DB ###########################
func SetDBClient() {
	var (
		host     = envVariable("DB_HOST")
		port     = envVariable("DB_PORT")
		user     = envVariable("DB_USER")
		dbname   = envVariable("DB_NAME")
		password = envVariable("DB_PASSWORD")
	)

	dns := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	// dsn1 := url.URL{
	//     User:     url.UserPassword(user, password),
	//     Scheme:   "postgres",
	//     Host:     fmt.Sprintf("%s:%d", host, port),
	//     Path:     dbname,
	//     RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	// }

	// dns := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	//DB, err = gorm.Open("postgres", dns)
	DB.AutoMigrate(User{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connection to the database is successful")

}
func envVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// ################ Application Model ####################
type User struct {
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (u *User) GeneratePasswordHarsh() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(bytes)
	return err
}
func (u *User) CheckPasswordHarsh(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
