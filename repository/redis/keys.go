package redis

import (
	"fmt"
	"time"
)

var InvertedIndexDbPathKeys = []string{InvertedIndexDbPathDayKey, InvertedIndexDbPathMonthKey, InvertedIndexDbPathSeasonKey}

var TrieTreeDbPathKey = []string{
	TrieTreeDbPathDayKey, TrieTreeDbPathMonthKey, TrieTreeDbPathSeasonKey,
}

const (
	InvertedIndexDbPathDayKey    = "index_platform:inverted_index:day"
	InvertedIndexDbPathMonthKey  = "index_platform:inverted_index:month:%s"
	InvertedIndexDbPathSeasonKey = "index_platform:inverted_index:season:%s"

	TrieTreeDbPathDayKey    = "index_platform:trie_tree:day"
	TrieTreeDbPathMonthKey  = "index_platform:trie_tree:month:%s"
	TrieTreeDbPathSeasonKey = "index_platform:trie_tree:season:%s"

	QueryTokenDocIds = "query_doc_id:%s"
	UserQueryToken   = "query_token:%d"
)

const (
	QueryTokenDocIdsDefaultTimeout = 10 * time.Minute
)

func getQueryTokenDocsIdsKey(term string) string {
	return fmt.Sprintf(QueryTokenDocIds, term)
}

func getUserQueryTokenKey(userId int64) string {
	return fmt.Sprintf(UserQueryToken, userId)
}

func getAllDbPaths(key string) string {
	return fmt.Sprintf(key, "*")
}

func GetInvertedIndexDbPathMonthKey(month string) string {
	return fmt.Sprintf(InvertedIndexDbPathMonthKey, month)
}

func GetInvertedIndexDbPathSeasonKey(season string) string {
	return fmt.Sprintf(InvertedIndexDbPathSeasonKey, season)
}
