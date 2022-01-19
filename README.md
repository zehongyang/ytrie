# ytrie
####trie;可组合词前缀树
###创建一个简单前缀树并使用
```
trie := NewSimpleTrie()
trie.Add("中国")
trie.Add("北京")
ok, words := trie.Find("中国你好，北京你好")
```
###创建一个组合词前缀树并使用
```
trie := NewCombineTrie()
trie.Add("中国,大陆")
trie.Add("北京")
ok, words := trie.Find("中国你好,大陆统一")
```
