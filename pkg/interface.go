package treeformator

type Printer interface {
	Show() string
}

type TreeFormator interface {
	Printer
	NextLevel() []*TreeFormator
	Root() *TreeFormator
}
