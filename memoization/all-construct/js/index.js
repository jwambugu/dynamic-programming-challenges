const allConstruct = (target, wordBank, memo = {}) => {
    if (target in memo) {
        return memo[target];
    }

    if (target === "") {
        return [[]];
    }

    const result = [];

    for (let word of wordBank) {
        if (target.indexOf(word) === 0) {
            const suffix = target.slice(word.length);

            const possibleSuffixWays = allConstruct(suffix, wordBank, memo);

            const targetWays = possibleSuffixWays.map((way) => [word, ...way]);

            result.push(...targetWays);
        }
    }

    memo[target] = result;
    return memo[target];
};

console.log(allConstruct("purple", ["purp", "p", "ur", "le", "purpl"]));
// [ [ 'purp', 'le' ], [ 'p', 'ur', 'p', 'le' ] ]

console.log(
    allConstruct("abcdef", ["ab", "abc", "cd", "def", "abcd", "ef", "c"])
);

// [
//     [ 'ab', 'cd', 'ef' ],
//     [ 'ab', 'c', 'def' ],
//     [ 'abc', 'def' ],
//     [ 'abcd', 'ef' ]
// ]

console.log(
    allConstruct("skateboard", ["bo", "rd", "ate", "ska", "sk", "boar"])
);
// []

console.log(
    allConstruct("aaaaaaaaaaaaaaaaaaaaaaaaaaaaz", [
        "a",
        "aa",
        "aaa",
        "aaaa",
        "aaaaa",
    ])
);
// []
