package readWrite

import (
	"io"
	"log"
	"net/http"
	"os"
)

// using a random 1MB test file
var url string = "http://speedtest.ftp.otenet.gr/files/test1Mb.db"

func readWrite() {
	r, err := http.Get(url)
	if err != nil {
		log.Println("Cannot get from URL", err)
	}
	defer r.Body.Close()
	data, _ := io.ReadAll(r.Body)
	os.WriteFile("rw.data", data, 0o755)
}
