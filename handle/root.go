package handle

import (
	"github.com/elastic/go-elasticsearch/v8"
	"html/template"
	"net/http"
	"nfinder/core"
)

var (
	elasticConfig      = elasticsearch.Config{Addresses: []string{"http://localhost:9200"}}
	elasticClient, err = elasticsearch.NewClient(elasticConfig)
)

func HandlerMain(w http.ResponseWriter, r *http.Request) {
	processor := core.NewSearchProcessor(elasticClient)
	processor.SetSearchIndex("note-test")
	processor.SetSearchType("test")
	processor.SetSearchFields("test", "tes2", "test3")
	processor.Run()
	ts, _ := template.ParseFiles("./templates/index.html")
	ts.Execute(w, "")
}
