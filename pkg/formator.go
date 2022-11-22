package treeformator

import (
	"fmt"
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

func ShowTreeStruct(r TreeFormator, opts ...Option) {
	printer := &treePrinterImpl{
		linewidth: 1,
		space:     "  ",
		last:      "\\-",
		line:      "| ",
		notlast:   "|-",
	}
	for _, opt := range opts {
		opt(printer)
	}

	strs := printer.Dfs(NewTreeFormatorNodeInfo(r, nil, true))
	for _, str := range strs {
		fmt.Println(str)
	}
}

func (r *TreeFormatorNodeInfo) Show() string {
	return r.node.Show()
}

type Option func(TreePrinter) error

type TreePrinter interface {
	BuildPrefix(r *TreeFormatorNodeInfo) string
	Dfs(r *TreeFormatorNodeInfo) (str []string)
	SetProperty(name string, value interface{}) error
}

type treePrinterImpl struct {
	linewidth int
	space     string
	last      string
	line      string
	notlast   string
}

func (tp *treePrinterImpl) BuildPrefix(r *TreeFormatorNodeInfo) string {
	if r == nil {
		return ""
	}
	cur := r.Show()
	if r.root == nil {
		return cur
	} else {
		if r.isLast {
			cur = tp.last + cur
		} else {
			cur = tp.notlast + cur
		}
	}
	curNode := r.root
	root := curNode.root
	for root != nil {
		if curNode.isLast {
			cur = tp.space + cur
		} else {
			cur = tp.line + cur
		}
		curNode = root
		root = curNode.root
	}
	return cur
}

func (tp *treePrinterImpl) Dfs(r *TreeFormatorNodeInfo) (str []string) {
	if r == nil {
		return []string{}
	}
	nextLevelNodes := r.NextLevel()
	str = append(str, tp.BuildPrefix(r))
	for i, child := range nextLevelNodes {
		str = append(str, tp.Dfs(NewTreeFormatorNodeInfo(child, r, i == len(nextLevelNodes)-1))...)
	}
	return
}

func (tp *treePrinterImpl) SetProperty(name string, value interface{}) error {
	switch name {
	case "linewidth":
		l, ok := value.(int)
		if !ok {
			return fmt.Errorf("value of linewidth must be int")
		}
		tp.linewidth = l
	case "line":
		l, ok := value.(string)
		if !ok {
			return fmt.Errorf("value of line must be string")
		}
		tp.line = l
	case "space":
		l, ok := value.(string)
		if !ok {
			return fmt.Errorf("value of space must be string")
		}
		tp.space = l
	case "notlast":
		l, ok := value.(string)
		if !ok {
			return fmt.Errorf("value of notlast must be string")
		}
		tp.notlast = l
	case "last":
		l, ok := value.(string)
		if !ok {
			return fmt.Errorf("value of last must be string")
		}
		tp.last = l
	default:
		return fmt.Errorf("undefined property")
	}
	return nil
}

func WithLinewidth(l int) Option {
	return func(tp TreePrinter) error {
		return tp.SetProperty("linewidth", l)
	}
}

func WithSpace(s string) Option {
	return func(tp TreePrinter) error {
		return tp.SetProperty("space", s)
	}
}

func WithLine(s string) Option {
	return func(tp TreePrinter) error {
		return tp.SetProperty("line", s)
	}
}

func WithLast(s string) Option {
	return func(tp TreePrinter) error {
		return tp.SetProperty("last", s)
	}
}

func WithNoLast(s string) Option {
	return func(tp TreePrinter) error {
		return tp.SetProperty("nolast", s)
	}
}
