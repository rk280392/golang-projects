package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rk280392/fileSystemInteractions/readLocalFile"
	"github.com/rk280392/fileSystemInteractions/readRemoteFile"
	"github.com/rk280392/fileSystemInteractions/simpleRestWebServer"
	"github.com/rk280392/fileSystemInteractions/writeToLocal"
)

func main() {
	fmt.Println("Exploring various file system operations using golang")
	fmt.Println()

	pathToRead := "files/testReadFile.txt"
	pathToWrite := "files/testWriteFile.txt"
	writeContent := "This is test write file\n"
	restURL := "127.0.0.1:8000"

	data, err := readLocalFile.ReadLocalFile(pathToRead)
	if err != nil {
		fmt.Printf("Failed to read the file with error: %s", err)
	}
	// data is of the []byte type, so if you would like to use it as a string, you can simply convert it into one by using s := string(data)
	fmt.Println(string(data))
	fmt.Println("Reading from local Success!!")

	err = writeToLocal.WriteToLocal(pathToWrite, writeContent)
	if err != nil {
		fmt.Printf("Cannot write to the local file with error :%s", err)
	}
	fmt.Println("Writing to local Success!!")

	// Start webserver to test read and write for remote server. Once rest call is done , server will be shutdown.
	serverErrCh := make(chan error)
	go func() {
		err = simpleRestWebServer.SimpleRestWebServer("./files", restURL)
		serverErrCh <- err
	}()

	/*

		In this code, a select statement is used to wait for an error from the serverErrCh channel or a timeout using time.After.
		If no error is received within the specified timeout (in this example, 10 seconds), the code will continue, allowing you to
		handle the case where the server started successfully without errors
	*/
	select {
	case serverErr := <-serverErrCh:
		if serverErr != nil {
			log.Fatalf("Server start failed: %s", serverErr)
		} else {
			fmt.Println("Server started successfully")
		}
	case <-time.After(10 * time.Millisecond):
		err = readRemoteFile.ReadRemoteFile("http://" + restURL + "/remoteReadFile.txt")
		if err != nil {
			fmt.Printf("Error while reading the file from remote: %s", err)
			os.Exit(1)
		}
		fmt.Println("Reading from remote Success!!")
	}
	if err := simpleRestWebServer.StopServer(); err != nil {
		log.Fatalf("Server error : %s", err)
	} else {
		log.Println("Server has gracefully shutdown")
	}
}
