/**
 * 
 * @param {number[]} piles 
 * @param {number} H 
 * @return {number}
 */
var minEatingSpeed = function(piles,H){
    let lo = 1,
        // 数组解构  用 ... 将数组中的每个字符解构出来
        hi = Math.max(...piles);
        console.log(lo,hi);
    // 求出最大香蕉数，在某堆中
    return lo;
}

minEatingSpeed([3, 6, 7, 11],8);