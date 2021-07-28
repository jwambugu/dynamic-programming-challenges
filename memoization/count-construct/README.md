# Count Construct

Write a function `countConstruct(target, wordBank)` that accepts a target string, and an array of strings. The function
should return the number of ways that the `target` can be constructed by concatenating elements of the `wordBank` array.

You may reuse the elements of the `wordBank` as many times as needed.

### Examples

 ```javascript
countConstruct('abcdef', ['ab', 'abc', 'cd', 'def', 'abcd']) // 1
// 'abc' + 'def'
```

 ```javascript
countConstruct('skateboard', ['bo', 'rd', 'ate', 'ska', 'sk', 'boar']) // 0
```

 ```javascript
countConstruct('', ['cat', 'dog', 'mouse']) // 1
```

 ```javascript
countConstruct('purple', ['purp', 'p', 'ur', 'le', 'purpl']) // 2
// 'purp' + 'le'
// 'p' + 'ur' + 'p' + 'le'
```
