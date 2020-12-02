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

// tweetCount represents a user (twitter handle) and the tweet count
// for the specified user
type tweetCount struct {
	user  string
	count int
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

// TodayTweetCounts returns a map containing all the number of tweets
// for today, for all the added users
func (t *Twitter) TodayTweetCounts() map[string]int {
	users := config.App.FollowedUsers
	counts := make(map[string]int)
	channels := make([]chan tweetCount, 0, len(users))

	for _, user := range users {
		c := make(chan tweetCount)
		channels = append(channels, c)
		go t.count(c, user)
	}

	for _, c := range channels {
		var data tweetCount
		data = <-c
		counts[data.user] = data.count
	}
	return counts
}

// count is a function which gets the tweet count of the specified
// user and sends the data back to a channel
func (t *Twitter) count(c chan tweetCount, user string) {
	count, err := t.TodayTweetCount(user)
	if err != nil {
		log.Printf("Could not get tweet count for user %s, %v", user, err)
		return
	}
	c <- tweetCount{user, count}
}
