# How Sum

Write a function `howSum(targetSum, numbers)` that takes in a target sum and an array of numbers as arguments.

The function should return an array containing any combinations of elements that add up to exactly target sum. If there
is no combination, then return null.

If there are multiple combinations possible, you may return any single one.

### Examples

 ```javascript
howSum(7, [5, 3, 4, 7])
// [3,4]
// [7]
```

 ```javascript
howSum(8, [2, 3, 5])
// [2,2,2,2]
// [2,3,3]
// [3,5]
```

 ```javascript
howSum(7, [2, 4])
// []
```

 ```javascript
howSum(0, [1, 2, 3])
// []
```
