## 147. 对链表进行插入排序

## Level: mid

## cost:4h

## value：值得复习

对链表进行插入排序。  
![pic](https://upload.wikimedia.org/wikipedia/commons/0/0f/Insertion-sort-example-300px.gif)    
插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。


插入排序算法：

插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。

```
示例 1：
输入: 4->2->1->3
输出: 1->2->3->4
```
```
示例 2：

输入: -1->5->3->4->0
输出: -1->0->3->4->5
```

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/insertion-sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。



##solution

参考[官方解法](https://leetcode-cn.com/problems/insertion-sort-list/solution/dui-lian-biao-jin-xing-cha-ru-pai-xu-by-leetcode-s/)

#### Tip:[虚拟头节点](https://mp.weixin.qq.com/s/slM1CH5Ew9XzK93YOQYSjA)
- 目的是让头节点和其他节点被一视同仁。

### TYPE :排序，链表

### 时间复杂度： 平方，空间复杂度： 常数

### Key：指针操作链表，双指针