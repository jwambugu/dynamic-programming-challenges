# How Sum

Write a function `bestSum(targetSum, numbers)` that takes in a target sum and an array of numbers as arguments.

The function should return an array containing `shortest` combinations of elements that add up to exactly target sum. If
there is no combination, then return null.

If there are multiple combinations possible, you may return any of the shortest.

### Examples

 ```javascript
bestSum(7, [5, 3, 4, 7])
// [3,4]
// [7] -> shortest
```

 ```javascript
bestSum(8, [2, 3, 5])
// [2,2,2,2]
// [2,3,3]
// [3,5] -> shortest
```

 ```javascript
bestSum(7, [2, 4])
// []
```

 ```javascript
howSum(0, [1, 2, 3])
// []
```
