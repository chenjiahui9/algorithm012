//定义前缀树
type TrieNode struct {
    word string
    children [26]*TrieNode
}
func findWords(board [][]byte, words []string) []string {
    //申明前缀树根
    root := &TrieNode{}
    //构造前缀树
    //遍历words 单词列表
    for _,word := range words {
        //将root赋给node
        node := root
        //遍历单个单词字母
        for _,w := range word {
            if node.children[w-'a'] == nil {
                node.children[w-'a'] = &TrieNode{}
            }
            node = node.children[w-'a']
        }
        node.word = word
    }
    
    result := make([]string,0)
    for i:=0; i<len(board); i++{
        for j:=0; j<len(board[0]); j++{
            dfs(i,j,board,root,&result)
        }
    }
    return result
}

//深度优先遍历
func dfs(i int,j int,broad [][]byte,node *TrieNode,result *[]string){
    //越界处理
    if i < 0 || j < 0 || i==len(broad) || j==len(broad[0]) {
        return
    }
    //获取值
    c := broad[i][j]
    if c == '#' || node.children[c-'a'] == nil {//访问过或为空
        return
    }

    node = node.children[c-'a']
    
    if node.word != "" {
        *result = append(*result,node.word)
        node.word = ""
    }

    broad[i][j] = '#'
    dfs(i+1,j,broad,node,result)
    dfs(i-1,j,broad,node,result)
    dfs(i,j+1,broad,node,result)
    dfs(i,j-1,broad,node,result)
    broad[i][j] = c
}