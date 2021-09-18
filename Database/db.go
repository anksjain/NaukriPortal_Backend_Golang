package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	return os.Getenv(key)
}
func getDb() (*sql.DB, error) {
	var myDataSource = fmt.Sprintf("%s:%s@tcp(%s)/%s", GetEnv("DB_USERNAME"), GetEnv("DB_PASSWORD"), GetEnv("DB_ADDRESS_PORT"), GetEnv("DB_DATABASE"))
	Db, err := sql.Open("mysql", myDataSource)
	if err != nil {
		return nil, err
	}
	return Db, nil
}
func MigrationSeed() {
	var myDataSource = fmt.Sprintf("%s:%s@tcp(%s)/%s", GetEnv("DB_USERNAME"), GetEnv("DB_PASSWORD"), GetEnv("DB_ADDRESS_PORT"), GetEnv("DB_DATABASE"))
	db, err := gorm.Open(mysql.Open(myDataSource), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}
	err = db.Migrator().CreateTable(&Users{})
	if err == nil {
		var users []Users
		users = append(users, Users{First_Name: "Super", Last_Name: "Admin", Email: "super@admin.com", Password: "superadmin", Role: "Admin", City: "Albora", Company_id: 0, Experience: 0})
		users = append(users, Users{First_Name: "Recru", Last_Name: "First", Email: "recruiter1@sb.com", Password: "recruiter1", Role: "Recruiter", City: "enji", Company_id: 1, Experience: 0})
		users = append(users, Users{First_Name: "Recru", Last_Name: "Second", Email: "recruiter2@sb.com", Password: "recruiter2", Role: "Recruiter", City: "simon", Company_id: 2, Experience: 0})
		users = append(users, Users{First_Name: "Candi", Last_Name: "First", Email: "candidate1@sb.com", Password: "candidate1", Role: "Candidate", City: "Symphiny", Company_id: 0, Experience: 3})
		users = append(users, Users{First_Name: "Candi", Last_Name: "Two", Email: "candidate2@sb.com", Password: "candidate2", Role: "Candidate", City: "cuzal", Company_id: 0, Experience: 5})
		db.Create(&users)
	}
	err = db.Migrator().CreateTable(&Companies{})
	if err == nil {
		var comapnies []Companies
		comapnies = append(comapnies, Companies{Name: "Square Boat", Email: "Squareboat.org", Address: "Near Mega Mall Gurgaon", Phone_Number: "789456130"})
		comapnies = append(comapnies, Companies{Name: "Google", Email: "Google.com", Address: "Near Jupiter", Phone_Number: "1478523692"})
		comapnies = append(comapnies, Companies{Name: "ASPER Boat", Email: "Asperboat.org", Address: "Near Venus", Phone_Number: "1234567890"})
		db.Create(&comapnies)
	}
	db.Migrator().CreateTable(&Jobs{})
	db.Migrator().CreateTable(&Applications{})

}
