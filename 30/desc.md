# desc

Substring with Concatenation of All Words

You are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of 
substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.

For example, given:
s: "barfoothefoobarman"
words: ["foo", "bar"]

You should return the indices: [0,9].
(order does not matter).

# tag

* Hash table
* Two Pointers
* String

# 总结

解题思路：

用一个 hashmap 来存放来自 words 的 所有单词，统计每个单词出现的次数（如果没有重复的话可以用 set）。 
对于 字符串 s，从第一次字母开始迭代， 每次考虑 len(words)*len(words[0]) 的子字符串ss。
对于每一子串，依次读取一个为 words 单个单词的子字符串 str， 从 保存 words 的  hash table中查看是否存在

* 如果 str 不存在 hashtable 中， 则说明 ss 不可能是一个包含 words 中所有词汇的子串，
* 如果存在且数量大于1， 则减去一个1
* 如果存在且数量大于1, 假设这是最后一次出现该 子串 str，删除该键值
 
处理完成之后如果所有 map 所有键值被删除， 则索 ss 是一个包含所有 words 的子串
