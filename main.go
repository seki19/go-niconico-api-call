package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type MusicData struct {
	StartTime     time.Time `json:"startTime"`
	LengthSeconds int       `json:"lengthSeconds"`
	ViewCounter   int       `json:"viewCounter"`
	ContentID     string    `json:"contentId"`
	Title         string    `json:"title"`
	ThumbnailURL  string    `json:"thumbnailUrl"`
}

type Response struct {
	Data []MusicData `json:"data"`
	Meta struct {
		ID         string `json:"id"`
		TotalCount int    `json:"totalCount"`
		Status     int    `json:"status"`
	} `json:"meta"`
}

func main() {
	resp, err := http.Get("https://api.search.nicovideo.jp/api/v2/snapshot/video/contents/search?q=" + url.QueryEscape("初音ミク") + "&targets=tagsExact&fields=contentId,title,viewCounter,startTime,thumbnailUrl,lengthSeconds&filters[categoryTags][0]=VOCALOID&_sort=-viewCounter&_offset=0&_limit=3&_context=apitest")
	fmt.Println(resp)
	if err != nil {
		log.Fatal(err)
	}

	// The client must close the response body when finished with it:
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Response // nil slice

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	for _, music := range data.Data {
		// fmt.Printf("%s %s\n", item.CreatedAt, item.Title)
		fmt.Println(music)
	}
}
