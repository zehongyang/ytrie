package ytrie


type Node struct {
	child map[rune]*Node
	finished bool
	data string
}

type SimpleTrie struct {
	root *Node
}

func NewSimpleTrie() *SimpleTrie {
	return &SimpleTrie{root: &Node{child: make(map[rune]*Node)}}
}

func (t *SimpleTrie) Add (data string) {
	rd := []rune(data)
	root := t.root
	for _, r := range rd {
		node,ok := root.child[r]
		if !ok {
			node = &Node{child: make(map[rune]*Node)}
			root.child[r] = node
		}
		root = node
	}
	root.finished = true
	root.data = data
}

func (t *SimpleTrie) find(data string) (bool,*Node) {
	rd := []rune(data)
	root := t.root
	for _, r := range rd {
		node,ok := root.child[r]
		if ok {
			if node.finished {
				return true,node
			}
			root = node
		}else{
			root = t.root
		}
	}
	return false,nil
}

func (t *SimpleTrie) Find (data string) (bool,string) {
	for i, _ := range data {
		ok, node := t.find(data[i:])
		if ok && node != nil {
			return ok,node.data
		}
	}
	return false,""
}
