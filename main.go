package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_user     = "root"
	db_password = "toor"
	db_address  = "127.0.0.1"
	db_db       = "test"
)

type Person struct {
	ID       int
	Name     string
	Age      int
	Location string
}

func main() {
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_password, db_address, db_db)
	fmt.Printf(s)
	db, err := sql.Open("mysql", s)

	defer db.Close() //When the main function ends, the connection will be closed

	if err != nil {
		log.Panic(err)
	}

	/*err = insertData(db)
	if err != nil {
		log.Fatal(err)
	}*/

	people, err := getAllData(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(people)
	fmt.Printf("We have %v people in our db", len(people))
	people, err = getAllAboveAge(db, 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(people)
	fmt.Printf("We have %v people which are older than 30", len(people))

	err = deleteAllAboveAge(db, 30)
	if err != nil {
		fmt.Println(err)
	}

	people, err = getAllAboveAge(db, 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(people)
	fmt.Printf("We have %v people which are older than 30", len(people))

	err = updatePersonAge(db, "Marianne C Reagan", 21)
	people, err = getAllData(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(people)
	fmt.Printf("We have %v people in our db", len(people))

	if err != nil {
		log.Fatal(err)
	}

}

func insertData(db *sql.DB) error {
	people := GetData()
	for _, person := range people {
		q := "INSERT INTO `person` (name, age, location) VALUES (?, ?, ?)"
		insert, err := db.Prepare(q) //Prepare a command
		defer insert.Close()

		if err != nil {
			return err
		}

		_, err = insert.Exec(person.Name, person.Age, person.Location) //Execute a command

		if err != nil {
			return err
		}

	}

	return nil
}

func getAllData(db *sql.DB) (people []Person, err error) {
	resp, err := db.Query("SELECT * from `person`") //Select data from database
	defer resp.Close()

	if err != nil {
		return people, err
	}

	for resp.Next() {
		var pPerson Person
		err = resp.Scan(&pPerson.ID, &pPerson.Name, &pPerson.Age, &pPerson.Location)
		if err != nil {
			return people, err
		}

		people = append(people, pPerson)
	}

	return people, nil
}

func getAllAboveAge(db *sql.DB, age int) (people []Person, err error) {
	q := "SELECT * FROM `person` WHERE `age` > ?"
	resp, err := db.Query(q, age) //Select data from database
	defer resp.Close()

	if err != nil {
		return people, err
	}

	for resp.Next() {
		var pPerson Person
		err = resp.Scan(&pPerson.ID, &pPerson.Name, &pPerson.Age, &pPerson.Location)
		if err != nil {
			return people, err
		}

		people = append(people, pPerson)
	}

	return people, nil
}

func deleteAllAboveAge(db *sql.DB, age int) error {
	q := "DELETE FROM `person` WHERE `age` > ?" //it will automaticaly set the age
	drop, err := db.Prepare(q)
	defer drop.Close()
	if err != nil {
		return err
	}

	_, err = drop.Exec(age)
	if err != nil {
		return err
	}

	return nil
}

func updatePersonAge(db *sql.DB, name string, age int) error {
	q := "UPDATE `person` SET `age` = ? WHERE `name` LIKE ?"

	update, err := db.Prepare(q)
	defer update.Close()

	if err != nil {
		return err
	}

	_, err = update.Exec(age, name)
	if err != nil {
		return err
	}

	return nil
}
