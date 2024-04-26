package path

// This struct type is used for storing path from the first vertice (root)
// to the current vertice.
type Object struct {
	// Struktur object linked-list untuk path
	before *Object

	// Stored value in this Object
	Value string
}

// Return the previous element of the list
func (o *Object) Before() *Object {
	if o == nil {
		return nil
	}
	if p := o.before; p != nil {
		return p
	}
	return nil
}

// The list of vertice (path)
type Path struct {
	root *Object
	len  int
}

// Initialize the path
func (path *Path) Init() *Path {
	path.root = nil
	path.len = 0
	return path
}

// Create new path and return it's pointer
func New() *Path { return new(Path).Init() }

func (path *Path) IsEmpty() bool {
	return path.root == nil
}

// Add new vertice to the list of path
func (path *Path) Add(newNode string) *Path {
	e := Object{path.root, newNode}
	newPath := Path{&e, path.len + 1}
	return &newPath
}

// Clear all vertice to avoid memory leaks
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

// Return the path list as array of string
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

// Getter for the length of the path list
func (path *Path) Len() int {
	return path.len
}
