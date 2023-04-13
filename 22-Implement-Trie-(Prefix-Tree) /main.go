package main

import (
	"log"
)

func main() {
	obj := Constructor()
	obj.Insert("hello")
	log.Println(obj.Search("hello"))
	log.Println(obj.Search("hello1"))
	log.Println(obj.StartsWith("hell"))
	log.Println(obj.StartsWith("hell1"))
}

type node struct {
	children    map[string]*node
	isEndOfWord bool
}

func newNode() *node {
	return &node{
		children:    make(map[string]*node),
		isEndOfWord: false,
	}
}

type Trie struct {
	root *node
}

func Constructor() Trie {
	return Trie{
		root: newNode(),
	}
}

func (t *Trie) Insert(word string) {
	cur := t.root
	runes := []rune(word)

	for _, r := range runes {
		if _, exist := cur.children[string(r)]; !exist {
			cur.children[string(r)] = newNode()
		}
		cur = cur.children[string(r)]
	}
	cur.isEndOfWord = true
}

func (t *Trie) Search(word string) bool {
	cur := t.root
	runes := []rune(word)

	for _, r := range runes {
		if _, exist := cur.children[string(r)]; !exist {
			return false
		}
		cur = cur.children[string(r)]
	}
	return cur.isEndOfWord
}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t.root
	runes := []rune(prefix)

	for _, r := range runes {
		if _, exist := cur.children[string(r)]; !exist {
			return false
		}
		cur = cur.children[string(r)]
	}
	return true
}
