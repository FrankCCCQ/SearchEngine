package trie

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTrieTree(t *testing.T) {
	t1 := NewTrie()
	t1.Insert("hello")
	t1.Insert("world")
	fmt.Println(t1)
	t1.Traverse()

	t2 := NewTrie()
	t2.Insert("hello")
	t2.Insert("golang")
	t2.Insert("program")
	fmt.Println(t2)
	t2.Traverse()

	t1.Merge(t2)
	fmt.Println(t1)
	t1.Traverse()

	r := t1.FindAllByPrefix("he")
	fmt.Println(r)
}

func TestBinaryTrieTree(t *testing.T) {
	t1 := NewTrie()
	t1.Insert("哈罗")
	t1.Insert("狗")
	t1.Traverse()
	rootByte, err := t1.Root.Children.MarshalJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("rootByte:", string(rootByte))

	// t2 := NewTrie()
	// t2.Root = NewTrieNode()
	// // 后面的节点地址没有分配,不能直接赋值
	// err = t2.Root.Children.UnmarshalJSON(rootByte)
	// if err != nil {
	// 	fmt.Println("unmarshalJSON err", err)
	// }
	// fmt.Println("starting", t2)
}

func TestJsonMarshal(t *testing.T) {
	a := "{\"则\":{\"is_end\":false,\"children\":{\"美\":{\"is_end\":true,\"children\":{}}}},\"啊\":{\"is_end\":false,\"children\":{\"啊\":{\"is_end\":false,\"children\":{\"啊\":{\"is_end\":true,\"children\":{}}}}}},\"成\":{\"is_end\":false,\"children\":{\"型\":{\"is_end\":true,\"children\":{}}}}}\n"
	b := []byte(a)
	replaced := bytes.Replace(b, []byte("children"), []byte("children_recall"), -1)
	fmt.Println(string(replaced))
	node, err := ParseTrieNode(string(replaced))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(node)
	trie := NewTrie()
	trie.Root = node
	fmt.Println("traverse")
	trie.TraverseRecall()
	alist := trie.FindAllByPrefix("成")
	fmt.Println(alist)

}
