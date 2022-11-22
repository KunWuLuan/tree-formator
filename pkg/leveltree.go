package treeformator

import (
	"fmt"

	"github.com/golang/glog"
)

type TreeFormatorLevelTreeImpl struct {
	Value  interface{}
	Child  []*TreeFormatorLevelTreeImpl
	Parent *TreeFormatorLevelTreeImpl
	root   *TreeFormatorLevelTreeImpl
}

func (t *TreeFormatorLevelTreeImpl) NextLevel() []TreeFormator {
	res := []TreeFormator{}
	for _, node := range t.Child {
		res = append(res, TreeFormator(node))
	}
	return res
}

func (t *TreeFormatorLevelTreeImpl) Root() TreeFormator {
	return t.root
}

func (t *TreeFormatorLevelTreeImpl) Show() string {
	return fmt.Sprintf("%+v", t.Value)
}

func (t *TreeFormatorLevelTreeImpl) AddChild(obj interface{}) *TreeFormatorLevelTreeImpl {
	newNode := &TreeFormatorLevelTreeImpl{Value: obj, Child: []*TreeFormatorLevelTreeImpl{}, Parent: t, root: t.root}
	t.Child = append(t.Child, newNode)
	return newNode
}

type StringLevelTree struct {
	TreeFormatorLevelTreeImpl
}

func (t *StringLevelTree) Show() string {
	if t.Value == nil {
		return ""
	}
	str, ok := t.Value.(string)
	if !ok {
		glog.Errorf("node value is not string type:%v", t.Value)
	}
	return fmt.Sprintf("%v", str)
}

func BuildTest() *StringLevelTree {
	t := &StringLevelTree{}
	tt := t.AddChild("abc")
	ttt := tt.AddChild("e")
	tt.AddChild("f")
	ttt.AddChild("g")
	return t
}
