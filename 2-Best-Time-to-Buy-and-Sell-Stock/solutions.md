https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

#### Solution 1: Brute Force
Use 2 for loops to iterate through the array. 
Calculate profit at each step and record the max profit

```
time - O(N*2)
space - O(1)
```

#### Solution 2: One Pass
If current num is greater than prev num, calculate the profit by subtracting 
from previous min and set max earning

```
time - O(N)
space - O(1)
```


