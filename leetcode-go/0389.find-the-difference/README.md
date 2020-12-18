## 0389. 找不同
给定两个字符串 s 和 t，它们只包含小写字母。

字符串t由字符串s随机重排，然后在随机位置添加一个字母。

请找出在 t 中被添加的字母。

```
输入：s = "abcd", t = "abcde"
输出："e"
解释：'e' 是那个被添加的字母。
```
```
输入：s = "", t = "y"
输出："y"
```
```
输入：s = "a", t = "aa"
输出："a"
```
```
输入：s = "ae", t = "aea"
输出："a"
```
提示：

- 0 <= s.length <= 1000
- t.length == s.length + 1
- s 和 t 只包含小写字母

来源：力扣（LeetCode）     
链接：https://leetcode-cn.com/problems/find-the-difference
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
#### 效率
- 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
- 内存消耗：2.2 MB, 在所有 Go 提交中击败了28.37%的用户
#### 关键词
- 位运算，哈希表
#### 思路
参考[官方题解](https://leetcode-cn.com/problems/find-the-difference/solution/zhao-bu-tong-by-leetcode-solution-mtqf/)     
- 方法一：因为题目描述不需要字母与顺序有关，所以用哈希表存储字母出现次数
- 方法二：位运算，出现奇数次的字符(TODO
)
- 方法三：求和，两个字符串的ASCII和相减