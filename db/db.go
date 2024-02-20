package db

import (
	"fmt"
)

type DB struct {
	dbName string
	dbStorage map[string]interface{}
}

func CreateDB(dbname string) (*DB,error){
	fmt.Println("This is inside the DB Package!!")



	newDBInstance := &DB{
		dbName: dbname,
		dbStorage:make(map[string]interface{}),
	}

	return newDBInstance,nil
}