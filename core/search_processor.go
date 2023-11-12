package core

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"strings"
)

type SearchInterface interface {
	SetSearchIndex(arg string)
	SetSearchArg(arg string)
	SetSearchType(arg string)
	SetSearchFields(arg ...string)
	SetPaginationEnabled(arg bool)
	SetPaginationCount(arg int)
	SetAnalyzerEnabled(arg bool)
	SetAnalyzerType(arg string)
	SetFuzzinessEnabled(arg bool)
	SetFuzzinessCount(arg int)
	Run()
}
type SearchProcessor struct {
	client       *elasticsearch.Client
	searchIndex  string
	searchArg    string
	searchType   string
	searchFields []string
	// TODO: Create class for search types
	paginationenabled bool
	paginationcount   int
	analyzerenabled   bool
	analyzertype      string
	// TODO: Create class for analyzers
	fuzzinessenabled bool
	fuzzinesscount   int
}

func NewSearchProcessor(client *elasticsearch.Client) *SearchProcessor {

	sp := new(SearchProcessor)
	sp.client = client
	sp.searchIndex = ""
	sp.searchArg = ""
	sp.searchType = ""
	sp.paginationenabled = true
	sp.paginationcount = 20
	sp.analyzerenabled = false
	sp.analyzertype = ""
	sp.fuzzinessenabled = false
	sp.fuzzinesscount = 0

	return sp
}

func (receiver SearchProcessor) SetSearchIndex(arg string) {
	receiver.searchIndex = arg
}
func (receiver SearchProcessor) SetSearchArg(arg string) {
	receiver.searchArg = arg
}
func (receiver SearchProcessor) SetSearchType(arg string) {
	receiver.searchType = arg
}
func (receiver SearchProcessor) SetSearchFields(arg ...string) {

	for _, e := range arg {
		receiver.searchFields = append(receiver.searchFields, e)
	}

}
func (receiver SearchProcessor) SetPaginationEnabled(arg bool) {
	receiver.paginationenabled = arg
}
func (receiver SearchProcessor) SetPaginationCount(arg int) {
	receiver.paginationcount = arg
}
func (receiver SearchProcessor) SetAnalyzerEnabled(arg bool) {
	receiver.analyzerenabled = arg
}
func (receiver SearchProcessor) SetAnalyzerType(arg string) {
	receiver.analyzertype = arg
}
func (receiver SearchProcessor) SetFuzzinessEnabled(arg bool) {
	receiver.fuzzinessenabled = arg
}
func (receiver SearchProcessor) SetFuzzinessCount(arg int) {
	receiver.fuzzinesscount = arg
}

func (receiver SearchProcessor) Run() Document {

	var arr []string
	var query string
	if receiver.searchType == "match" {
		for i, _ := range receiver.searchFields {
			if i >= 1 {
				log.Fatal("Error in processing query. Selected type \"Match\", " +
					"but got more than one fields in in searchFields")
			}
		}
		query = fmt.Sprintf(`{ 
			"query": { 
				"match": { 
					"%s": "%s"
				} 
			} 
		}`, receiver.searchFields, receiver.searchArg[0])
	} else if receiver.searchType == "multimatch" {
		fields := ""
		for i, e := range receiver.searchFields {
			if i == 0 {
				fields = fields + "\"" + e + "\""
			} else {
				fields = fields + "," + "\"" + e + "\""
			}
		}
		qstr := fmt.Sprintf("\"query\": %s", receiver.searchArg)
		fstr := fmt.Sprintf("\"fields\": [%s]", fields)
		query = fmt.Sprintf(`
			{
			  "query": {
				"multi_match": {
				%s,
				%s
				}
			  }, 
			}`, qstr, fstr)
	}
	if receiver.searchType == "title" {
		arr = append(arr,
			"title.default",
			"title.en",
			"title.ru",
		)
	}

	receiver.client.Search(
		receiver.client.Search.WithIndex(receiver.searchIndex),
		receiver.client.Search.WithPretty(),
		receiver.client.Search.WithSource(arr...),
		receiver.client.Search.WithBody(strings.NewReader(query)),
	)
}
