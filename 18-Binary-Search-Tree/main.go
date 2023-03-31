package main

import (
	"log"
)

type bst struct {
	val   int
	left  *bst
	right *bst
}

func (b *bst) insert(val int) {
	if val < b.val {
		if b.left != nil {
			b.left.insert(val)
		} else {
			b.left = &bst{val: val}
		}
	}

	if val > b.val {
		if b.right != nil {
			b.right.insert(val)
		} else {
			b.right = &bst{val: val}
		}
	}
}

func (b *bst) search(val int) bool {
	if b == nil {
		return false
	}

	if val < b.val {
		return b.left.search(val)
	}

	if val > b.val {
		return b.right.search(val)
	}

	return true
}

func main() {
	root := &bst{val: 8}
	root.insert(6)
	root.insert(10)

	log.Println(root)
	log.Println(root.search(88))
	log.Println(root.search(6))
	log.Println(root.search(10))
	log.Println(root.search(101))
}
