package db

import (
	"errors"
	"fmt"
	"reflect"
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
	fmt.Printf("Here is the whole DB!!\n\n")
	fmt.Println("DB Name is-->",db.DBName)
	fmt.Println("DB Storage is-->",db.DBStorage)
}

func (db *DB) Set(key string,value interface{}) {
	db.DBStorage[key] = value
} 

func (db *DB) Get(key string) (interface{}) {

	value,ok := db.DBStorage[key]
	if !ok{
		 return nil
	}

	return value
}

func (db *DB) Delete(key string) (int) {
	_,ok := db.DBStorage[key]
	if !ok{
		return 0
	}
	delete(db.DBStorage,key)
	return 1

}

func (db *DB) Increment(key string ,byCount int) (int,error){

	
	value,ok := db.DBStorage[key]
	
	if !ok{
		return 0,errors.New("key does not exist")
	}
    ok = false
	
	switch temp:=value.(type) {
		case int:
			ok = true
			fmt.Println(temp)
		default:
			ok = false
			
	}
	
	// value,_  = value.(int)
	// value = string(value)
	// intValue,ok := strconv.Atoi(value)
	if !ok{
		fmt.Println(value)
		fmt.Println(reflect.TypeOf(value))
		return 0,errors.New("value of the key is not of Int type , can not increment")
	}

	intValue,_ := value.(int)

	
	intValue+=byCount

	db.DBStorage[key] = intValue

	return intValue,nil
}