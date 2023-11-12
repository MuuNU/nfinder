package handle

import (
	"github.com/elastic/go-elasticsearch/v8"
	"html/template"
	"net/http"
	"nfinder/core"
)

var (
	elasticConfig    = elasticsearch.Config{Addresses: []string{"http://localhost:9200"}}
	elasticClient, _ = elasticsearch.NewClient(elasticConfig)
)

func HandlerMain(w http.ResponseWriter, r *http.Request) {
	print(r)
	processor := core.NewSearchProcessor(elasticClient)
	processor.SetSearchIndex("note-test")
	processor.SetSearchType("match")
	processor.SetSearchFields("content.en", "content.ru")
	processor.SetSearchArg("сосать")
	doc := processor.Run()
	ts, _ := template.ParseFiles("./templates/index.html")
	err := ts.Execute(w, doc)
	if err != nil {
		return
	}
}
