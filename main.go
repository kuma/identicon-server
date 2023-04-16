package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgryski/go-identicon"
)

func handler(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.Path, "/")
	args = args[1:]

	if len(args) != 1 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	item := args[0]
	if !strings.HasSuffix(item, ".png") {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	item = strings.TrimSuffix(item, ".png")

	key := []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff}
	icon := identicon.New7x7(key)

	data := []byte(item)
	pngdata := icon.Render(data)

	w.Header().Set("Content-Type", "image/png")
	w.Write(pngdata)

	return
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
