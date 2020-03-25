package main // import "server"

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 配列の宣言が分からない
	// arr := make(map[string]string, 10)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping test",
		})
		ExtractItemUrl()
	})
	r.Run(":8888")
}

func ExtractItemUrl() {
	// get document //
	// 検索結果をここで挿入したいが、今のところ決め打ち
	// ジャンルも絞っていない
	searchValue := "Maison+Margiela"
	doc, err := goquery.NewDocument("https://zozo.jp/search/?p_keyv=" + searchValue)
	if err != nil {
		fmt.Println(err.Error())
	}

	// search部分 //
	// 取得したデータに対する処理
	// 1ページmax120のitems
	// URLを配列に入れて取り出したい
	// 各itemのURLは取得OK
	doc.Find(".catalog-link").Each(func(i int, s *goquery.Selection) {
		target, _ := s.Attr("href")
		fmt.Println("https://zozo.jp" + target)

		// fmt.Printf("Title: %s %s\n", s.Find(".caption").Text(), s.Find(".by").Text())
		// fmt.Printf("%s\n\n", strings.Replace(s.Find(".text").Text(), "\n", "", -1))
	})
	return
}
