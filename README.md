# A Simple Key-Value DB in Go 
A Simple Implementation of Redis DB in Golang , the tool is built as a CLI that allows user to interact witha simple in-memory key-value database through a TCP server. User can execute various commands using command-line interface forexecuting various commands and to get results.

# Installation Guide
Need to have Go installed and , also need ncat like tool for connecting to TCP server.  

1.Clone the repo.  

2. Navigate to the project directory and run:

   ```bash
   go build -o go-redis-db


3. Run the exe file using.

   
    ``` ./go-redis-db  ```

 # Usage

 1. Create a .env file as given in the repo , the environment file should contain PORT number on which the TCP server listens to and the DB_COUNT variable will have the
    number of in memory databases.
    ``` APP_PORT="9736"  ```
    ``` DB_COUNT=16   ```

 2. Run command  ``` ./go-redis-db ``` to run the TCP server.
 3. Open another terminal window and use ncat tool to connect to TCP using command.
    ``` ncat localhost 9736 ```
 4. Now the CLI tool is connected and you can interact with the DB using following.
    ```  COMMAND [arg1] [arg2] ```
 5. Commands supported by the Tool are , also type ```  HELP  ```  to get the available commands.

    
    <img width="247" alt="image" src="https://github.com/dg222599/go-redis-db/assets/56475367/35d711d0-a738-43dc-8f1b-9227666a8d81">

    Change the values for ``` key,value,db-number ``` accordingly in the commands , also the commands are case insensitive so can use in that manner.


  6. After entering command , tool will show the result or error(if any) on the window.
  7. To close the connection use   ``` CTRL + C ```   or  ``` EXIT ```  command.
     

# Usage with Docker

1. Current repo version has DockerFile in it, if Docker is installed in the user system then can build the image and run the container directly.

2. Steps for building the image and running the container:

    ```
    docker build -t redis-server-image:latest .
    ```

    ```
    docker run -p 9736:9736 redis-server-image:test
    ```

3. After the server starts successfully on the container, use the same steps as listed above with ncat for interacting with the key-value db server.

