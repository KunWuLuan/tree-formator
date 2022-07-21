package treeformator

import "fmt"

const (
	space = "  "
	last  = "|-"
	line  = "| "
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
	strs := []string{}
	dfs(NewTreeFormatorNodeInfo(r, nil, true), strs)
	for _, str := range strs {
		fmt.Println(str)
	}
}

func (r *TreeFormatorNodeInfo) Show() string {
	return r.node.Show()
}

func buildPrefix(r *TreeFormatorNodeInfo) string {
	root := r.root
	hasLast := false
	cur := ""
	for root != nil {
		if root.isLast {
			if !hasLast {
				cur = last + r.Show()
			} else {
				cur = space + r.Show()
			}
		}
		root = root.root
	}
	return cur
}

func dfs(r *TreeFormatorNodeInfo, str []string) {
	if r == nil {
		return
	}
	nextLevelNodes := r.NextLevel()
	str = append(str, buildPrefix(r))
	for i, child := range nextLevelNodes {
		dfs(NewTreeFormatorNodeInfo(child, r, i == len(nextLevelNodes)), str)
	}
}

type TreeFormatorBinTreeImpl struct {
	value string
	left  *TreeFormatorBinTreeImpl
	right *TreeFormatorBinTreeImpl
	root  *TreeFormatorBinTreeImpl
}

func (t *TreeFormatorBinTreeImpl) Show() string {
	return t.value
}

func (t *TreeFormatorBinTreeImpl) NextLevel() []TreeFormator {
	res := []TreeFormator{t.right}
	cur := t.right
	for cur != nil {
		res = append(res, cur)
		cur = cur.right
	}
	return res
}

func (t *TreeFormatorBinTreeImpl) Root() TreeFormator {
	return t.root
}

func BuildSampleTree() (root *TreeFormatorBinTreeImpl) {
	root = &TreeFormatorBinTreeImpl{
		"a",
		&TreeFormatorBinTreeImpl{"b", nil, nil, root},
		&TreeFormatorBinTreeImpl{"c", nil, nil, root},
		nil,
	}
	return
}
