'use strict';

/**
 * Selection sort
 *
 * @param {Array} a
 * @returns {Array} sorted array
 */
function selectionSort( a ) {
    var j,
        i,
        k,
        s,
        n = a.length;

    for (j = 0; j < n - 1; j += 1) {
        // Current key
        k = a[j];

        // Find smallest
        s = j;
        for (i = j; i < n; i += 1) {
            if (a[i] < a[s]) {
                s = i;
            }
        }

        // Exchange keys
        a[j] = a[s];
        a[s] = k;
    }

    return a;
}

// Export algorithm
module.exports = selectionSort;
