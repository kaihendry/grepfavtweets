package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"regexp"
	"strings"
)

func main() {

	// https://github.com/ChimeraCoder/anaconda

	anaconda.SetConsumerKey(consumerkey)
	anaconda.SetConsumerSecret(consumersecret)
	api := anaconda.NewTwitterApi(accesstoken, accesstokensecret)

	// https://dev.twitter.com/rest/reference/get/favorites/list

	v := url.Values{}
	v.Set("screen_name", "kaihendry")
	v.Set("count", "200")

	searchResult, err := api.GetFavorites(v)

	if err != nil {
		panic(err)
	}

	reg, err := regexp.Compile(`[\s]+`)
	if err != nil {
		panic(err)
	}

	// https://godoc.org/github.com/ChimeraCoder/anaconda#Tweet
	for _, tweet := range searchResult {
		// fmt.Println("Count", k)
		// fmt.Println(tweet)
		tweet.Text = reg.ReplaceAllString(tweet.Text, " ")
		tweet.Text = strings.TrimSpace(tweet.Text)
		t := fmt.Sprintf("%s|%s|%s|%s", tweet.IdStr, tweet.User.ScreenName, tweet.CreatedAt, tweet.Text)

		// if tweet.Truncated {
		// 	fmt.Println("Truncated!")
		// }

		//fmt.Println("O", t)

		// https://godoc.org/github.com/ChimeraCoder/anaconda#Entities

		for _, u := range tweet.Entities.Urls {
			//fmt.Println(u)
			//fmt.Println("Expanded URL", u.Url, u.Expanded_url)
			t = strings.Replace(t, u.Url, u.Expanded_url, -1)
		}

		for _, m := range tweet.Entities.Media {
			//fmt.Println("Media URL", m.Media_url)
			t = strings.Replace(t, m.Url, m.Media_url, -1)
		}
		fmt.Println(t)
	}
}
