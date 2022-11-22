package treeformator

import "fmt"

const (
	space   = "  "
	last    = "\\-"
	line    = "| "
	notlast = "|-"
)

type TreeFormatorNodeInfo struct {
	node   TreeFormator
	root   *TreeFormatorNodeInfo
	isLast bool
}

func (r TreeFormatorNodeInfo) NextLevel() []TreeFormator {
	if r.node == nil {
		return nil
	}
	return r.node.NextLevel()
}

func NewTreeFormatorNodeInfo(r TreeFormator, root *TreeFormatorNodeInfo, isLast bool) *TreeFormatorNodeInfo {
	return &TreeFormatorNodeInfo{r, root, isLast}
}

func ShowTreeStruct(r TreeFormator) {
	strs := dfs(NewTreeFormatorNodeInfo(r, nil, true))
	for _, str := range strs {
		fmt.Println(str)
	}
}

func (r *TreeFormatorNodeInfo) Show() string {
	return r.node.Show()
}

func buildPrefix(r *TreeFormatorNodeInfo) string {
	if r == nil {
		return ""
	}
	cur := r.Show()
	if r.root == nil {
		return cur
	} else {
		if r.isLast {
			cur = last + cur
		} else {
			cur = notlast + cur
		}
	}
	curNode := r.root
	root := curNode.root
	for root != nil {
		if curNode.isLast {
			cur = space + cur
		} else {
			cur = line + cur
		}
		curNode = root
		root = curNode.root
	}
	return cur
}

func dfs(r *TreeFormatorNodeInfo) (str []string) {
	if r == nil {
		return []string{}
	}
	nextLevelNodes := r.NextLevel()
	str = append(str, buildPrefix(r))
	for i, child := range nextLevelNodes {
		str = append(str, dfs(NewTreeFormatorNodeInfo(child, r, i == len(nextLevelNodes)-1))...)
	}
	return
}
