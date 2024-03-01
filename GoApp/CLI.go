package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	// bufio 包实现了 I/O 缓冲。
	// 它封装了一个 io.Reader 或 io.Writer 对象，创建了另一个对象（Reader 或 Writer），
	// 也实现了接口，并为文本 I/O 提供了缓冲和一些帮助
	in *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}
}

func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	// Scanner.Scan() 逐行读取内容
	cli.in.Scan()
	// 返回 scanner 读取的 string
	return cli.in.Text()
}
