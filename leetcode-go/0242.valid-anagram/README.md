## 242. 有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

#### Level: simple

#### cost:0.5h

#### value：扩展值得复习


```
示例1:

输入: s = "anagram", t = "nagaram"
输出: true
```
```
示例 2:

输入: s = "rat", t = "car"
输出: false
```



说明:
**你可以假设字符串只包含小写字母。**

进阶:
**如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？**

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-anagram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


## Solution

#### Tip:[虚拟头节点](https://mp.weixin.qq.com/s/slM1CH5Ew9XzK93YOQYSjA)
- 目的是让头节点和其他节点被一视同仁。

### TYPE :排序，哈希表

### 时间复杂度： N，空间复杂度： 常数

### TODO: 如果输入字符串包含 unicode 字符怎么办？---------是啥意思。。