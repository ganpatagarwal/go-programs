https://leetcode.com/problems/two-sum/

#### Solution 1: Brute Force
Use 2 for loops to iterate through the array.
find if there is matching num which equals target-num

```
time - O(N*2)
space - O(1)
```

#### Solution 2: One Pass using map
iterate through array once and check if there is an entry in map for 
`target-num`

if not, add the num to map and procedd

```
time - O(N)
space - O(N)
```