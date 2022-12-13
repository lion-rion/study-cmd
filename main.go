package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var i = flag.Int("i", 0, "数値")
	var s = flag.String("s", "", "文字列")
	var b = flag.Bool("b", false, "真偽値")
	flag.Parse()
	fmt.Println(*i, *s, *b)
}

//勉強時間を計測する関数
func timeTrack() {
	start := time.Now()
	elapsed := time.Since(start)
	fmt.Printf("勉強開始")
	fmt.Println(start.Format("2006-01-02 15:04:05"))
	var button string
	fmt.Scan(&button)
	fmt.Printf("勉強終了")
	fmt.Printf("勉強時間は%vです", elapsed)

	// var str string
	// fmt.Println("文字列を入力してください")
	// fmt.Scan(&str) // データを格納する変数のアドレスを指定
	// fmt.Println(str)
	// start := time.Now()
	// fmt.Println("終了したらおしえてください")
	// fmt.Scan(&str) // データを格納する変数のアドレスを指定
	// end := time.Now()
	// fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
