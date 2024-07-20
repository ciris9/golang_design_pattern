package main

import "fmt"

type Component interface {
	search(string)
}

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) search(key string) {
	for _, composite := range f.components {
		composite.search(key)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}

type File struct {
	name string
}

func (f *File) search(key string) {
	fmt.Printf("Searching for keyword %s in file %s\n", key, f.name)
}

func (f *File) getName() string {
	return f.name
}

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}
