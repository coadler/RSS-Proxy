package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

// RSSData holds the structure of the JSON data for GetRSS
type RSSData struct {
	URL string `json:"url"`
}

// JSONError holds the structure for the JSON error response
type JSONError struct {
	ErrorCode string `json:"code"`
	Message   string `json:"message"`
}

var (
	redditUser string
	auth       string
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	if r := viper.GetString("reddit_user"); r != "" {
		redditUser = r
	} else {
		panic("reddit_user not set")
	}

	if a := viper.GetString("auth"); a != "" {
		auth = a
	} else {
		panic("auth not set")
	}
}

// Index responds with Hello World so it can easily be tested if the API is running
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "s-senpai, please don't hurt me ;_;\n")
}

// GetRSS is the handler function for /v1/get, which returns the requested RSS document
func GetRSS(w http.ResponseWriter, r *http.Request) {
	// super secret password authentication :eyes:
	if r.Header.Get("Authorization") != auth {
		fmt.Fprint(w, ":thonking:\n"+
			"get off my api you idiot\n"+
			"i've logged ur ip, prepare for ddos")
		fmt.Println("bad guy: " + r.RemoteAddr)
		return
	}

	client := &http.Client{}
	var data RSSData
	var reddit bool

	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &data)
	if err != nil {
		fmt.Println("Error unmarshaling json." + err.Error())
		http.Error(w, "Server error", 500) // todo: informational json errors
		return
	}

	if !strings.HasPrefix(data.URL, "http") {
		data.URL = "http://" + data.URL
		fmt.Println("Malformed URL received, fixing " + data.URL)
	}

	if strings.Contains(data.URL, "reddit.com") {
		reddit = true
		oldURL := data.URL
		// Force HTTPS on Reddit because it doesn't like HTTP on their API
		data.URL = strings.Replace(data.URL, "http://", "https://", 1)
		if oldURL != data.URL {
			fmt.Println("Reddit URL fixed. " + data.URL)
		}
	}

	req, err := http.NewRequest("GET", data.URL, nil)
	if err != nil {
		fmt.Println("Error creating request. " + err.Error())
		http.Error(w, "Server error", 500) // todo: informational json errors
		return
	}

	if reddit {
		// Set User-Agent to follow Reddit bot rules
		req.Header.Set("User-Agent", "GoLang:RSS-Proxy:v1.0.1 (by /u/"+redditUser)
	} else {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36")
	}

	// insert IP of whitehouse.gov so they heckin go to jail if they think they're clever enough to look thru the headers
	req.Header.Add("X-Forwarded-For", "104.126.136.25")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error retrieving RSS. " + err.Error())
		http.Error(w, "Server error", 500) // todo: informational json errors
		return
	}

	defer resp.Body.Close()
	rss, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body. " + err.Error())
		http.Error(w, "Server error", 500) // todo: informational json errors
		return
	}

	fmt.Fprint(w, string(rss))
}
