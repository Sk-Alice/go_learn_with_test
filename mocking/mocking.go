//package main

package mocking

import (
	"fmt"
	"io"
	"time"
)

const finalWorld = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	// 跟踪记录 Sleep() 被调用了多少次
	Calls int
}

type ConfigurableSleeper struct {
	duration time.Duration
}

// Sleep 真正运行时调用
func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

// Sleep 测试运行时调用，通过记录调用的次数模拟真实的时间等待行为
// 避免了 缓慢的测试会破坏开发人员的生产力
func (s *SpySleeper) Sleep() {
	s.Calls++
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
