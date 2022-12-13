package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	// var i = flag.Int("i", 0, "数値")
	// var s = flag.String("s", "", "文字列")
	// var b = flag.Bool("b", false, "真偽値")
	// var setup = flag.Bool("setup", false, "真偽値")
	// flag.Parse()
	// if *b == true {
	// 	timeTrack()
	// }

	// if *setup == true {
	// 	Setup()
	// }

	//計測開始
	timeTrack()

}

func Setup() {
	//gitの設定を行う
	err := exec.Command("git init").Run()
	fmt.Println("git init...")
	if err != nil {
		fmt.Println("git init error")
	}

	err = exec.Command("git add .").Run()
	fmt.Println("git add...")
	if err != nil {
		fmt.Println("git add error")
	}

	err = exec.Command("git commit -m 'initial commit'").Run()
	fmt.Println("git commit...")
	if err != nil {
		fmt.Println("git commit error")
	}

	err = exec.Command("git branch -M main").Run()
	fmt.Println("git branch -M main...")
	if err != nil {
		fmt.Println("git branch -M main error")
	}

	//入力を受け付ける
	var str string
	fmt.Print("リモートリポジトリのURLを入力してください : ")
	fmt.Scan(&str)
	err = exec.Command(str).Run()
	fmt.Println("git setting...")

	err = exec.Command("git push -u origin main").Run()
	fmt.Println("git push...")

}

//勉強時間を計測する関数
func timeTrack() {

	start := time.Now()
	//fmt.Println(start.Format("2006-01-02 15:04:05"))
	fmt.Println("------------勉強開始------------")

	fmt.Println("-----終了時はEnterを押してください-----")
	fmt.Scanln() //エンター受付
	end := time.Now()
	fmt.Printf("---------勉強終了----------")
	fmt.Println("%d秒勉強しました", getdiff(start, end))

	var answer string
	fmt.Println("勉強時間を記録しますか？(y/n)")
	fmt.Scan(&answer)
	if answer == "y" {
		fmt.Println("記録します")
		Createlog(start, end)

	} else if answer == "n" {
		fmt.Println("記録しません")
	} else {
		fmt.Println("yかnで答えてください")
	}
}

func Createlog(start time.Time, end time.Time) {
	//ファイルを作成する

	month := strconv.Itoa(int(start.Month()))
	year := strconv.Itoa(start.Year())
	day := strconv.Itoa(start.Day())

	var dir string = year + "/" + month
	var filename string = dir + "/" + day + ".txt"

	os.MkdirAll(dir, 0777)
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//書き込み
	data := []byte(
		"\n開始 : " + start.Format("2006-01-02 15:04:05") + "\n終了 : " + end.Format("2006-01-02 15:04:05") + "\n勉強時間 : " + getdiff(start, end) + "\n",
	)
	_, err = f.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	gitpush(month, day)

}

func getdiff(start time.Time, end time.Time) string {

	_duration := end.Sub(start)
	_microsecond := int(_duration / 1000)
	_millisecond := int(_microsecond / 1000)
	_second := int(_millisecond / 1000)

	//秒数を時間に変換
	if _second > 3600 {
		_hour := _second / 3600
		_minute := _second % 3600 / 60
		_second = _second % 3600 % 60
		diff := strconv.Itoa(_hour) + "時間" + strconv.Itoa(_minute) + "分" + strconv.Itoa(_second) + "秒"
		return diff
	} else if _second > 60 {
		_minute := _second / 60
		_second = _second % 60
		diff := strconv.Itoa(_minute) + "分" + strconv.Itoa(_second) + "秒"
		return diff
	}
	diff := strconv.Itoa(_second) + "秒"
	return diff
}

func gitpush(month string, day string) {
	//gitの設定を行う
	err := exec.Command("git add .").Run()
	fmt.Println("git add...")
	if err != nil {
		fmt.Println("git add error")
	}

	err = exec.Command("git", "commit", "-m", month+"/"+day).Run()
	fmt.Println("git commit...")
	if err != nil {
		fmt.Println("git add error")
	}

}
