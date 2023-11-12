package core

// Response
// Structure for serialization answer from ElasticSearch
// Provide all fields from shards info
// .
type Response struct {
	Took     int      `json:"took"`
	TimedOut bool     `json:"timed_out"`
	Shards   Shards   `json:"_shards"`
	HitsInfo HitsInfo `json:"hits"`
}

// Shards
// Structure for serialization answer from ElasticSearch
// Provide all fields from shards info
// ._shards
type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

// HitsInfo
// Structure for serialization answer from ElasticSearch
// Provide all fields from shards info
// .hits
type HitsInfo struct {
	Total struct {
		Value    int    `json:"value"`
		Relation string `json:"relation"`
	} `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hits  `json:"hits"`
}

// Hits
// Structure for serialization answer from ElasticSearch
// Provide all fields from shards info
// .hits.hits
type Hits struct {
	Index  string   `json:"_index"`
	ID     string   `json:"_id"`
	Score  float64  `json:"_score"`
	Source Document `json:"_source"`
}
