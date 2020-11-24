## 0222.完全二叉树的节点个数
#### Level: mid

#### cost:1.5h

#### value：值得品味
给出一个完全二叉树，求出该树的节点个数。

说明：

完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

示例:

```
输入: 
    1
   / \
  2   3
 / \  /
4  5 6
```

输出: 6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-complete-tree-nodes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

##solution

#### Tip:   
- 利用完全二叉树性质再递归，或者暴力递归.
- [参考答案](https://leetcode-cn.com/problems/count-complete-tree-nodes/solution/tu-jie-222-wan-quan-er-cha-shu-de-jie-dian-ge-shu-/)
- N个节点的完全二叉树高度\[log2(N)\]+1。（下取整）
- 
    1.Complete Binary Tree  
    2.Prefect Binary Tree   
    3.Full Binary Tree.
- 注意：算N次方用左移右移1<<N是2的N次方,不要用^(TODO)。
### TYPE :树，完全二叉树，递归，二分法

### TODO: 时间复杂度：O(logN*logN) ，空间复杂度： -

### TODO: 官方标准答案[二分法解题](https://leetcode-cn.com/problems/count-complete-tree-nodes/solution/wan-quan-er-cha-shu-de-jie-dian-ge-shu-by-leetco-2/)
