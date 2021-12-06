# go-twitter-walker

## usage
- passは画像を保存するディレクトリを指定する。

$ go run main.go {path}

## Step

### Step1
- Twitterのtokenを設定する

$ vi conf/conf.go

```
package conf

import (
	"github.com/ChimeraCoder/anaconda"
)

func InitTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey("XXXXX")
	anaconda.SetConsumerSecret("XXXXXX")
	api := anaconda.NewTwitterApi("XXXXX", "XXXXX")
	return api
}

```