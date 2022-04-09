<h2><a href="https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/">1356. Sort Integers by The Number of 1 Bits</a></h2><h3>Easy</h3><hr><div><p><font papago-translate="cached" papago-id="16">You are given an integer array </font><code>arr</code><font papago-translate="cached" papago-id="17">. Sort the integers in the array&nbsp;in ascending order by the number of </font><code>1</code><font papago-translate="cached" papago-id="18">'s&nbsp;in their binary representation and in case of two or more integers have the same number of </font><code>1</code><font papago-translate="cached" papago-id="19">'s you have to sort them in ascending order.</font></p>

<p papago-id="20" papago-translate="cached">Return <em papago-id="20-1">the array after sorting it</em>.</p>

<p>&nbsp;</p>
<p><strong papago-id="21" papago-translate="translated">Example 1:</strong></p>

<pre papago-id="22" papago-translate="cached"><strong papago-id="22-0">Input:</strong> arr = [0,1,2,3,4,5,6,7,8]
<strong papago-id="22-2">Output:</strong> [0,1,2,4,8,3,5,6,7]
<strong papago-id="22-4">Explantion:</strong> [0] is the only integer with 0 bits.
[1,2,4,8] all have 1 bit.
[3,5,6] have 2 bits.
[7] has 3 bits.
The sorted array by bits is [0,1,2,4,8,3,5,6,7]
</pre>

<p><strong papago-id="23" papago-translate="translated">Example 2:</strong></p>

<pre papago-id="24" papago-translate="cached"><strong papago-id="24-0">Input:</strong> arr = [1024,512,256,128,64,32,16,8,4,2,1]
<strong papago-id="24-2">Output:</strong> [1,2,4,8,16,32,64,128,256,512,1024]
<strong papago-id="24-4">Explantion:</strong> All integers have 1 bit in the binary representation, you should just sort them in ascending order.
</pre>

<p>&nbsp;</p>
<p><strong papago-id="25" papago-translate="translated">Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= arr.length &lt;= 500</code></li>
	<li><code>0 &lt;= arr[i] &lt;= 10<sup>4</sup></code></li>
</ul>
</div>