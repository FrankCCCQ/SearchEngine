package types

import "github.com/RoaringBitmap/roaring"

type Document struct {
	DocId int64  `json:"doc_id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Data2Starrocks struct {
	DocId int64   `json:"doc_id"`
	Url   string  `json:"url"`
	Title string  `json:"title"`
	Desc  string  `json:"desc"`
	Score float64 `json:"score"`
}

type Task struct {
	Columns    []string `json:"columns"`
	BiTable    string   `json:"bi_title"`
	SourceType int      `json:"source_type"`
}

type DictTrieTree struct {
	Value string `json:"value"`
	Score int64  `json:"score"`
}

type InvertedIndexValue struct {
	Token        string        `json:"token"`
	PostingsList *PostingsList `json:"posting_list"`
	DocCount     int64         `json:"doc_count"`
	PostingCount int64         `json:"posting_count"`
	TermValues   *TermValue    `json:"term_values"`
}

type TermValue struct {
	DocCount int64 `json:"doc_count"`
	Offset   int64 `json:"offset"`
	Size     int64 `json:"size"`
}

type PostingsList struct {
	Term      string          `json:"term"`
	Position  []int64         `json:"position"` //标红?
	TermCount int64           `json:"term_count"`
	DocIds    *roaring.Bitmap `json:"doc_ids"`
}
