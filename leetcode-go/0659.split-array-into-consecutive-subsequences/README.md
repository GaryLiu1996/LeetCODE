[参考的解法](https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences/solution/tan-xin-suan-fa-jian-cha-shu-zu-neng-fou-bei-fen-w/)
#### 重点
- 两个哈希表
    - 存储原数组中数字i出现的次数
    - 存储以数字i结尾的且符合题意的连续子序列个数
- 贪心算法
#### 思考
- for range遍历出的是无序的，每次顺序都不同。
#### 效率
- 执行用时：100 ms, 在所有 Go 提交中击败了65.63%的用户
- 内存消耗：7 MB, 在所有 Go 提交中击败了65.63%的用户