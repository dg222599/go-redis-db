package command

import (
	"fmt"
)

const (
	SET = "SET"
	GET = "GET"
	ALTER = "ALTER"
	INCR = "INCR"
	INCRBY = "INCRBY"
	DEL = "DEL"
	SHOW = "SHOW"
	MULTI = "MULTI"
	EXEC = "EXEC"
	DISCARD = "DISCARD"
	SELECT = "SELECT"
	HELP = "HELP"
	
)


type Command struct {
	Name  string
	Key   string
	Value interface{}
}



func NewCommand(name string,args ...interface{}) Command {
      var key string
	  var value interface{}

	  if len(args) > 0 {
		key  = fmt.Sprintf("%v",args[0])
	  }

	  if len(args) > 1 {
		value = args[1]
	  }

	  return Command{
		Name : name,
		Key : key,
		Value : value,
	  }

}

func (cmd Command) IsEndCommand() bool {
	
	   if cmd.Name == EXEC || cmd.Name == DISCARD {
		return true
	   }

	   return false
}

func (cmd Command) ValidateCommand() (bool,error){
    switch cmd.Name {
	case SET,INCRBY:
		tempCmd:="set"
		if cmd.Name == "INCRBY" {
			tempCmd = "incrby"

		}
		if cmd.Value == nil {
			return false,fmt.Errorf("(error) ERR wrong number of arguments , looks like the value is missing for Command -> %s",tempCmd)

		}
		return true,nil
	case DEL,GET,INCR:
		tempCmd:="del"
		if cmd.Name == "GET" {
			tempCmd = "get"
		} else if cmd.Name == "INCR" {
			tempCmd = "incr"
		}

		if cmd.Key == "" {
			return false,fmt.Errorf("(error)-ERR in the cmd-> %s...looks like the key is missing",tempCmd)
		}
		return true,nil
	case SELECT:
		tempCmd:="select"
		if cmd.Key == "" {
			return false,fmt.Errorf("(error)-ERR in the cmd-> %s...looks like the key is missing",tempCmd)
		}
		return true,nil
	case MULTI,EXEC,DISCARD,SHOW,HELP:
		return true,nil
	}

	//handling unknown commands
	unknownArgs := ""

	if cmd.Key != "" {
		unknownArgs = fmt.Sprintf("`%s`," ,cmd.Key)
		
		if cmd.Value != nil {
			unknownArgs = fmt.Sprintf("`%s` , `%v` " , cmd.Key , cmd.Value)

		}
	}

	return false,fmt.Errorf("(error) ERR unknown command `%s` entered with args as `%s` ...please enter correct command " , cmd.Name,unknownArgs)


}


