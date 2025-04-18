package tire

type Node struct {
	Pattern  string // 不为空表示匹配
	Part     string
	Children []*Node
}

func (n *Node) mathChild(part string) *Node {
	for _, child := range n.Children {
		if child.Part == part {
			return child
		}
	}
	return nil
}

func (n *Node) Insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.Pattern = pattern
		return
	}
	part := parts[height]
	child := n.mathChild(part)
	if child == nil {
		child = &Node{Part: part}
		n.Children = append(n.Children, child)
	}
	child.Insert(pattern, parts, height+1)
}

func (n *Node) matchChildren(part string) []*Node {
	var list []*Node
	for _, child := range n.Children {
		if child.Part == part || child.Part[0] == ':' || child.Part[0] == '*' {
			list = append(list, child)
		}
	}
	return list
}

func (n *Node) Search(parts []string, height int) *Node {
	if (n.Pattern != "" && len(parts) == height) || len(n.Part) > 0 && n.Part[0] == '*' {
		return n
	}
	part := parts[height]
	list := n.matchChildren(part)

	for _, child := range list {
		res := child.Search(parts, height+1)
		if res != nil {
			return res
		}
	}
	return nil
}
