package tree

type Node struct {
	nodeName  string
	nodeIsDir bool
}

func NewNode(name string, isDir bool) Node {
	return Node{
		nodeName:  name,
		nodeIsDir: isDir,
	}
}

func (e *Node) Name() string {
	return e.nodeName
}

func (e *Node) IsDir() bool {
	return e.nodeIsDir
}
