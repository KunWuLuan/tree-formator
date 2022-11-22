package main

import treeformator "github.com/KunWuLuan/tree-formator/pkg"

func main() {
	tree := treeformator.BuildTest()
	treeformator.ShowTreeStruct(tree,
		treeformator.WithLinewidth(1),
		treeformator.WithLine(" │ "),
		treeformator.WithNotLast(" ├─"),
		treeformator.WithLast(" └─"),
		treeformator.WithSpace("   "),
	)
}
