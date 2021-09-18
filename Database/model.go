package database

var DBSession = make(map[string]string)

type User struct {
	Id         int    `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	City       string `json:"city"`
	Role       string `json:"role"`
	Experience int    `json:"experience"`
	Company_id int    `json:"company_id"`
}

type Job struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Company_id  int    `json:"company_id"`
	User_id     int    `json:"user_id"`
}

type Company struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	Phone_Number string `json:"phone_number"`
}

type Users struct {
	Id         int `gorm:"autoIncrement"`
	First_Name string
	Last_Name  string
	Email      string
	Password   string
	City       string
	Role       string
	Experience int
	Company_id int
}

type Jobs struct {
	Id          int `gorm:"autoIncrement"`
	Title       string
	Description string
	Company_id  int
	User_id     int
}

type Companies struct {
	Id           int `gorm:"autoIncrement"`
	Name         string
	Email        string
	Address      string
	Phone_Number string
}
type Applications struct {
	Id      int `gorm:"autoIncrement"`
	Job_id  int
	User_id int
}
