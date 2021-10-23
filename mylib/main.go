package mylib

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"time"

	"go_learning_2/calc"
)

const (
	c1 = iota
	c2 = iota
	c3 = iota
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func main() {
	fmt.Println("main")
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(Average(s))

	fmt.Println(calc.Sum(1, 2))

	person := calc.Person{Name: "Mike", Age: 20}
	fmt.Println(person)
	fmt.Println(calc.Public)

	// エラー
	// fmt.Println(calc.private)

	// spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	// fmt.Print(spy.CSV())
	// rsi2 := talib.Rsi(spy.Close, 2)
	// fmt.Println(rsi2)

	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Year(), t.Month(), t.Hour(), t.Minute(), t.Second())

	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	// 何度も同じ正規表現を使うときは下記のように定数化する
	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)

	i := []int{5, 3, 2, 8, 7}
	s2 := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}

	fmt.Println(i, s2, p)
	sort.Ints(i)
	sort.Strings(s2)
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(i, s2, p)

	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB)

	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	go longProcess(ctx, ch)

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case <-ch:
			fmt.Println("success")
			return
		}
	}

}

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}
