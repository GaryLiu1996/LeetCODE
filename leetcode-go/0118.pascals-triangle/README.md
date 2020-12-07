## 0118.杨辉三角
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。    
![gif](./PascalTriangleAnimated2.gif)   
在杨辉三角中，每个数是它左上方和右上方的数的和。
示例:
```
输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
```

### 解法
1. [定义法](https://leetcode-cn.com/problems/pascals-triangle/solution/yang-hui-san-jiao-by-leetcode-solution-lew9/)
2. [错位相加(取巧)](https://leetcode-cn.com/problems/pascals-triangle/solution/qu-qiao-jie-fa-cuo-yi-wei-zai-zhu-ge-xiang-jia-28m/)
- Type:数组
- Tip: 首尾都是1不需要加入循环
- 时间复杂度O(numRows*numRows)
- 空间复杂度O(1)