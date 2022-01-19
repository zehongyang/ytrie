package ytrie

import "strings"

type CombineTrie struct {
	root *CombineNode
}

type CombineNode struct {
	child map[rune]*CombineNode
	finished bool
	data string
	ct *CombineTrie
}

func NewCombineTrie() *CombineTrie {
	return &CombineTrie{root: &CombineNode{child: make(map[rune]*CombineNode)}}
}

func (t *CombineTrie) Add (data string) {
	sd := strings.Split(data, ",")
	root := t.root
	for i, s := range sd {
		rd := []rune(s)
		for _, r := range rd {
			node,ok := root.child[r]
			if !ok {
				node = &CombineNode{child: make(map[rune]*CombineNode)}
				root.child[r] = node
			}
			root = node
		}
		if i == len(sd) - 1 {
			root.finished = true
			root.data = data
		}else{
			if root.ct == nil {
				root.ct = NewCombineTrie()
			}
			root = root.ct.root
		}
	}
}

func (t *CombineTrie) find (data string) (bool,*CombineNode) {
	rd := []rune(data)
	root := t.root
	for _, r := range rd {
		node,ok := root.child[r]
		if ok {
			if node.finished {
				return true,node
			}
			root = node
			if node.ct != nil {
				ok, cNode := node.ct.find(data)
				if ok {
					return ok,cNode
				}
			}
		}else{
			root = t.root
		}
	}
	return false,nil
}

func (t *CombineTrie) Find (data string) (bool,string) {
	for i, _:= range data {
		ok, node := t.find(data[i:])
		if ok && node != nil {
			return ok,node.data
		}
	}
	return false,""
}
