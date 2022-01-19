package ytrie

import (
	"bufio"
	"io"
	"os"
	"testing"
)

func TestSimpleTrie(t *testing.T)  {
	trie := NewSimpleTrie()
	fn, err := os.Open("./simple_words.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fn.Close()
	reader := bufio.NewReader(fn)
	for  {
		ld, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		trie.Add(string(ld))
	}
	ok, data := trie.Find("就覅我二姐覅哦我就佛前大陆微积分安抚鸡尾酒佛i去；卡魄可分为")
	if ok {
		t.Log("find node:",data)
	}
}


func TestCombineTrie(t *testing.T)  {
	trie := NewCombineTrie()
	fn, err := os.Open("./combine_words.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fn.Close()
	reader := bufio.NewReader(fn)
	for  {
		ld, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		trie.Add(string(ld))
	}
	ok, data := trie.Find("就覅我二姐覅哦我小日就佛前大陆微积分安抚鸡尾酒佛i去；卡魄可分为")
	if ok {
		t.Log("find node:",data)
	}
}
