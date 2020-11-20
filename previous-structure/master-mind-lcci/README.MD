### 题目
珠玑妙算游戏（the game of master mind）的玩法如下。

计算机有4个槽，每个槽放一个球，颜色可能是红色（R）、黄色（Y）、绿色（G）或蓝色（B）。例如，计算机可能有RGGB 4种（槽1为红色，槽2、3为绿色，槽4为蓝色）。作为用户，你试图猜出颜色组合。打个比方，你可能会猜YRGB。要是猜对某个槽的颜色，则算一次“猜中”；要是只猜对颜色但槽位猜错了，则算一次“伪猜中”。注意，“猜中”不能算入“伪猜中”。

给定一种颜色组合solution和一个猜测guess，编写一个方法，返回猜中和伪猜中的次数answer，其中answer[0]为猜中的次数，answer[1]为伪猜中的次数。

示例：

输入： solution="RGBY",guess="GGRR"
输出： [1,1]
解释： 猜中1次，伪猜中1次。
提示：

len(solution) = len(guess) = 4
solution和guess仅包含"R","G","B","Y"这4种字符

```来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/master-mind-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

### 解题
#### 解题思路
- map[string]int用来储存solution的字符出现频率
- guess数组去掉完全匹配的字符，去掉对应位置在soulution的字符
#### 难点
- 题目理解:     
首先去掉guess和solution的猜中元素，剩余的元素，solution中计算每个字符出现频率，      
遍历guess字符串，每遍历一个guess的元素，若匹配上solution的map的key，统计的       
数字减去一，伪猜中+1。
- go基础
#### 知识点
- map和slice的操作
#### 拓展
- rune : int32别名
- 初始化map[string]int ---> value=0
- go string底层是byte数组

```
作者：garyliu1996
链接：https://leetcode-cn.com/problems/master-mind-lcci/solution/20201030-by-garyliu1996/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
```
