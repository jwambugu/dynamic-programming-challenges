# Memoization

In computing, memoization is an optimization technique used primarily to speed up computer programs by storing the
results of expensive function calls and returning the cached result when the same inputs occur again
([Wikipedia](https://en.wikipedia.org/wiki/Memoization))

## Memoization Recipe

### 1. Make it work

    - Visualize the problem as a tree.
    - Implement the tree using recursion.
    - Test the solution.

### 2. Make it efficient

    - Add the memo object that is shared among all the recursive calls.
    - Add a new base case to return the memo values.
    - Store the return values in the memo object.
