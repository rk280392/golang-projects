package readRemoteFile

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func ReadRemoteFile(url string) error {

	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		return fmt.Errorf("error in creating the request: %s", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error in client.do : %s", err)
	}

	/*
			The client.Do function in Go's net/http package returns an http.Response and an error,
			but the error is usually related to network issues, such as DNS resolution problems or
			connection timeouts, not the HTTP status code returned by the server.

		To check for specific HTTP status codes like 501, you should examine the resp.StatusCode and handle it accordingly in your code, as shown
	*/

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received a non-OK status code: %d", resp.StatusCode)
	}
	flags := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	/*
		Create the file if it doesn't exist: os.O_CREATE.
		Write to a file: os.O_WRONLY.
		If the file exists, truncate it versus append to it: os.O_TRUNC. https://pkg.go.dev/os#pkg-constants.
	*/
	f, err := os.OpenFile("./files/readandWriteFromRemote.txt", flags, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	/* schedules the f.Close() method to be called when the surrounding function  (the function containing this defer statement) exits.
	In our case defer will be called at the end of ReadRemoteFile function to close the file.
	*/
	if _, err := io.Copy(f, resp.Body); err != nil {
		return err
	}

	return nil
}
