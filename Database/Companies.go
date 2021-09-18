package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllCompany() ([]Company, error) {
	Db, err := getDb()
	if err != nil {
		return nil, err
	}
	defer Db.Close()
	res, er := Db.Query("select * from companies")
	if er != nil {
		return nil, er
	}
	defer res.Close()
	var companies []Company
	for res.Next() {
		var company Company
		errr := res.Scan(&company.Id, &company.Name, &company.Email, &company.Address, &company.Phone_Number)
		if errr != nil {
			fmt.Println(errr)
			return nil, errr
		}
		companies = append(companies, company)
	}
	return companies, nil
}
func GetCompany(company_id int) (Company, error) {
	var company Company
	db, err := getDb()
	if err != nil {
		return company, err
	}
	qry := fmt.Sprintf("select * from companies where id=%d", company_id)
	fmt.Println(qry)
	res, err := db.Query(qry)
	if err != nil {
		return company, err
	}
	for res.Next() {
		err := res.Scan(&company.Id, &company.Name, &company.Email, &company.Address, &company.Phone_Number)
		if err != nil {
			return company, err
		}
	}
	return company, nil
}
func PostCompany(company Company) error {
	Db, err := getDb()
	if err != nil {
		return err
	}
	defer Db.Close()
	query := "insert into companies(name,email,address,phone_number) values(?,?,?,?)"
	fmt.Println(query)
	inst, er := Db.Prepare(query)
	if er != nil {
		fmt.Println("kijkj")
		return er
	}
	defer inst.Close()
	_, errr := inst.Exec(company.Name, company.Email, company.Address, company.Phone_Number)
	if errr != nil {
		fmt.Println("exec")
		return errr
	}
	return nil
}

func DeleteCompany(company_id int) (int64, error) {
	db, err := getDb()
	if err != nil {
		return -1, err
	}
	qry := fmt.Sprintf("delete from companies where id=%d", company_id)
	fmt.Println(qry)
	res, err := db.Prepare(qry)
	if err != nil {
		return -1, err
	}
	defer res.Close()
	rr, err := res.Exec()
	if err != nil {
		return -1, err
	}
	val, err := rr.RowsAffected()
	if err != nil {
		return -1, err
	}
	if val >= 1 {
		err := deleteapplication(db, company_id)
		if err != nil {
			return -1, err
		}
		err = deletejob(db, company_id)
		if err != nil {
			return -1, err
		}
		err = deleteUser(db, company_id)
		if err != nil {
			return -1, err
		}
	}
	return val, nil
}
func deletejob(db *sql.DB, company_id int) error {
	qry := fmt.Sprintf("delete from jobs where company_id=%d", company_id)
	fmt.Println(qry)
	res, err := db.Prepare(qry)
	if err != nil {
		return err
	}
	_, err = res.Exec()
	if err != nil {
		return err
	}
	return nil
}
func deleteapplication(db *sql.DB, company_id int) error {
	qry := fmt.Sprintf("delete from applications where job_id in (select id from jobs where company_id=%d)", company_id)
	fmt.Println(qry)
	res, err := db.Prepare(qry)
	if err != nil {
		return err
	}
	_, err = res.Exec()
	if err != nil {
		return err
	}
	return nil
}
func deleteUser(db *sql.DB, company_id int) error {
	qry := fmt.Sprintf("delete from users where company_id=%d", company_id)
	fmt.Println(qry)
	res, err := db.Prepare(qry)
	if err != nil {
		return err
	}
	_, err = res.Exec()
	if err != nil {
		return err
	}
	return nil
}
