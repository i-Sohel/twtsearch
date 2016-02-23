package main

import (
	"log"
	"net/http"
	"os"
	"html/template"
    "fmt"
    "encoding/json"
	//for extracting service credentials from VCAP_SERVICES
	//"github.com/cloudfoundry-community/go-cfenv"
)

type Ms struct {
     Tweets []struct {
		Message struct {
			//PostedTime time.Time `json:"postedTime"`
			Verb string `json:"verb"`
			Link string `json:"link"`
			Generator struct {
				DisplayName string `json:"displayName"`
				Link string `json:"link"`
			} `json:"generator"`
			Body string `json:"body"`
			FavoritesCount int `json:"favoritesCount"`
			ObjectType string `json:"objectType"`
			Actor struct {
				Summary string `json:"summary"`
				Image string `json:"image"`
				StatusesCount int `json:"statusesCount"`
				UtcOffset string `json:"utcOffset"`
				Languages []string `json:"languages"`
				PreferredUsername string `json:"preferredUsername"`
				DisplayName string `json:"displayName"`
				//PostedTime time.Time `json:"postedTime"`
				Link string `json:"link"`
				Verified bool `json:"verified"`
				FriendsCount int `json:"friendsCount"`
				TwitterTimeZone string `json:"twitterTimeZone"`
				FavoritesCount int `json:"favoritesCount"`
				ListedCount int `json:"listedCount"`
				ObjectType string `json:"objectType"`
				Links []struct {
					Rel string `json:"rel"`
					Href string `json:"href"`
				} `json:"links"`
				Location struct {
					DisplayName string `json:"displayName"`
					ObjectType string `json:"objectType"`
				} `json:"location"`
				ID string `json:"id"`
				FollowersCount int `json:"followersCount"`
			} `json:"actor"`
			Provider struct {
				DisplayName string `json:"displayName"`
				Link string `json:"link"`
				ObjectType string `json:"objectType"`
			} `json:"provider"`
			TwitterEntities struct {
				Urls []interface{} `json:"urls"`
				Hashtags []interface{} `json:"hashtags"`
				UserMentions []struct {
					Indices []int `json:"indices"`
					ScreenName string `json:"screen_name"`
					IDStr string `json:"id_str"`
					Name string `json:"name"`
					ID int `json:"id"`
				} `json:"user_mentions"`
				Symbols []interface{} `json:"symbols"`
			} `json:"twitter_entities"`
			TwitterLang string `json:"twitter_lang"`
			ID string `json:"id"`
			RetweetCount int `json:"retweetCount"`
			Object struct {
				Summary string `json:"summary"`
				//PostedTime time.Time `json:"postedTime"`
				Link string `json:"link"`
				ID string `json:"id"`
				ObjectType string `json:"objectType"`
			} `json:"object"`
		} `json:"message"`
	} `json:"tweets"`
}

const (
	DEFAULT_PORT = "8080"
    size = "5"
    DEF_URL = "" // API URL you got from Bluemix service
)


var templates = template.Must(template.ParseFiles(
  "templates/index.html",
  "templates/About.html",
))


func httpserve(w http.ResponseWriter, req *http.Request) {
  templates.ExecuteTemplate(w, "indexPage", nil)
}
// Get Twitter Serach results in JSON Format using GET request method
func getJSON(url string, target interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()
    return json.NewDecoder(r.Body).Decode(target)
}

func results(w http.ResponseWriter, req *http.Request){
    req.ParseForm()
    query := req.FormValue("key") // Get the keywords from the template
    var url = DEF_URL + query + "&size=" + size // Construct the request url. The size here represents the number of returned tweets which is 5 at the moment
    res := new(Ms)
    getJSON(url, res) // Get the data
    fmt.Println(res)
    // Return the data to the template
    templates.ExecuteTemplate(w,"indexPage", res)
}

func about(w http.ResponseWriter, req *http.Request){
    templates.ExecuteTemplate(w , "aboutPage", nil)
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}
        
	http.HandleFunc("/", httpserve)
    http.HandleFunc("/results", results)
    http.HandleFunc("/about", about)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	log.Printf("Starting app on port %+v\n", port)
	http.ListenAndServe(":"+port, nil)
}
