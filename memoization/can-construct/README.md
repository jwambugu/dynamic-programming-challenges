# Can Construct

Write a function `canConstruct(target, wordBank)` that accepts a target string, and an array of strings. The function
should return a boolean indicating whether the `target` can be constructed by concatenating elements of the `wordBank`
array.

You may reuse the elements of the `wordBank` as many times as needed.

### Examples

 ```javascript
canConstruct('abcdef', ['ab', 'abc', 'cd', 'def', 'abcd']) // true
// 'abc' + 'def'
```

 ```javascript
canConstruct('skateboard', ['bo', 'rd', 'ate', 'ska', 'sk', 'boar']) // false
```

 ```javascript
canConstruct('', ['cat', 'dog', 'mouse']) // true
```
