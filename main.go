package main

import (
	"net/http"
	"sync"
	"log"
	"time"
)


func main() {
	url := "http://localhost:8080"			// アクセスするURLだよ！

	maxConnection := make(chan bool,200)	// 同時に並列する数をしていできるよ！（第二引数）
	wg := &sync.WaitGroup{}					// 並列処理が終わるまでSleepしてくれる便利なやつだよ！

	count := 0								// いくつアクセスが成功したかをアカウントするよ！
	start := time.Now()						// 処理にかかった時間を測定するよ！
	for maxRequest := 0; maxRequest < 10000; maxRequest ++{		// 10000回リクエストを送るよ！
		wg.Add(1)						// wg.add(1)とすると並列処理が一つ動いていることを便利な奴に教えるよ！
		maxConnection <- true				// ここは並列する数を抑制する奴だよ！詳しくはググって！
		go func() {							// go func(){/*処理*/}とやると並列処理を開始してくれるよ！
			defer wg.Done()					// wg.Done()を呼ぶと並列処理が一つ終わったことを便利な奴に教えるよ！

			resp, err := http.Get(url)		// GETリクエストでアクセスするよ！
			if err != nil {					// err ってのはエラーの時にエラーの内容が入ってくるよ！
				return						// 回線が狭かったりするとここでエラーが帰ってくるよ！
			}
			defer resp.Body.Close() 		// 関数が終了するとなんかクローズするよ！（おまじない的な）

			count++							// アクセスが成功したことをカウントするよ！
			<-maxConnection					// ここは並列する数を抑制する奴だよ！詳しくはググって！
		}()
	}
	wg.Wait()								// ここは便利な奴が並列処理が終わるのを待つよ！
	end := time.Now()						// 処理にかかった時間を測定するよ！
	log.Printf("%d 回のリクエストに成功しました！\n", count)	// 成功したリクエストの数を表示してくれるよ！
	log.Printf("%f 秒処理に時間がかかりました！\n",(end.Sub(start)).Seconds())			//何秒かかったかを表示するよ！
}