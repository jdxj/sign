package task

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	ticker := time.NewTicker(time.Minute)
	//timer := time.NewTimer(time.Minute)

	for {
		f := func() {
			fmt.Println("模拟执行任务:", rand.Int())
		}
		dur := rand.Int63n(int64(time.Minute))
		time.AfterFunc(time.Duration(dur), f)
		<-ticker.C
	}
}

func TestDur(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dur := time.Duration(r.Int63n(86400) * 1e9)
	fmt.Println(dur)
}

func TestTimer2(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 现在时刻, 假设是昨天, 只是测试
	now := time.Now().Add(-24 * time.Hour)
	// 明天这个时候
	tomNow := now.Add(24 * time.Hour)
	// 明天0点
	tom0AM := time.Date(tomNow.Year(), tomNow.Month(), tomNow.Day(), 0, 0, 0, 0, tomNow.Location())
	tom830AM := tom0AM.Add(8 * time.Hour).Add(30 * time.Minute)
	tom830PM := tom830AM.Add(12 * time.Hour)

	for i := 0; i < 1000; i++ {
		// 明天随便几点 8:30~20:30
		inc := time.Duration(r.Intn(12 * 60 * 60))
		tomSome := tom830AM.Add(inc * time.Second)

		if tomSome.Unix() > tom830PM.Unix() {
			// 随机失败
			panic("fail")
		}
		time.Sleep(time.Millisecond)
	}
}

func TestRandTime(t *testing.T) {
	for i := 0; i < 100; i++ {
		tim := randTime()
		fmt.Println(tim)
		time.Sleep(time.Millisecond)
	}
}

func TestFormatTime(t *testing.T) {
	fmt.Printf("%s\n", time.Now().Format("2006-01-02 15:04:05"))
}
