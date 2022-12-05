package prototype

import "fmt"

type File struct {
	name string
}

func (f *File) print(id string) {
	fmt.Println(id + f.name)
}

func (f *File) clone() INode {
	return &File{name: f.name + "_clone"}
}
