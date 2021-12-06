// https://pkg.go.dev/github.com/ChimeraCoder/anaconda
package main

import (
	. "go-twitter-walker/conf"
	"fmt"
	"net/url"
	"regexp"
	"io"
    "net/http"
    "os"
	// "database/sql"
	// _"github.com/go-sql-driver/mysql"
)

func main() {

	api := InitTwitterApi()

	v := url.Values{}
	v.Set("screen_name", "moe_five")
	v.Set("count", "1000")

	tweets, err := api.GetUserTimeline(v)
	if err != nil {
		panic(err)
	}

	for _, tweet := range tweets {
		/**
		fmt.Print(tweet.Id)
		fmt.Print(" : ", tweet.FavoriteCount)
		fmt.Print(" :tweet: ", tweet.User.ScreenName)
		fmt.Println(" : ", tweet.FullText)
		**/
		mediaList := []string{}
		for _, media := range tweet.Entities.Media {
			reg := regexp.MustCompile(`([^\/]+?)(\.jpg|\.jpeg|\.gif|\.png)$`)
			fmt.Println(reg.FindString(media.Media_url_https))
			getImg(media.Media_url_https)
			mediaList = append(mediaList, media.Media_url_https)
			// fmt.Println(mediaList)
		}
		// fmt.Println(" : ", tweet.Entities)
	}
/**
	db, err := sql.Open("mysql", "root:rootpassword@tcp(127.0.0.1:3306)/twitte
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
	
*/
}
func getImg(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	reg := regexp.MustCompile(`([^\/]+?)(\.jpg|\.jpeg|\.gif|\.png)$`)
	file, err := os.Create("./test/" + reg.FindString(url))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, response.Body)
}