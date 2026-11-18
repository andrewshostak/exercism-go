package treebuilding

import (
	"cmp"
	"errors"
	"slices"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	slices.SortFunc(records, func(i, j Record) int {
		return cmp.Compare(i.ID, j.ID)
	})

	childrenOf := make(map[int][]int)
	for i, record := range records {
		if i == 0 && (record.Parent != 0 || record.ID != 0) {
			return nil, errors.New("root node is invalid")
		}

		if i == 0 {
			continue
		}

		if i != record.ID || record.ID <= record.Parent {
			return nil, errors.New("non continuous")
		}

		childrenOf[record.Parent] = append(childrenOf[record.Parent], record.ID)
	}

	return createNode(&Node{ID: 0}, childrenOf), nil
}

func createNode(node *Node, childrenOf map[int][]int) *Node {
	childIDs := childrenOf[node.ID]
	if len(childIDs) != 0 {
		node.Children = make([]*Node, 0, len(childIDs))
		for _, id := range childIDs {
			node.Children = append(node.Children, createNode(&Node{ID: id}, childrenOf))
		}
	}
	return node
}
