package core

type Document struct {
	Created string `json:"created"` // May be use a time format instead of string
	Title   struct {
		Default   string `json:"default"`
		Ru        string `json:"ru"`
		En        string `json:"en"`
		Language  string `json:"language"`
		Supported bool   `json:"supported"`
	} `json:"title"`
	Updated string `json:"updated"` // May be use a time format instead of string
	Content struct {
		Default   string `json:"default"`
		Ru        string `json:"ru"`
		En        string `json:"en"`
		Language  string `json:"language"`
		Supported bool   `json:"supported"`
	} `json:"content"`
	Tags   []string `json:"tags.tag"`
	NoteID string   `json:"note_id"`
}
