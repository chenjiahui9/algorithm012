学习笔记
(DFS)深度优先模板:
递归写法：
visited = set()

def dfs(node,visited)
	visited(node)
	
	for next_node in node.children():
		if not next_node in visited :
			dfs(next node,visited)

非递归写法：
def dfs(self,tree):
	if tree.root is None:
		return []
	
	visited,stack = [],[tree.root]

	while stack:
		node = stack.pop()
		visited.add(node)
		
		process(node)
		nodes = generate_related_nodes(node)
		stack.push(nodes)
		

广度优先：
def BFS(graph,start,end):
	queue = []  //队列
	queue.append([start]) //入队
	visited.add(start) //入集合
	
	while queue: //队列不为空
		node = queue.pop() //出队
		visited.add(node) //入计划
		
		process()
		nodes = generate_related_nodes(node) //产生node
		queue.push(nodes
		

贪心算法 Greedy
贪心算法是一种在每一步选择中都采取最好或最优（既最有利）的选择，从而希望导致结果是全局最好或最优的算法。
贪心算法和动态规划最大的不同在于它对每个子问题的解决方案都做出选择，并且不能做出回退。动态规划会保存以前的结果，
并根据以前的结果对当前做出选择，有回退功能。


二分查找的前提
1.目标函数的单调性（单调递增或单调递减）
2.存在上下界(bounded)
3.能够通过索引访问(index accessible)
代码模板:
left,right = 0,len(array) - 1
while left <= right:
	mid = (left + right)/2
	if array(mid) == target:
		#find the target
		break or return result
		elif array[mid] < target:
	elif array[mid] < target:
		left = mid + 1
	else :
		right = mid -1





























