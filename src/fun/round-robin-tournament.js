'use strict';

function *pairs(a) {
    a = [].concat(a) // copy to avoid side-effects
    const n = a.length;

    for(let z = 0; z < n - 1; z += 1) {
        const p = [];
        for(let i = 0; i < n / 2; i += 1) {
            p.push([a[i], a[n - i - 1]]);
        }
        yield p;
        // rotate clockwise with fixed left column
        a.splice(1, 0, a.pop());
    }
}

/**
 * Create max number of pairs distributed across buckets
 *
 * @param {Array} a
 * @returns {Array} sessions
 */
function roundRobinTournament( a ) {
    const n = a.length - 1;
    const s = [];
    const p = pairs(a);
    while(s.push(p.next().value) < n);
    return s;
}

// Export algorithm
module.exports = roundRobinTournament;
