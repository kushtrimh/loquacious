package twitter

import (
	"encoding/json"
	"github.com/kushtrimh/loquacious/config"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"
	"time"
)

// Tweet represents the data of a tweet
type Tweet struct {
	Id        string `json:"id_str"`
	CreatedAt string `json:"created_at"`
}

// UserTimeline returns tweets from the timeline of the specified user.
// To retrieve the tweets user should have a public profile
func (t *Twitter) UserTimeline(user string) ([]Tweet, error) {
	queryParams := url.Values{}
	queryParams.Add("screen_name", user)
	queryParams.Add("count", strconv.Itoa(config.App.UserTimelineTweetCount))
	response, err := t.client.Get("/statuses/user_timeline.json", queryParams)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	tweets := make([]Tweet, 100)
	err = json.Unmarshal(body, &tweets)
	if err != nil {
		return nil, err
	}
	log.Printf("Response status code for user timeline %d request. User @%s, tweets returned %d",
		response.StatusCode, user, len(tweets))
	return tweets, nil
}

// TodayTweetCount returns the number of tweets made today
// for the specified user
func (t *Twitter) TodayTweetCount(user string) (int, error) {
	tweets, err := t.UserTimeline(user)
	if err != nil {
		return 0, err
	}
	count := 0
	thisDay, thisMonth, thisYear := time.Now().Date()
	for _, tweet := range tweets {
		tweetTime, err := time.Parse(tweetDateTimeLayout, tweet.CreatedAt)
		if err != nil {
			log.Println(err)
			continue
		}
		day, month, year := tweetTime.Date()
		if day == thisDay && month == thisMonth && year == thisYear {
			count++
		}
	}
	return count, nil
}
