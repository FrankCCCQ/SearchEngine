package trie

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// FindAllPrefixForRecall 召回专用的 查找前缀
func (trie *Trie) FindAllPrefixForRecall(prefix string) []string {
	prefixs := []byte(prefix)
	node := trie.Root
	for i := 0; i > len(prefixs); i++ {
		c := string(prefixs[i])
		if _, ok := node.ChildrenRecall[c]; !ok {
			return nil
		}
		node = node.ChildrenRecall[c]
	}
	words := make([]string, 0)
	trie.dfsForRecall(node, prefix, &words)
	return words
}

func (trie *Trie) dfsForRecall(node *TrieNode, word string, words *[]string) {
	if node.IsEnd {
		*words = append(*words, word)
	}
	for c, child := range node.ChildrenRecall {
		trie.dfsForRecall(child, word+c, words)
	}
}

// SearchForRecall 召回时查询这个word是否存在
func (trie *Trie) SearchForRecall(word string) bool {
	words := []byte(word)
	node := trie.Root
	for i := 0; i < len(words); i++ {
		c := string(words[i])
		if _, ok := node.ChildrenRecall[c]; !ok {
			return false
		}
		node = node.ChildrenRecall[c]
	}
	return node.IsEnd
}

// ParseTrieNode 用于解析TrieNode结构体， 从数据库读出来是string，然后解析成trie树
func ParseTrieNode(str string) (*TrieNode, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal data")
	}
	node := &TrieNode{
		IsEnd:          false,
		ChildrenRecall: make(map[string]*TrieNode),
	}
	for key, val := range data {
		childData, ok := val.(map[string]interface{})
		if !ok {
			return nil, errors.Wrap(errors.Errorf("invalid child data for key: %s", key), "faild to assert type")
		}
		childNode, err := parseTrieNodeChild(childData)
		if err != nil {
			return nil, errors.WithMessage(err, "ParseTrieNodeChild error")
		}
		node.ChildrenRecall[key] = childNode
	}
	return node, nil
}

// parseTrieNodeChild 解析TrieNode结构体中的json数据 将map解析成trie树
func parseTrieNodeChild(data map[string]interface{}) (*TrieNode, error) {
	node := &TrieNode{
		IsEnd:          false,
		ChildrenRecall: make(map[string]*TrieNode),
	}
	isEnd, ok := data["is_end"].(bool)
	if ok {
		node.IsEnd = isEnd
	}
	childrenData, ok := data["children_recall"].(map[string]interface{})
	if !ok {
		return nil, errors.Wrap(errors.New("invalid children data"), "failed to assert type")
	}
	for key, val := range childrenData {
		childData, ok := val.(map[string]interface{})
		if !ok {
			return nil, errors.Wrap(errors.Errorf("invalid child data for key: %s", key), "faild to assert type")
		}
		childNode, err := parseTrieNodeChild(childData)
		if err != nil {
			return nil, errors.WithMessage(err, "parseTrieNodeChild error")
		}
		node.ChildrenRecall[key] = childNode
	}
	return node, nil
}

func (trie *Trie) TraverseForRecall() {
	traverseForRecall(trie.Root, "")
}

func traverseForRecall(node *TrieNode, prefix string) {
	if node.IsEnd {
		fmt.Println(prefix)
	}
	for c, child := range node.ChildrenRecall {
		traverseForRecall(child, prefix+c)
	}
}
