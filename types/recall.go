package types

type SearchItem struct {
	DocId        int64   `json:"doc_id"`
	Content      string  `json:"content"`
	Title        string  `json:"title"`
	Score        float64 `json:"score"`
	DocCount     int64   `json:"doc_count"`
	ContentScore float64 `json:"content_score"`
}

type SearchItemList []*SearchItem

func (ds SearchItemList) Len() int {
	return len(ds)
}

func (ds SearchItemList) Less(i, j int) bool {
	return ds[i].Score < ds[j].Score
}
func (ds SearchItemList) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}

type queryTokenHash struct {
	token         string
	invertedIndex *InvertedIndexValue
	fetchPostings *PostingsList
}

type searchCursor struct {
	doc     *PostingsList
	current *PostingsList
}

type phraseCursor struct {
	positions []int64
	base      int64
	current   *int64
	index     int
}
