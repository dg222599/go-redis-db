package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/dg222599/go-redis-db/command"
	"github.com/dg222599/go-redis-db/db"
	"github.com/joho/godotenv"
)

type TxnObject struct {
	DBName string
	PORT int64

}
func main(){

	HandleTermination()
	
	fmt.Print("\n\n====****    Welcome to Redis DB     ****=====\n\n")

	err:= godotenv.Load()
	if err!=nil{
		log.Fatal(err.Error())
	}

	dbCount := os.Getenv("DB_COUNT")
	dbInstance:= db.CreateDB(dbCount)
	

	newOperationDB := command.NewDBOperation(dbInstance)

	server,err := StartTCPServer(fmt.Sprintf(":%s",os.Getenv("APP_PORT")))
	if err!=nil{
		log.Fatalf("Failed to start the TCP instance!!")
	}

	defer server.Close()

	for {
		conn,err := server.Accept()
		if err!=nil{
			fmt.Printf("Failed to accept connection: %v\n",err)
			continue
		}

		go HandleNewConnection(conn,newOperationDB)

	}

}

func StartTCPServer(port string) (net.Listener , error) {
	server,err := net.Listen("tcp",port)
	if err!=nil{
		return nil,err
	}

	fmt.Println("Server started on PORT:",port)
	return server,nil
}

func HandleNewConnection(conn net.Conn ,operationDB command.DBOperation) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	
	for {
		writer.Flush()

		//Get Input from the User
		command,err := TakeCommand(reader)
		if err!=nil{
			showResult(writer,err)
			break
		}

		var result interface{}
		result  =  command.Execute(command)
		showResult(writer,result)


	}
}

func TakeCommand(reader *bufio.Reader) (command.Command , error) {
	line,err := reader.ReadString('\n')
	if err!=nil{
		return command.Command{},err
	}

	line = strings.TrimSpace(line)

	words := strings.Split(line," ")

	args := make([]interface{},0,10)

	count:=0
	longArgs := false
	qoutesEnded := true

	completeWord:= ""
	for _,currWord := range words {
		
		if strings.HasPrefix(currWord,`"`) {
			longArgs = true
			qoutesEnded = false
		}

		if longArgs {
			completeWord = fmt.Sprintf("%s %s",completeWord,currWord)

			if strings.HasSuffix(currWord,`"`) {
				args = append(args,strings.ReplaceAll(strings.TrimSpace(currWord),`"`,""))
				completeWord = ""
				longArgs = false
				qoutesEnded = false
				count++

			} }else {
				args = append(args,strings.ReplaceAll(currWord,`"`,""))
				count++
			}
		}

		if count<1{
			return command.Command{},fmt.Errorf("empty command")
		}

		if !qoutesEnded {
			return command.Command{},fmt.Errorf("Wrong format error - unbalanced Qoutes in args")

		}

		cmdName:= strings.ToUpper(strings.TrimSpace(args[0].(string)))

		return command.NewCommand(cmdName),nil

		
}

func showResult(writer *bufio.Writer,result interface{}) {
	switch res:=result.(type) {
	case []interface{}:
		for i,item := range res {
			fmt.Fprintf(writer,"%d) %v\n",i+1,item)
		}
	default:
		fmt.Fprintf(writer,"%v\n",result)
	}
	writer.Flush()
}

func HandleTermination() {

	// Terminating the Redis Client in case of CTRL + C
	c:=make(chan os.Signal,1)
	signal.Notify(c,syscall.SIGINT,syscall.SIGTERM)
	go func(){
		<-c
		LastMessage()
		
		
	}()
}
func LastMessage() {
	fmt.Println("It seems you have terminated the operations...bye!!")
	os.Exit(1)
}

