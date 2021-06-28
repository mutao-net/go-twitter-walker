// https://pkg.go.dev/github.com/ChimeraCoder/anaconda
package main

import (
	. "go-twitter-walker/conf"
	"fmt"
	"net/url"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func main() {

	api := InitTwitterApi()

	v := url.Values{}
	v.Set("count", "100")

	tweets, err := api.GetHomeTimeline(v)
	if err != nil {
		panic(err)
	}

	for _, tweet := range tweets {
		fmt.Print(tweet.Id)
		fmt.Print(" : ", tweet.FavoriteCount)
		fmt.Print(" :tweet: ", tweet.User.ScreenName)
		fmt.Println(" : ", tweet.FullText)
	}

	db, err := sql.Open("mysql", "root:rootpassword@tcp(127.0.0.1:3306)/twitter")
	if err != nil {
		panic(err.Error())
		fmt.Print("connect failed.")
	}
	defer db.Close()

	stmt, err := db.Prepare(`
	INSERT INTO timeline (id, name, tweet, favorite_count)
	VALUES (?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE id = ?
	`)
	if err != nil {
		return
	}
	for _, tweet := range tweets {
		fmt.Print(tweet.Id)
		fmt.Print(" : ", tweet.FavoriteCount)
		fmt.Print(" :tweet: ", tweet.User.ScreenName)
		fmt.Println(" : ", tweet.FullText)
		_, err := stmt.Exec(tweet.Id, tweet.User.ScreenName, tweet.FullText, tweet.FavoriteCount, tweet.Id)
		if err != nil {
			panic(err.Error())
			fmt.Print("commit failed.")
		}
	}
}
