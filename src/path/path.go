package path

type Object struct {
	// Struktur object linked-list untuk path
	before *Object

	// Stored value in this Object
	Value string
}

func (o *Object) Before() *Object {
	if o == nil {
		return nil
	}
	if p := o.before; p != nil {
		return p
	}
	return nil
}

type Path struct {
	root *Object
	len  int
}

func (path *Path) Init() *Path {
	path.root = nil
	path.len = 0
	return path
}

func New() *Path { return new(Path).Init() }

func (path *Path) IsEmpty() bool {
	return path.root == nil
}

func (path *Path) Add(newNode string) *Path {
	e := Object{path.root, newNode}
	newPath := Path{&e, path.len + 1}
	return &newPath
}

func (path *Path) Clear() {
	curr := path.root
	for curr != nil {
		before := curr.before
		curr.before = nil
		curr = before
	}
	path.root = nil
	path.len = 0
}

func (path *Path) GetNodes() []string {
	nodes := []string{}
	curr := path.root
	for curr != nil {
		nodes = append(nodes, curr.Value)
		before := curr.before
		curr = before
	}
	return nodes
}

func (path *Path) Len() int {
	return path.len
}
