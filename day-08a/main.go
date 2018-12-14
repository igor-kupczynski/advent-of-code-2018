package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type input struct {
	tokens []int
	idx int
}

func (i *input) Next() int {
	if i.idx >= len(i.tokens) {
		log.Fatalf("Trying to read after end of file @ idx=%d\n", i.idx)
	}
	i.idx++
	return i.tokens[i.idx - 1]
}

type node struct {
	children []*node
	metadata []int
}

func (n *node) totalMetadata() int {
	sum := 0
	for _, ch := range n.children {
		sum += ch.totalMetadata()
	}
	for _, m := range n.metadata {
		sum += m
	}
	return sum
}

func readNode(input *input) node {
	noChildren := input.Next()
	noMetadata := input.Next()

	children := make([]*node, noChildren)
	for i := 0; i < noChildren; i++ {
		child := readNode(input)
		children[i] = &child
	}

	metadata := make([]int, noMetadata)
	for i := 0; i < noMetadata; i++ {
		metadata[i] = input.Next()
	}

	return node{children, metadata}
}

func main() {
	tokens := []int{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		for _, x := range strings.Split(s.Text(), " ") {
			var num int
			fmt.Sscanf(x, "%d", &num)
			tokens = append(tokens, num)
		}
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	input := input{tokens, 0}
	tree := readNode(&input)

	fmt.Println(tree.totalMetadata())

}
