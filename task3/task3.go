package main

import (
	"fmt"
	"net/http"
	"io"
	"html/template"
	"encoding/json"
)

const url = "https://everyday-cc.github.io/ai/api/news/"

type newsItem struct {
	Page   int `json:"page"`
   Value   []struct {
	Path            string `json:"path"`
	Title           string `json:"title"`
	Excerpt         string `json:"excerpt"`
	SourceURL       any    `json:"sourceUrl"`
	WebURL          string `json:"webUrl"`
	OriginalURL     string `json:"originalUrl"`
	FeaturedContent any    `json:"featuredContent"`
	Highlight       any    `json:"highlight"`
	Heat            int    `json:"heat"`
	Tags            any    `json:"tags"`
	Images          []struct {
		URL         string `json:"url"`
		Width       int    `json:"width"`
		Height      int    `json:"height"`
		Title       any    `json:"title"`
		Attribution any    `json:"attribution"`
		IsCached    bool   `json:"isCached"`
	} `json:"images"`
	Content           string `json:"content"`
	Type              string `json:"type"`
	AmpWebURL         any    `json:"ampWebUrl"`
	CdnAmpWebURL      any    `json:"cdnAmpWebUrl"`
	PublishedDateTime string `json:"publishedDateTime"`
	UpdatedDateTime   any    `json:"updatedDateTime"`
	Provider          struct {
		Name   string `json:"name"`
		Domain string `json:"domain"`
		Images []struct {
			URL         string `json:"url"`
			Width       int    `json:"width"`
			Height      int    `json:"height"`
			Title       any    `json:"title"`
			Attribution any    `json:"attribution"`
			IsCached    bool   `json:"isCached"`
		} `json:"images"`
		Publishers any `json:"publishers"`
		Authors    any `json:"authors"`
	} `json:"provider"`
	Locale     string   `json:"locale"`
	Categories []string `json:"categories"`
	Topics     []string `json:"topics"`
   }
   NextPage int `json:"nextPage"`
}

func main() {

	http.HandleFunc("/", aiNews)
    	http.ListenAndServe(":8080", nil)

}

func aiNews ( w http.ResponseWriter, r *http.Request) {

	var data newsItem
	req, _ := http.NewRequest("GET", url, nil)

	// Create Template
	t, err := template.New("foo").Parse(`<table><th>AI News</th>{{range .}}<tr onMouseOver="this.style.backgroundColor='#F8F8F8'" onMouseOut="this.style.backgroundColor='#FFFFFF'"><td> <a href="{{.WebURL}}">{{.Title}}</a><br/><table><tr><td><span style="font-size:small">{{.Excerpt}}</span></br><span style="font-size:small"><a href="{{.Provider.Domain}}">{{.Provider.Name}}</a> | {{.Topics}} | {{.PublishedDateTime}} </span><br/></td></tr></table></td></tr>{{end}}</table>`)

	// Create client
        res, _ := http.DefaultClient.Do(req)
        defer res.Body.Close()

	// Get news data
        body, _:= io.ReadAll(res.Body)
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Printf("Could not unmarshal json: %s\n", err)
		return
	}

	err = t.Execute(w, data.Value)

}

