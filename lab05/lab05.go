package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// TODO: Please create a struct to include the information of a video
type PageData struct {
	Title        string
	Id           string
	ChannelTitle string
	LikeCount    string
	ViewCount    string
	PublishedAt  string
	CommentCount string
}

func error_handle() {

}

func num_handle(input string) string {
	var output string
	var cnt int = 0
	for i := len(input) - 1; i >= 0; i-- {
		if cnt%3 == 0 && i != len(input)-1 {
			output = "," + output
		}
		output = string(input[i]) + output
		cnt++
	}
	return output
}

func time_handle(input string) string {
	var output string
	var year int
	var month int
	var date int
	var time string
	fmt.Sscanf(input, "%d-%d-%dT%s", &year, &date, &month, &time)
	output = strconv.Itoa(year) + "年" + strconv.Itoa(month) + "月" + strconv.Itoa(date) + "日"
	return output

}

func YouTubePage(w http.ResponseWriter, r *http.Request) {
	var filename string
	var title string
	var id string
	var ctitle string
	var lcnt string
	var vcnt string
	var pa string
	var ccnt string
	data := PageData{
		Title:        title,
		Id:           id,
		ChannelTitle: ctitle,
		LikeCount:    lcnt,
		ViewCount:    vcnt,
		PublishedAt:  pa,
		CommentCount: ccnt,
	}
	// TODO: Get API token from .env file
	err := godotenv.Load() // Load environment variable from .env file
	if err != nil {
		fmt.Println("========ERROR==========\n")
		filename = "error.html"
		tmpl, err := template.ParseFiles("error.html")
		if err != nil {
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	password := os.Getenv("YOUTUBE_API_KEY") // Get value from environment variable
	// fmt.Printf("password: %s\n", password)

	// TODO: Get video ID from URL query `v`
	key := "v"
	var videoid string = r.URL.Query().Get(key)
	if videoid == "" {
		fmt.Println("========ERROR==========\n")
		filename = "error.html"
		tmpl, err := template.ParseFiles("error.html")
		if err != nil {
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	// fmt.Printf("video id: %s\n", videoid)

	// TODO: Get video information from YouTube API
	url := "https://www.googleapis.com/youtube/v3/videos?key=" + password + "&id=" + videoid + "&part=" + "statistics,snippet"
	resp, err := http.Get(url) // resp is the response of the url
	if err != nil {
		fmt.Println("========ERROR1==========\n")
		filename = "error.html"
		tmpl, err := template.ParseFiles("error.html")
		if err != nil {
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("========ERROR==========\n")
		filename = "error.html"
		tmpl, err := template.ParseFiles("error.html")
		if err != nil {
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	fmt.Println(string(body))

	// TODO: Parse the JSON response and store the information into a struct
	var m map[string]interface{}
	err1 := json.Unmarshal(body, &m)
	if err1 != nil {
		fmt.Println("========ERROR==========\n")
		filename = "error.html"
		tmpl, err := template.ParseFiles("error.html")
		if err != nil {
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	items := m["items"].([]interface{})
	if len(items) == 0 {
		fmt.Println("========ERROR==========\n")
		filename = "error.html"
		tmpl, err := template.ParseFiles("error.html")
		if err != nil {
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	item := items[0].(map[string]interface{})
	if err1 != nil {
		fmt.Println("========ERROR==========\n")
		filename = "error.html"
		tmpl, err := template.ParseFiles("error.html")
		if err != nil {
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	statistics := item["statistics"].(map[string]interface{})
	snippet := item["snippet"].(map[string]interface{})

	if err != nil {
		fmt.Println("========ERROR==========\n")
		filename = "error.html"
		tmpl, err := template.ParseFiles("error.html")
		if err != nil {
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	// TODO: Display the information in an HTML page through `template`

	title = snippet["title"].(string)
	id = item["id"].(string)
	ctitle = snippet["channelTitle"].(string)
	lcnt = num_handle(statistics["likeCount"].(string))
	vcnt = num_handle(statistics["viewCount"].(string))
	pa = time_handle(snippet["publishedAt"].(string))
	ccnt = num_handle(statistics["commentCount"].(string))

	filename = "index.html"
	tmpl, err := template.ParseFiles(filename)
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", YouTubePage)
	log.Fatal(http.ListenAndServe(":8085", nil))
}
