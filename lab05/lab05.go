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
	// data := PageData{
	// 	Title:        title,
	// 	Id:           id,
	// 	ChannelTitle: ctitle,
	// 	LikeCount:    lcnt,
	// 	ViewCount:    vcnt,
	// 	PublishedAt:  pa,
	// 	CommentCount: ccnt,
	// }
	// TODO: Get API token from .env file
	err := godotenv.Load() // Load environment variable from .env file
	if err != nil {
		fmt.Println("========ERROR==========9\n")
		http.ServeFile(w, r, "error.html")
		return
	}
	password := os.Getenv("YOUTUBE_API_KEY") // Get value from environment variable
	// fmt.Printf("password: %s\n", password)

	// TODO: Get video ID from URL query `v`
	fmt.Println("?" + r.URL.String())
	if r.URL.String() == "/favicon.ico" {
		http.NotFound(w, r)
		return
	}
	key := "v"
	var videoid string = r.URL.Query().Get(key)
	fmt.Println("?" + videoid)

	if videoid == "" {
		http.ServeFile(w, r, "error.html")
		return
	}
	// fmt.Printf("video id: %s\n", videoid)

	// TODO: Get video information from YouTube API
	url := "https://www.googleapis.com/youtube/v3/videos?key=" + password + "&id=" + videoid + "&part=" + "statistics,snippet"
	resp, err := http.Get(url) // resp is the response of the url
	if err != nil {
		http.ServeFile(w, r, "error.html")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.ServeFile(w, r, "error.html")
		return
	}
	// fmt.Println(string(body))

	// TODO: Parse the JSON response and store the information into a struct
	var m map[string]interface{}
	err1 := json.Unmarshal(body, &m)
	if err1 != nil {
		http.ServeFile(w, r, "error.html")
		return
	}
	items := m["items"].([]interface{})
	if len(items) == 0 {
		http.ServeFile(w, r, "error.html")
		return
	}
	item := items[0].(map[string]interface{})
	if err1 != nil {
		http.ServeFile(w, r, "error.html")
		return
	}
	statistics := item["statistics"].(map[string]interface{})
	snippet := item["snippet"].(map[string]interface{})

	if err != nil {
		http.ServeFile(w, r, "error.html")
		return
	}
	fmt.Println("######################")

	// TODO: Display the information in an HTML page through `template`

	title = snippet["title"].(string)
	// fmt.Println(title)
	id = item["id"].(string)
	// fmt.Println(id)
	ctitle = snippet["channelTitle"].(string)
	// fmt.Println(ctitle)
	lcnt = num_handle(statistics["likeCount"].(string))
	// fmt.Println(lcnt)
	vcnt = num_handle(statistics["viewCount"].(string))
	// fmt.Println(vcnt)
	pa = time_handle(snippet["publishedAt"].(string))
	// fmt.Println(pa)
	ccnt = num_handle(statistics["commentCount"].(string))
	// fmt.Println(ccnt)
	data := PageData{
		Title:        title,
		Id:           id,
		ChannelTitle: ctitle,
		LikeCount:    lcnt,
		ViewCount:    vcnt,
		PublishedAt:  pa,
		CommentCount: ccnt,
	}
	filename = "index.html"
	tmpl, err := template.ParseFiles(filename)
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("errorerrorerror")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("??????????????")
	return
}

func main() {
	http.HandleFunc("/", YouTubePage)
	log.Fatal(http.ListenAndServe(":8085", nil))
}
