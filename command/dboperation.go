package command

import (
	"fmt"
	"strconv"

	"github.com/dg222599/go-redis-db/db"
)

type DBOperation struct {
	db db.DB
	runningMultiTxn bool
	cmds []Command
}

func NewDBOperation(db db.DB) DBOperation {
	return DBOperation{db : db , runningMultiTxn:false , cmds : []Command{}}
}

func (dbOP *DBOperation) enqueue(cmd Command) {

	dbOP.cmds = append(dbOP.cmds, cmd)
}

func (dbOP *DBOperation) Execute(dbIdx int,cmd Command) (int,interface{}) {
    
	  _,err := cmd.ValidateCommand()
	  if err!=nil{
		return dbIdx, err.Error()
	  }

	  if dbOP.runningMultiTxn && !cmd.IsEndCommand() {
		dbOP.enqueue(cmd)
		return dbIdx, "QUEUED"
	  }

	  switch cmd.Name {
		case SELECT :
			dbIdx,err  = dbOP.db.Select(cmd.Key)
			if err!=nil{
				return dbIdx , err.Error()
			}
			return dbIdx , "OK"
		case MULTI:
			dbOP.runningMultiTxn = true
			return dbIdx,"OK"
		case DISCARD:
			dbOP.runningMultiTxn = false
			dbOP.cmds = nil
			return dbIdx,"OK"
		case EXEC:
			dbOP.runningMultiTxn = false
			return dbIdx,dbOP.executeMulti(dbIdx)
		case COMPACT:
			var results []interface{}
			
			for keyValPair := range dbOP.db.Show(dbIdx) {
				results = append(results, fmt.Sprintf("SET %s ",keyValPair))
			}
			return dbIdx,results

		case SET:
			dbOP.db.Set(dbIdx,cmd.Key,cmd.Value)
			return dbIdx,"OK"
		case GET:
			return dbIdx,dbOP.db.Get(dbIdx,cmd.Key)
		case DEL:
			return dbIdx,dbOP.db.Del(dbIdx,cmd.Key)
		case INCR:
			value:=dbOP.db.Get(dbIdx,cmd.Key)
			if value == nil {
				newValue := "1"
				dbOP.db.Set(dbIdx,cmd.Key,newValue)
				return dbIdx,newValue

			}

			currValue,err := strconv.Atoi(value.(string))
			if err!=nil{
				return dbIdx , fmt.Errorf("(error) value is not an intger or out of range")

			}

			newValue := fmt.Sprintf("%v",currValue+1)
			dbOP.db.Set(dbIdx,cmd.Key , newValue)
			return dbIdx,newValue
		case INCRBY:
			value:=dbOP.db.Get(dbIdx,cmd.Key)
			if value == nil {
				newValue:=cmd.Value
				dbOP.db.Set(dbIdx,cmd.Key,newValue)
				return dbIdx,newValue
			}

			currValue,err := strconv.Atoi(value.(string))
			if err!=nil{
				return dbIdx,fmt.Errorf("(error) - current value is not an intger or out of range")
			}

			newValue,err := strconv.Atoi(cmd.Value.(string))
			if err!=nil{
				return dbIdx,fmt.Errorf("(error)  - new value is not an intger or out of range")
			}

			finalValue := fmt.Sprintf("%v",currValue + newValue)
			
			dbOP.db.Set(dbIdx,cmd.Key,finalValue)
			return dbIdx,finalValue
		case HELP,EXIT:
			return dbIdx,""
		}

		return dbIdx, fmt.Errorf("(error) - Error unknown command  '%s' ",cmd.Key)
}

func (dbOP *DBOperation) executeMulti(dbIdx int) interface{} {
	 		var results []interface{}

			for _,cmd := range dbOP.cmds {
				_,result := dbOP.Execute(dbIdx,cmd)
				results = append(results, result)
			}

			dbOP.cmds = nil

			return results
}



