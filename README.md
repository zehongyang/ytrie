# ytrie
trie,组合词前缀树,检测敏感词、违禁词
```
go get github.com/zehongyang/ytrie
```
### 创建一个简单前缀树并使用
```go
trie := ytrie.NewSimpleTrie()
trie.Add("中国")
trie.Add("北京")
//ok为true,并返回触发的词
ok, words := trie.Find("中国你好，北京你好")
```
### 创建一个组合词前缀树并使用
```go
trie := ytrie.NewCombineTrie()
trie.Add("中国,大陆")
trie.Add("北京")
//ok为true,并返回触发的词
ok, words := trie.Find("中国你好,大陆统一")
```
