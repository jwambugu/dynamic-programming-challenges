# All Construct

Write a function `allConstruct(target, wordBank)` that accepts a target string, and an array of strings. The function
should return a 2D array containing all the ways that the `target` can be constructed by concatenating elements of
the `wordBank` array. Each element of the 2D array should represent one combination that constructs the `target`

You may reuse the elements of the `wordBank` as many times as needed.

### Examples

```javascript
allConstruct('purple', ['purp', 'p', 'ur', 'le', 'purpl'])
// [['purp','le'], ['p','ur','p','le']]
```

 ```javascript
allConstruct('abcdef', ['ab', 'abc', 'cd', 'def', 'abcd', 'ef', 'c'])
// [['ab','cd','ef'], ['ab','c','def'], ['abc', 'def'], ['abcd', 'ef']]
```

 ```javascript
allConstruct('skateboard', ['bo', 'rd', 'ate', 'ska', 'sk', 'boar'])
// []
```

 ```javascript
allConstruct('', ['cat', 'dog', 'mouse'])
// [[]]
```
