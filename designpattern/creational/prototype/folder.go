package prototype

import "fmt"

type Folder struct {
	children []INode
	name     string
}

func (f *Folder) print(id string) {
	fmt.Println(id + f.name)
	for _, child := range f.children {
		child.print(id + id)
	}
}

func (f *Folder) clone() INode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []INode
	for _, i := range f.children {
		copyNode := i.clone()
		tempChildren = append(tempChildren, copyNode)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
