package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/dg222599/go-redis-db/command"
	"github.com/dg222599/go-redis-db/db"
)

type txnObject struct {
	DBName string
	PORT int64

}
func main(){
	
	fmt.Print("\n\n====****    Welcome to Redis DB     ****=====\n\n")
	
	// got the PORT where db needs to be run and dbName for the DB
	portNumber,dbName:=InitialMessage()

	fmt.Println(portNumber)

	//currentTxn := &txnObject{dbName:dbName,PORT:portNumber}

	
    // Creating a new DB instance for that DB name
	dbInstance,err := db.CreateDB(dbName)
	if err != nil {
		fmt.Println(err.Error())
	}
	
	// taking and handling the user commands until exit
	c:=make(chan os.Signal,1)
	signal.Notify(c,syscall.SIGINT,syscall.SIGTERM)
	go func(){
		<-c
		LastMessage()
		//os.Exit(1)
		
	}()

	for {
		command.HandleCommand(dbInstance)
	}
	
}

// to get the inital details for PORT and DB name
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

	//also add option for the user to stop the app anytime

	

	dbName =  args[1]

	//Connect from here only to the DB on Port

	return int64(PORT),dbName
}

func LastMessage() {
	fmt.Println("It seems you have terminated the operations...bye!!")
	os.Exit(1)
}

