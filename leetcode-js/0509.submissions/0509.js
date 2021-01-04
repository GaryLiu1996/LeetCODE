/**
 * @param {number} n
 * @return {number}
 */
var fib = function (n) {
    if (n === 0 || n === 1) {
        return n
    }
    let num0 = 0, num1 = 1
    for (let i = 1; i < n; i++) {
        let t = num1
        num1 = num0 + num1
        num0 = t
    }
    return num1
};