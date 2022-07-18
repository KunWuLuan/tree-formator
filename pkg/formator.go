package treeformator

const (
	space = "  "
	last  = "|-"
	line  = "| "
)

type TreeFormatorNodeInfo struct {
	node   *TreeFormator
	root   *TreeFormatorNodeInfo
	isLast bool
}

func (r TreeFormatorNodeInfo) NextLevel() []*TreeFormator {
	if r.node == nil {
		return nil
	}
	return (*r.node).NextLevel()
}

func NewTreeFormatorNodeInfo(r *TreeFormator, root *TreeFormatorNodeInfo, isLast bool) *TreeFormatorNodeInfo {
	return &TreeFormatorNodeInfo{r, root, isLast}
}

func ShowTreeStruct(r TreeFormator) {
	strs := []string{}
	dfs(NewTreeFormatorNodeInfo(&r, nil, true), strs)
}

func (r *TreeFormatorNodeInfo) Show() string {
	return (*r.node).Show()
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

type treeFormatorBinTreeImpl struct {
	value string
	left  *treeFormatorBinTreeImpl
	right *treeFormatorBinTreeImpl
	root  *treeFormatorBinTreeImpl
}

func (t treeFormatorBinTreeImpl) Show() string {
	return t.value
}

func (t treeFormatorBinTreeImpl) NextLevel() []*treeFormatorBinTreeImpl {
	res := []*treeFormatorBinTreeImpl{t.right}
	cur := t.right
	for cur != nil {
		res = append(res, cur)
		cur = cur.right
	}
	return res
}

func (t treeFormatorBinTreeImpl) Root() *treeFormatorBinTreeImpl {
	return t.root
}
