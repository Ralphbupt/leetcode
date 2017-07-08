# First Missing Positive

Given an unsorted integer array, find the first missing positive integer.

For example,
Given [1,2,0] return 3,
and [3,4,-1,1] return 2.

Your algorithm should run in O(n) time and uses constant space.

# 思路

寻找第一个缺失的正整数，要满足 O（n） 的时间复杂度和常数级别的空间。
O（n） 的时间复杂度要求只能对数组进行有限次数的遍历，O（C）的空间复杂度
要求尽可能使用数组已有的空间。  

那么考虑将从 1 开始的所有正整数放到数组的第 i 位上， 如正整数 5 放到第 第五位，
下标为4 

# 总结

* 这道题目还是花了很长时间才做出来，而且借鉴了别人的思路。 一开始没思路的原因还在于没有认真的去思考和分析各种限定条件

在后来了解了思路之后，细节也掌握的不够到位。


