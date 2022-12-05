package composite

import "fmt"

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, f.name)

	for _, component := range f.components {
		component.search(keyword)
	}
}

func (f *Folder) add(component Component) {
	f.components = append(f.components, component)
}
