## 290. 单词规律
给定一种规律 pattern和一个字符串str，判断 str 是否遵循相同的规律。

这里的遵循指完全匹配，例如，pattern里的每个字母和字符串str中的每个非空单词之间存在着双向连接的对应规律。

示例1:
```
输入: pattern = "abba", str = "dog cat cat dog"
输出: true
```
```
输入:pattern = "abba", str = "dog cat cat fish"
输出: false
```
```
输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false
```
```
输入: pattern = "abba", str = "dog dog dog dog"
输出: false
```

说明:
你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。 
        
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-pattern
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

#### 效率
- 执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
- 内存消耗： 1.9 MB, 在所有 Go 提交中击败了 64.96% 的用户
#### 关键词
- 哈希表
#### 思路
- 两个哈希表     
    1.用来记录某个s中的单词对应pattern中的某个字母    
    2.记录pattern中的字母是否已经被使用
- 如果s中的单词没有被和S中的字母对应，并且S中的字母没有被使用，则对应，并标记S中的字谜已使用
- 否则，比较当前单词的对应的字母和S串中同位置的字母byte值是否相同。
#### 注意
- map[string]byte{}的key未被赋值的value,默认初始值是byte(0)
#### 学习
- strings.Split(s, s2)：根据字符串s2把字符串s切割成数组