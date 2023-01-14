https://leetcode.com/problems/contains-duplicate/

### Solution 1
Brute force

Use 2 for loops to iterate over the numbers and find matching numbers
```
time - O(N*2)
space - O(1)
```


### Solution 2
Step 1: Sort the array ( O(NlogN) )
Step 2: Iterate through array once and check if next number is matching ( O(N))

```
time - O(NlogN)
space - O(1)
```

### Solution 3
Use a map.
Step 1: Iterate through array once and check if number exist in map.
If not, add the number to map

```
time - O(N)
space - O(N)
```