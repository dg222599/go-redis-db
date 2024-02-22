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

func ValidateCommand(cmdLine string) (bool,Command){
    cmdArgs := strings.Fields(cmdLine)
	totalArgs := len(cmdArgs)
	if totalArgs < 1 {
		fmt.Println("You have entered empty command...please enter again ")
		return false,Command{}
	}
	currentCmd:=Command{
		Name : strings.TrimSpace(cmdArgs[0]) ,
		Key : "" , 
		Value : nil,
	}

	switch currentCmd.Name {
		case SET:
			if totalArgs <=2 {
				fmt.Println("All the arguments are not provided need -> SET key value")
				return false,Command{}
			}
			currentCmd.Key  = strings.TrimSpace(cmdArgs[1])
			currentCmd.Value = strings.TrimSpace(strings.Join(cmdArgs[2:]," "))
		case GET:
			if totalArgs!=2{
				fmt.Println("Key not present in the command , need --> GET key")
				return false,Command{}
			}
			currentCmd.Key = strings.TrimSpace(cmdArgs[1])
		case DEL:
			if totalArgs!=2{
				fmt.Println("Key not present in the command , need --> DEL key")
				return false,Command{}
			}
			currentCmd.Key = strings.TrimSpace(cmdArgs[1])
		case SHOW:
			if totalArgs !=1{
				fmt.Println("SHOW command is in wrong format  , just enter --> SHOW")
				return false,Command{}
			}
		default:
			fmt.Println("Other commands ...so frepass!!")
	}
	
	return true,currentCmd
}

func  HandleCommand(dbInstance *db.DB) {
	fmt.Println("Enter the Command or press CTRL+C to exit")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err.Error())
	}

	validationStatus,currentCmd := ValidateCommand(line)



	if !validationStatus{
		fmt.Println("InValid Command format ...pls enter again")
		HandleCommand(dbInstance)
	} else {

		//fmt.Println("Congrats the command is verified!!~")
		
		switch currentCmd.Name {
		case SET:
			dbInstance.Set(currentCmd.Key,currentCmd.Value)
		case GET:
			value:=dbInstance.Get(currentCmd.Key)
			if value== nil {
				fmt.Println("Key does not exist!!")
			} else {
				fmt.Println("Value for key is -->",value)
			}
		case DEL:
			dbInstance.Delete(currentCmd.Key)
		case SHOW:
			fmt.Println("You have asked to see the whole DB")
			dbInstance.Show()
		default:
			fmt.Println("We are printing the default command!!")
		}
		
	}
}



