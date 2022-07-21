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

type TreeFormatorBinTreeImpl struct {
	Value interface{}
	Left  *TreeFormatorBinTreeImpl
	Right *TreeFormatorBinTreeImpl
	root  *TreeFormatorBinTreeImpl
}

func (t *TreeFormatorBinTreeImpl) Show() string {
	return fmt.Sprintf("%+v", t.Value)
}

func (t *TreeFormatorBinTreeImpl) NextLevel() []TreeFormator {
	if t.Right == nil {
		return []TreeFormator{}
	}
	res := []TreeFormator{}
	if t.Left != nil {
		res = append(res, t.Left)
	}
	if t.Right != nil {
		res = append(res, t.Right)
	}
	return res
}

func (t *TreeFormatorBinTreeImpl) Root() TreeFormator {
	return t.root
}

func (t *TreeFormatorBinTreeImpl) SetLeft(obj interface{}) TreeFormator {
	t.Left = &TreeFormatorBinTreeImpl{obj, nil, nil, t}
	return t.Left
}

func (t *TreeFormatorBinTreeImpl) SetRight(obj interface{}) TreeFormator {
	t.Right = &TreeFormatorBinTreeImpl{obj, nil, nil, t}
	return t.Right
}

func BuildSampleTree() (root *TreeFormatorBinTreeImpl) {
	root = &TreeFormatorBinTreeImpl{
		"a",
		&TreeFormatorBinTreeImpl{"b", nil, nil, root},
		&TreeFormatorBinTreeImpl{"c", nil, nil, root},
		nil,
	}
	root.Left.SetLeft("d")
	root.Left.SetRight("e")
	root.Right.SetLeft("f")
	root.Right.SetRight("g")
	root.Right.Left.SetRight("h")
	root.Right.Left.Right.SetRight("i")
	return
}
