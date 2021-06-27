package conf

import (
	"github.com/ChimeraCoder/anaconda"
)

func InitTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey("")
	anaconda.SetConsumerSecret("")
	api := anaconda.NewTwitterApi("", "")
	return api
}
