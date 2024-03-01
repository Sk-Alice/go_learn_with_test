package main

// https://studygolang.gitbook.io/learn-go-with-tests/gou-jian-ying-yong-cheng-xu/app-intro

import (
	"go_learn_with_test/GoApp"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	server := poker.NewPlayerServer(store)

	// 启动 HTTP 服务器，并将其监听在 5000 端口
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
