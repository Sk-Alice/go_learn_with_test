package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	// 用于获取文件的元信息
	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		return nil, fmt.Errorf("problem creating file system player store, %v ", err)
	}

	return store, nil
}

type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	// 根据给定的比较函数，Slice 对提供的切片进行排序
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})

	return f.league
}

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.GetLeague().Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Encode(f.league)
}
