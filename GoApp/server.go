package poker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PlayerStore stores score information about players
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

// Player stores a name with a number of wins
type Player struct {
	Name string
	Wins int
}

// PlayerServer is a HTTP interface for player information
type PlayerServer struct {
	store        PlayerStore
	http.Handler // 嵌入
}

const jsonContentType = "application/json"

// NewPlayerServer creates a PlayerServer with routing configured
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)

	// json.NewEncoder(w) 创建了一个 JSON 编码器，它将 JSON 数据写入到 w 中。
	// Encode(leagueTable) 方法将 leagueTable 编码成 JSON 格式，并将其写入到 w 中
	_ = json.NewEncoder(w).Encode(p.store.GetLeague())

	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		// 设置 HTTP 响应的状态码
		w.WriteHeader(http.StatusNotFound)
	}

	_, _ = fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
