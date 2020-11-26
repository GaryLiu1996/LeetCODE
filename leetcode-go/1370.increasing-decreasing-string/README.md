## 1370. 上升下降字符串
### Level: simple

### cost: 1h

### value：值得品味

给你一个字符串s，请你根据下面的算法重新构造字符串：

1.从 s中选出 最小的字符，将它 接在结果字符串的后面。   
2.从 s剩余字符中选出最小的字符，且该字符比上一个添加的字符大，将它 接在结果字符串后面.   
3.重复步骤 2 ，直到你没法从 s中选择字符。    
4.从 s中选出 最大的字符，将它 接在结果字符串的后面。   
5.从 s剩余字符中选出最大的字符，且该字符比上一个添加的字符小，将它 接在结果字符串后面。  
6.重复步骤 5，直到你没法从 s中选择字符。     
7.重复步骤 1 到 6 ，直到 s中所有字符都已经被选过。  
          
在任何一步中，如果最小或者最大字符不止一个，你可以选择其中任意一个，并将其添加到结果字符串。 
请你返回将s中字符重新排序后的 **结果字符串** 。

```
示例 1：

输入：s = "aaaabbbbcccc"
输出："abccbaabccba"
解释：第一轮的步骤 1，2，3 后，结果字符串为 result = "abc"
第一轮的步骤 4，5，6 后，结果字符串为 result = "abccba"
第一轮结束，现在 s = "aabbcc" ，我们再次回到步骤 1
第二轮的步骤 1，2，3 后，结果字符串为 result = "abccbaabc"
第二轮的步骤 4，5，6 后，结果字符串为 result = "abccbaabccba"
```
```
示例 2：

输入：s = "rat"
输出："art"
解释：单词 "rat" 在上述算法重排序以后变成 "art"
```
```
示例 3：

输入：s = "leetcode"
输出："cdelotee"
```
```
示例 4：

输入：s = "ggggggg"
输出："ggggggg"
```
```
示例 5：

输入：s = "spo"
输出："ops"

```

提示：

1 <= s.length <= 500
s只包含小写英文字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/increasing-decreasing-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

## Solution

#### Tip
- 我们只关注字符本身，而不关注字符在原字符串中的位置。因此我们可以直接创建一个大小为 26 的桶，记录每种字符的数量。每次提取最长的上升或下降字符串时，我们直接顺序或逆序遍历这个桶。
  
  作者：LeetCode-Solution
  链接：https://leetcode-cn.com/problems/increasing-decreasing-string/solution/shang-sheng-xia-jiang-zi-fu-chuan-by-leetcode-solu/
  来源：力扣（LeetCode）
  著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

### TYPE :排序，字符串，[桶计数](https://leetcode-cn.com/problems/increasing-decreasing-string/solution/shang-sheng-xia-jiang-zi-fu-chuan-by-leetcode-solu/) ，哈希表

### 时间复杂度： O(∣Σ∣×∣s∣)，其中 \SigmaΣ 为字符集，ss 为传入的字符串，空间复杂度： O(|\Sigma|)O(∣Σ∣)

### TODO:
- 桶计数，哈希表，哈希桶的异同
- int32(rune),byte(utf-8),int的异同
- 时间复杂度的计算

- 基数排序与计数排序、桶排序这三种排序算法都利用了桶的概念，但对桶的使用方法上有明显差异：  

    - 基数排序：根据键值的每位数字来分配桶；
    - 计数排序：每个桶只存储单一键值；
    - 桶排序：每个桶存储一定范围的数值；(桶计数)
    - 基数排序不是直接根据元素整体的大小进行元素比较，而是将原始列表元素分成多个部分，对每一部分按一定的规则进行排序，进而形成最终的有序列表。