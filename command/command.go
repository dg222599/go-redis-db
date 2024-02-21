package command

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"strings"

	"github.com/dg222599/go-redis-db/db"
)

const (
	SET = "SET"
	GET = "GET"
	ALTER = "ALTER"
	INCR = "INCR"
	DEL = "DEL"
	SHOW = "SHOW"
	
)
type Command struct {
	Name  string
	Key   string
	Value interface{}
}

func InitCommand() Command {
	initialCmd := Command{
		Name:  "command_name",
		Key:   "key_name",
		Value: 1,
	}

	return initialCmd
}

func ValidateCommand(cmdLine string) (bool,[]string,error){
    cmdArgs := strings.Split(cmdLine," ")

	if len(cmdArgs) < 1 {
		fmt.Println("You have entered empty command...please enter again ")
		return false,[]string{},nil
	}

    // to add proper validation for all the command types

	return true,cmdArgs,nil
}

func  HandleCommand(dbInstance *db.DB) {
	fmt.Println("Enter the Command or press CTRL+C to exit")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err.Error())
	}

	validationStatus,cmdArgs, _ := ValidateCommand(line)



	if !validationStatus{
		fmt.Println("InValid Command format ...pls enter again")
		HandleCommand(dbInstance)
	} else {

		fmt.Println("Congrats the command is verified!!~")
		fmt.Println(cmdArgs)
		switch cmdArgs[0] {
		case SET:
			fmt.Println("It is a SET command")
			dbInstance.Set(cmdArgs[1],cmdArgs[2])
			fmt.Println(*&(dbInstance).DBStorage)
		case GET:
			fmt.Println("It is a GET command")
			value:=dbInstance.Get(cmdArgs[1])
			fmt.Println(value)
		case DEL:
			fmt.Println("it is a DELETE command")
		case SHOW:
			fmt.Println("You have asked to see the whole DB")
			dbInstance.Show()
		default:
			fmt.Println("We are printing the default command!!")
		}
		
	}
}



