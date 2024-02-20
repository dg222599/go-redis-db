package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dg222599/go-redis-db/command"
)

type txnObject struct {
	dbName string
	PORT int64

}



func main(){
	
	fmt.Print("\n\n====****    Welcome to Redis DB     ****=====\n\n")
	portNumber,dbName:=InitialMessage()

	currentTxn := &txnObject{dbName:dbName,PORT:portNumber}
	
	currentTxn.HandleCommand()
	
	
	
     
}

func (currentTxn *txnObject) HandleCommand(){
	fmt.Println("Enter the Command - please enter in this format--> SET key value")
	reader := bufio.NewReader(os.Stdin)
	line,err := reader.ReadString('\n')
	if err!=nil{
		log.Fatal(err.Error())
	}

	validationStatus,_:=ValidateCommand(line)

	if validationStatus == false {
		fmt.Println("InValid Command format ...pls enter again")
		currentTxn.HandleCommand()
	}


}

func ValidateCommand(cmdLine string) (bool,error){
    cmdArgs := strings.Split(cmdLine," ")

	if len(cmdArgs) < 1 {
		fmt.Println("You have entered empty command...please enter again ")
		return false,nil
	}

    userCmd := command.InitCommand()

	fmt.Println(userCmd.Name)

	fmt.Println(userCmd)

	return true,nil
	
}

func InitialMessage() (int64,string) {
	var PORT int
	var dbName string
	var err error

	
	fmt.Println("Please enter the PORT , DB name from which you want to connect to DB server , ex - 6379 root")
	
	reader := bufio.NewReader(os.Stdin)
	line,err := reader.ReadString('\n')
    
	args := strings.Split(line, " ")
    
	if len(args) !=2 {
		fmt.Println("Please enter both the PORT number and db Name")
		InitialMessage()
	}

	PORT,err = strconv.Atoi(args[0])
	if err!=nil{
		fmt.Println("Error in Port Number format--> , PLease enter again correctly")
		fmt.Println(err.Error())
		InitialMessage()
		
	}

	

	dbName =  args[1]

	//Connect from here only to the DB on Port

	return int64(PORT),dbName
}
