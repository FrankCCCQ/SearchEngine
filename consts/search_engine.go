package consts

const LayoutTimeFormat = "2006-01-02 15:04:05"

const (
	TermBucket     = "term"
	InvertedBucket = "inverted"
	TrieTreeBucket = "trie_tree"

	EngineBufSize         = 10000
	ForwardCountInitValue = "0"
)

const (
	DataSourceCSV = iota + 1
	DataSourceCrawl
)

const (
	BatchCreateSize = 1000
	BatchSize       = 100
)
