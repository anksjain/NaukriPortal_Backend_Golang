package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsers() ([]User, error) {
	Db, err := getDb()
	if err != nil {
		return nil, err
	}
	defer Db.Close()
	res, er := Db.Query("select * from users")
	if er != nil {
		return nil, er
	}
	defer res.Close()
	var users []User
	for res.Next() {
		var user User
		errr := res.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email, &user.Password, &user.City, &user.Role, &user.Experience, &user.Company_id)
		if errr != nil {
			return nil, errr
		}
		users = append(users, user)
	}
	return users, nil
}
func GetAllUsersByType(types string) ([]User, error) {
	Db, err := getDb()
	if err != nil {
		return nil, err
	}
	defer Db.Close()
	qry := fmt.Sprintf(`select * from users where role="%s"`, types)
	fmt.Println(qry)
	res, er := Db.Query(qry)
	if er != nil {
		return nil, er
	}
	defer res.Close()
	var users []User
	for res.Next() {
		var user User
		errr := res.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email, &user.Password, &user.City, &user.Role, &user.Experience, &user.Company_id)
		if errr != nil {
			return nil, errr
		}
		users = append(users, user)
	}
	return users, nil
}
func GetAllUsersByJob(job_id int) ([]User, error) {
	Db, err := getDb()
	if err != nil {
		return nil, err
	}
	defer Db.Close()
	qry := fmt.Sprintf(`select * from users where id in (select user_id from applications where job_id=%d)`, job_id)
	fmt.Println(qry)
	res, er := Db.Query(qry)
	if er != nil {
		return nil, er
	}
	defer res.Close()
	var users []User
	for res.Next() {
		var user User
		errr := res.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email, &user.Password, &user.City, &user.Role, &user.Experience, &user.Company_id)
		if errr != nil {
			return nil, errr
		}
		users = append(users, user)
	}
	return users, nil
}
func CreateUser(user User) error {
	Db, err := getDb()
	if err != nil {
		return err
	}
	defer Db.Close()
	query := "insert into users(email,password,first_name,last_name,role,city,experience,company_id) values(?,?,?,?,?,?,?,?)"
	inst, er := Db.Prepare(query)
	if er != nil {
		return er
	}
	defer inst.Close()
	_, errr := inst.Exec(user.Email, user.Password, user.First_Name, user.Last_Name, user.Role, user.City, user.Experience, user.Company_id)
	if errr != nil {
		return errr
	}
	return nil
}
func VerfiyCredentials(username string, password string) (User, bool, error) {
	var user User
	Db, err := getDb()
	if err != nil {
		return user, false, err
	}
	defer Db.Close()
	qry := fmt.Sprintf(`select * from users where email="%s" and password="%s"`, username, password)
	res, er := Db.Query(qry)
	if er != nil {
		return user, false, er
	}
	defer res.Close()
	count := 0
	for res.Next() {
		count++
		errr := res.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email, &user.Password, &user.City, &user.Role, &user.Experience, &user.Company_id)
		if errr != nil {
			return user, false, errr
		}

	}
	if count == 1 {
		return user, true, nil
	}
	return user, false, nil

}
func IsUsernameExist(email string) (bool, error) {

	Db, err := getDb()
	if err != nil {
		return false, err
	}
	defer Db.Close()
	qry := fmt.Sprintf(`select * from users where email="%s"`, email)
	res, er := Db.Query(qry)
	if er != nil {
		return false, er
	}
	defer res.Close()
	return res.Next(), nil
}

func GetUser(id int) (User, error) {
	var user User
	Db, err := getDb()
	if err != nil {
		return user, err
	}
	defer Db.Close()
	qry := fmt.Sprintf(`select * from users where id=%d`, id)
	res, er := Db.Query(qry)
	if er != nil {
		return user, er
	}
	defer res.Close()
	for res.Next() {
		errr := res.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email, &user.Password, &user.City, &user.Role, &user.Experience, &user.Company_id)
		if errr != nil {
			return user, errr
		}
	}
	return user, nil
}
func DeleteUser(user_id int) (int64, error) {
	db, err := getDb()
	if err != nil {
		return -1, err
	}
	defer db.Close()
	qry := fmt.Sprintf(`delete from users where id=%d `, user_id)
	fmt.Println(qry)
	res, err := db.Prepare(qry)
	if err != nil {
		return -1, err
	}
	defer res.Close()
	r, err := res.Exec()
	if err != nil {
		return -1, err
	}
	val, err := r.RowsAffected()
	if err != nil {
		return -1, err
	}
	if val >= 1 {
		// if user is recruiter then we should delete all applications related to that recruiter
		qry = fmt.Sprintf(`delete from applications where job_id in (select id from jobs where user_id=%d) `, user_id)
		fmt.Println(qry)
		res, err := db.Prepare(qry)
		if err != nil {
			return -1, err
		}
		_, err = res.Exec()
		if err != nil {
			return -1, err
		}
		// if user is recruiter
		qry := fmt.Sprintf(`delete from jobs where user_id=%d `, user_id)
		fmt.Println(qry)
		res, err = db.Prepare(qry)
		if err != nil {
			return -1, err
		}
		_, err = res.Exec()
		if err != nil {
			return -1, err
		}

		// if user is candidate
		qry = fmt.Sprintf(`delete from applications where user_id=%d `, user_id)
		fmt.Println(qry)
		res, err = db.Prepare(qry)
		if err != nil {
			return -1, err
		}
		_, err = res.Exec()
		if err != nil {
			return -1, err
		}
	}
	return val, nil
}
func UpdateUser(user User) (int64, error) {
	db, err := getDb()
	if err != nil {
		return -1, err
	}
	qry := fmt.Sprintf(`update users set first_name="%s" ,last_name="%s",password="%s" ,city="%s" ,experience=%d where id=%d`, user.First_Name, user.Last_Name, user.Password, user.City, user.Experience, user.Id)
	fmt.Println(qry)
	res, err := db.Prepare(qry)
	if err != nil {
		return -1, err
	}
	r, err := res.Exec()
	if err != nil {
		return -1, err
	}
	val, err := r.RowsAffected()
	if err != nil {
		return -1, err
	}
	return val, nil
}
