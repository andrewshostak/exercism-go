package tree

import (
	"errors"
	"sort"
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

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make([]*Node, len(records))
	for i, record := range records {
		nodes[i] = &Node{ID: record.ID}

		if i == 0 && (record.Parent != 0 || record.ID != 0) {
			return nil, errors.New("root node is invalid")
		}

		if i == 0 {
			continue
		}

		if i != record.ID || record.ID <= record.Parent {
			return nil, errors.New("records are incorrect")
		}

		if parent := nodes[record.Parent]; parent != nil {
			parent.Children = append(parent.Children, nodes[i])
		}
	}

	return nodes[0], nil
}
