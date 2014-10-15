'use strict';

/**
 * Insertion sort
 *
 * @param {Array} a
 * @returns {Array} sorted array
 */
function insertionSort( a ) {
    var j,
        key,
        i,
        n = a.length;

    for (j = 1; j < n; j += 1) {
        key = a[j];
        i = j - 1;
        while (i > -1 && a[i] > key) {
            a[i + 1] = a[i];
            i -= 1;
        }
        a[i + 1] = key;
    }

    return a;
}

// Export algorithm
module.exports = insertionSort;
