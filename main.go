// https://pkg.go.dev/github.com/ChimeraCoder/anaconda
package main

import (
	. "go-twitter-walker/conf"
	"fmt"
	"net/url"
)

func main() {

	api := InitTwitterApi()

	v := url.Values{}
	v.Set("count", "10")

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

}
