package command

import "github.com/dg222599/go-redis-db/db"

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

		case SET:
			dbOP.db.Set(dbIdx,cmd.Key,cmd.Value)
			return dbIdx,"OK"
		case GET:
			return dbIdx,dbOP.db.Get(dbIdx,cmd.Key)
		case DEL:
			return dbIdx,dbOP.db.Del(dbIdx,cmd.Key)
		case INCR:
			v:=

			


	  }
}

