package main

func main() {
	//input := map[int][]int{
	//	1: {2, 4},
	//	2: {1, 3},
	//	3: {2, 4},
	//	4: {1, 3},
	//}

}

type Node struct {
	Val       int
	Neighbors []*Node
}

func newNode(val int) *Node {
	return &Node{
		Val:       val,
		Neighbors: []*Node{},
	}
}

func cloneGraph(node *Node) *Node {
	nodeMap := make(map[*Node]*Node)
	return clone(node, nodeMap)
}

func clone(node *Node, nodeMap map[*Node]*Node) *Node {
	if node == nil {
		return node
	}

	val, exist := nodeMap[node]
	if exist {
		return val
	}

	copiedNode := newNode(node.Val)
	nodeMap[node] = copiedNode

	for _, n := range node.Neighbors {
		copiedNode.Neighbors = append(copiedNode.Neighbors, clone(n, nodeMap))
	}

	return copiedNode
}
