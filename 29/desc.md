# desc
Divide two integers without using multiplication, division and mod operator.

If it is overflow, return MAX_INT.

Subscribe to see which companies asked this question.i

# Tag:

* Math
* Binary Search

# 总结

理解思路: 拿 15 ／ 3 来举例。 15／3 = 5， 而 5 换成二进制是101，

即 15  = (1 * 2^2 + 0 * 2^1 + 1*2^0) * 3, 先找到最高的比特位 再依次下降下降，判断每一个比特对不对

需要注意的是对于各种条件的判断， 对于整数要区分正负。  可以用本题的方法去处理，确定结果的符号后即将数全部换成正整数就行。
