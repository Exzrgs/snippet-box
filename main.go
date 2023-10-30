package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, req *http.Request) {
	/*
		これで、subtree pathでも他を受け付けないようにできる.

		このifでreturnしなかったら、レスポンスは改行されて"Hello from snippetBox"が追記される形になる！
	*/
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	w.Write([]byte("Hello from snippetBox"))
}

func snippetView(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil || id < 1 {
		log.Println(err)
		http.Error(w, "bad query param", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		// 空白とか入るとダメ(Allow Methodみたいな)
		w.Header().Set("Allow", "POST")

		/*
			さっきの処理のヘルパー関数。ヘッダーとボディに別々に書き込まなくていい
		*/
		http.Error(w, "mothod not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("server start at port 4000")

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
