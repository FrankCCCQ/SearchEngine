package trie

import (
	"fmt"

	cmap "github.com/orcaman/concurrent-map/v2"
)

type TrieNode struct {
	IsEnd          bool                                  `json:"is_end"`   //标记该节点是否为一个单词的末尾
	Children       cmap.ConcurrentMap[string, *TrieNode] `json:"children"` // 存储子节点的指针
	ChildrenRecall map[string]*TrieNode                  `json:"children_recall"`
}

func NewTrieNode() *TrieNode {
	m := cmap.New[*TrieNode]()
	return &TrieNode{
		IsEnd:    false,
		Children: m,
	}
}

type Trie struct {
	Root *TrieNode `json:"root"` //存储root节点
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(),
	}
}

func (trie *Trie) Insert(word string) {
	words := []rune(word)
	node := trie.Root
	for i := 0; i < len(words); i++ {
		c := string(words[i])
		if _, ok := node.Children.Get(c); !ok {
			node.Children.Set(c, NewTrieNode())
		}
		node, _ = node.Children.Get(c)
	}
	node.IsEnd = true
}

// Search 查询word是否存在
func (trie *Trie) Search(word string) bool {
	words := []rune(word)
	node := trie.Root
	for i := 0; i < len(words); i++ {
		c := string(words[i])
		if _, ok := node.Children.Get(c); !ok {
			return false
		}
		node, _ = node.Children.Get(c)
	}
	// IsEnd 表示word已经结束 否则 查找的word是存储的子串
	return node.IsEnd
}

// StartWith 判读是否包含前缀
func (trie *Trie) StartWith(prefix string) bool {
	prefixs := []rune(prefix)
	node := trie.Root
	for i := 0; i < len(prefixs); i++ {
		c := string(prefixs[i])
		if _, ok := node.Children.Get(c); !ok {
			return false
		}
		node, _ = node.Children.Get(c)
	}
	return true
}

func (trie *Trie) FindAllByPrefix(prefix string) []string {
	prefixs := []rune(prefix)
	node := trie.Root
	for i := 0; i < len(prefixs); i++ {
		c := string(prefixs[i])
		if _, ok := node.Children.Get(c); !ok {
			return nil
		}
		node, _ = node.Children.Get(c)
	}
	words := make([]string, 0)
	trie.dfs(node, prefix, &words)
	return words
}

func (trie *Trie) dfs(node *TrieNode, word string, words *[]string) {
	if node.IsEnd {
		*words = append(*words, word)
	}
	for c, child := range node.Children.Items() {
		trie.dfs(child, word+c, words)
	}
}

func (trie *Trie) Merge(other *Trie) {
	if other == nil {
		return
	}
	var MergeNode func(n1, n2 *TrieNode)
	MergeNode = func(n1, n2 *TrieNode) {
		for c, child := range n2.Children.Items() {
			if v, ok := n1.Children.Get(c); ok {
				MergeNode(v, child)
			} else {
				n1.Children.Set(c, child)
			}
		}
		n1.IsEnd = n1.IsEnd || n2.IsEnd
	}
	MergeNode(trie.Root, other.Root)
}

func traverse(node *TrieNode, prefix string) {
	if node.IsEnd {
		fmt.Println(prefix)
	}
	for c, child := range node.Children.Items() {
		traverse(child, prefix+c)
	}
}

func (trie *Trie) Traverse() {
	traverse(trie.Root, "")
}
