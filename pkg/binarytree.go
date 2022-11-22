package treeformator

import "fmt"

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
