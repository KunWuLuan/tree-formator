package main

import treeformator "github.com/tree-formator/pkg"

func main() {
	tree := treeformator.BuildSampleTree()
	treeformator.ShowTreeStruct(tree)
}
