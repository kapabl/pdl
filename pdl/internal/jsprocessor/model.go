package jsprocessor

type FileInfo struct {
	Name string
	Dir  string
}

type ClassNode struct {
	FullName        string
	Name            string
	Namespace       string
	Path            string
	UsageName       string
	IsClass         bool
	IsFloatingClass bool
	IsIntrinsic     bool
	IsArray         bool
	IsImported      bool
	FileInfo        FileInfo
}

type Property struct {
	Name      string
	ClassNode ClassNode
}

type ClassData struct {
	JSFullFilename  string
	ClassNode       ClassNode
	ParentClassNode ClassNode
	Properties      []Property
	Imports         []string
	Last            bool
}

func (class ClassData) IsValidClass() bool {
	return class.ClassNode.Name != "" && class.ClassNode.Namespace != ""
}

type GlobalIndexClass struct {
	Namespace string
	Classes   []ClassData
}

type NamespaceFile struct {
	Root     string
	Filename string
}
