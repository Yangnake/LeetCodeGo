BFS的核心思想就是把问题抽象成图， 从一个点开始，向四周开始扩散。一般来说平常写的算法都是用**队列**这种数据结构，每次将一个节点周围所有节点加入队列。

BFS相对DFS的最主要区别是：BFS找到的路径一定是最短的，但代价就是空间复杂度可能比DFS大很多。

```
// 计算从起点start到终点target的最近距离
int BFS(Node start, Node target) {
    Queue<Node> q; //核心数据结构
    Set<Node> visited; //避免走回头路
    
    q.offer(start); //将起点加入队列
    visited.add(start);
    int step = 0; //记录扩散的步数
    while(q not empty) {
        int sz = q.size();
        /*将当前队列中的所有节点向四周扩散*/
        for (int i = 0; i < sz; i++) {
            Node cur = q.poll();
            /*划重点： 这里判断是否到达终点*/
            if (cur is target)
                return step;
			/*将cur的相邻节点加入队列*/
			if (Node x : cur.adj()) {
			    if (x not in visited) {
			        q.offer(x);
			        visited.add(x);
			    }
			}
        }
        /*划重点: 更新步数在这里*/
        step++;
    }
}
```
队列 q 就不说了，BFS 的核心数据结构；cur.adj() 泛指 cur 相邻的节点，比如说二维数组中，cur 上下左右四面的位置就是相邻节点；visited 的主要作用是防止走回头路，大部分时候都是必须的，但是像一般的二叉树结构，没有子节点到父节点的指针，不会走回头路就不需要 visited。

Dijkstra算法待更新。。。

双向BFS优化

拼图问题，将问题转化为BFS，即给图中的每个点建立索引是一个难点