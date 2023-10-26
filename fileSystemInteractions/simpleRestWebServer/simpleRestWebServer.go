package simpleRestWebServer

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func SimpleRestWebServer(path, serverPort string) error {
	dir := path
	fileServer := http.FileServer(http.Dir(dir))
	http.Handle("/", fileServer)

	port := serverPort
	log.Printf("Server listening on %s", port)

	fmt.Println(port)
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		return fmt.Errorf("server error :%s", err)
	}
	return nil

}

func StopServer() error {
	var server *http.Server
	if server == nil {
		return nil // Server not started
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return server.Shutdown(ctx)
}
