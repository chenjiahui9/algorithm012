学习笔记

递归模板
public void recur(int level,int param) {

	if level > max {//递归终止条件 1
		return
	}
	
	process(level,param)//处理当前层逻辑 2
	
	recur(level,param)//递归到下一层 3
	
	//有需要的话恢复当前层状态	4
}

分治代码模板
def divide_conquer(problem,param1,param2):
	
	# recursion terminator(递归终止条件)
	if problem is None:
		print_result
		return
	
	# prepare data(拆分子问题)
	data = prepare_data(problem)
	subproblems = split_problem(problem,data)
	
	# conquer subproblems （调子问题的递归函数）
	subresult1 = self.divide_conquer(subproblems[0],p1,...)
	subresult2 = self.divide_conquer(subproblems[1],p2,...)
	subresult3 = self.divide_conquer(subproblems[2],p3,...)
	...
	
	#process and generate the final result （合并结果）
	result = process_result(subresult1,subresult2,subresult3,...)
	#reveret the current level states 恢复当前层状态
	


1.动态规划的实现及关键点
1） 人肉递归低效，很累
2） 找到最近最简的方法，将其拆解成可重复解决的问题
3） 数学归纳法思维（抵制人肉递归的诱惑）
	