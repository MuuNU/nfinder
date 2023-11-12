package core

import (
	"bytes"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"io"
	"log"
)

type (
	SearchInterface interface {
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
)
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

func (receiver *SearchProcessor) SetSearchIndex(arg string) {
	receiver.searchIndex = arg
}
func (receiver *SearchProcessor) SetSearchArg(arg string) {
	receiver.searchArg = arg
}
func (receiver *SearchProcessor) SetSearchType(arg string) {
	receiver.searchType = arg

}
func (receiver *SearchProcessor) SetSearchFields(arg ...string) {

	for _, e := range arg {
		receiver.searchFields = append(receiver.searchFields, e)
	}

}
func (receiver *SearchProcessor) SetPaginationEnabled(arg bool) {
	receiver.paginationenabled = arg
}
func (receiver *SearchProcessor) SetPaginationCount(arg int) {
	receiver.paginationcount = arg
}
func (receiver *SearchProcessor) SetAnalyzerEnabled(arg bool) {
	receiver.analyzerenabled = arg
}
func (receiver *SearchProcessor) SetAnalyzerType(arg string) {
	receiver.analyzertype = arg
}
func (receiver *SearchProcessor) SetFuzzinessEnabled(arg bool) {
	receiver.fuzzinessenabled = arg
}
func (receiver *SearchProcessor) SetFuzzinessCount(arg int) {
	receiver.fuzzinesscount = arg
}

type Match struct {
	Query struct {
		Match struct {
			NoteID string `json:"note_id"`
		} `json:"match"`
	} `json:"query"`
}

type Multimatch struct {
	Query struct {
		MultiMatch struct {
			Query  string   `json:"query"`
			Fields []string `json:"fields"`
		} `json:"multi_match"`
	} `json:"query"`
}

func (receiver SearchProcessor) Run() []Hits {

	var arr []string
	//var query string
	var jsonData []byte
	var response Response
	//var respRaw *esapi.Response
	if receiver.searchType == "all" {
		jsonData = []byte("")
	}
	if receiver.searchType == "match" {

		for i, _ := range receiver.searchFields {
			if i >= 1 {
				log.Fatal("Error in processing query. Selected type \"Match\", " +
					"but got more than one fields in in searchFields")
			}
		}
		var q Match
		q.Query.Match.NoteID = receiver.searchArg
		jsonData, _ = json.Marshal(q)

	}
	if receiver.searchType == "multimatch" {
		var q Multimatch
		q.Query.MultiMatch.Fields = receiver.searchFields
		q.Query.MultiMatch.Query = receiver.searchArg
		jsonData, _ = json.Marshal(q)
	}
	respRaw, _ := receiver.client.Search(
		receiver.client.Search.WithIndex(receiver.searchIndex),
		receiver.client.Search.WithPretty(),
		receiver.client.Search.WithSource(arr...),
		receiver.client.Search.WithBody(bytes.NewReader(jsonData)),
	)
	responceStr, _ := io.ReadAll(respRaw.Body)
	print(string(responceStr))
	json.Unmarshal(responceStr, &response)

	return response.HitsInfo.Hits

}
