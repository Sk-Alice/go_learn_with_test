//package main

package mocking

import (
	"fmt"
	"io"
	"time"
)

const finalWorld = "Go!"
const countdownStart = 3

const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
}

// CountdownOperationSpy 监视操作顺序是否正确
// 本次测试中应为：
//
//	Sleep
//	Print N
//	Sleep
//	Print N-1
//	Sleep
//	etc
type CountdownOperationsSpy struct {
	Calls []string
}

// 实现了 Sleeper 接口
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// 实现了 io.Writer 接口
func (s *CountdownOperationsSpy) Write([]byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// Sleep 真正运行时调用
func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		_, _ = fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	_, _ = fmt.Fprint(out, finalWorld)
}

//func main() {
//	sleeper := &ConfigurableSleeper{1 * time.Second}
//	Countdown(os.Stdout, sleeper)
//}
