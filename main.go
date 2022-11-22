package main

import treeformator "github.com/KunWuLuan/tree-formator/pkg"

func main() {
	tree := treeformator.BuildTest()
	treeformator.ShowTreeStruct(tree)
}
