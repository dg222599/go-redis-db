package db

import (
	"fmt"
	"strconv"
)
type DB struct {
	DBCount int
	DBStorage map[int]map[string]interface{}
}

func CreateDB(dbCountstr string) DB {
	fmt.Println("This is inside the DB Package!!")
	dbCount,err := strconv.Atoi(dbCountstr)

	if err!=nil{
		dbCount = 16
	}

	temp := make(map[int]map[string]interface{})

	for i:=0 ; i<dbCount ; i++ {
		temp[i] = make(map[string]interface{})

	}

	return DB{
		DBCount: dbCount,
		DBStorage: temp,
	}
}

func (db DB) Select(dbIdxStr string) (int , error) {
	dbIdx,err := strconv.Atoi(dbIdxStr)
	if err!=nil{
		return 0,fmt.Errorf("(error) - DB number is not an integer")
	}
	if dbIdx < 0 || dbIdx > db.DBCount - 1 {
		return 0,fmt.Errorf("(error) - DB Index is out of range")
	}

	return dbIdx,nil
}

func (db DB) Set(dbIdx int,key string ,value interface{}) {
	dbInstance:= db.DBStorage[dbIdx]
	dbInstance[key] = value
} 

func (db DB) Get(dbIdx int,key string) interface{}{
	dbInstance:=db.DBStorage[dbIdx]
	value,ok:=dbInstance[key]
	if !ok{
		return nil
	}
	return value
}

func (db DB) Del(dbIdx int,key string) interface{}{
	dbInstance:=db.DBStorage[dbIdx]
	_,ok:=dbInstance[key]
	if !ok{
		return 0
	}
	delete(dbInstance,key)
	return 1
}

func (db DB) Show(dbIdx int) <-chan string{
	resultChan := make(chan string)
	
	go func(){
		for key,value := range db.DBStorage[dbIdx] {
			resultChan <- fmt.Sprintf("`%s` -> `%s`" , key,value)
		}
		close(resultChan)
	}()

	return resultChan
}