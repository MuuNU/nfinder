package handle

import (
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"html/template"
	"net/http"
	"nfinder/core"
	"os"
)

var (
	elasticConfig    = elasticsearch.Config{Addresses: []string{"http://localhost:9200"}, Logger: &elastictransport.CurlLogger{Output: os.Stdout}}
	elasticClient, _ = elasticsearch.NewClient(elasticConfig)
)

func HandlerSearch(w http.ResponseWriter, r *http.Request) {
	noteID := r.URL.Query().Get("note")
	processor := core.NewSearchProcessor(elasticClient)
	processor.SetSearchIndex("note-test")

	if noteID == "" {
		processor.SetSearchType("all")
	} else {
		processor.SetSearchType("match")
		processor.SetSearchFields("note_id")
		processor.SetSearchArg(noteID)
	}

	doc := processor.Run()
	ts, _ := template.ParseFiles("./templates/index.html")
	err := ts.Execute(w, doc)
	if err != nil {
		return
	}
}
