var largeGroupPositions = function (s) {
    let arr = []
    let count = 1;
    let start = 0, end = 0;
    for (let i = 0; i + 1 < s.length; i++) {
        if (s[i + 1] === s[i]) {
            if (count === 1) {
                start = i;
            }
            count++;
        }
        if (s[i + 1] !== s[i]&&count >= 3) {
            end = i;
            conut = 1;
            arr.push([start, end])
        }
        if (s[i + 1] !== s[i]) {
            count = 1;
        }
    }
    return arr
};
console.log(largeGroupPositions("aaa"))
//https://www.bilibili.com/video/av754360392
