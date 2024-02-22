package db

import (
	"fmt"
)
type DB struct {
	DBName    string
	DBStorage map[string]interface{}
}

func CreateDB(dbname string) (*DB, error) {
	fmt.Println("This is inside the DB Package!!")
	newDBInstance := &DB{
		DBName:    dbname,
		DBStorage: map[string]interface{}{},
	}

	return newDBInstance, nil
}

func (db *DB) Show() {
	fmt.Println("Here is the whole DB!!")
	fmt.Println("DB Name is-->",db.DBName)
	fmt.Println("DB Storage is-->",db.DBStorage)
}

func (db *DB) Set(key string,value string) {
	//setted the DB key
	fmt.Println(db)
	db.DBStorage[key] = value
} 

func (db *DB) Get(key string) (interface{}) {

	value,ok := db.DBStorage[key]
	if !ok{
		 return nil
	}

	return value
}

func (db *DB) Delete(key string) {
	delete(db.DBStorage,key)
	
}