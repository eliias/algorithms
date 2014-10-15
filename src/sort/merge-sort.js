'use strict';

/**
 * Merge
 *
 * p <= q < r
 *
 * @see Cormen et al., Introduction to Algorithms 2009, 31.
 * @param {Array} a Array to merge
 * @param {Number} p Index - start
 * @param {Number} q Index - mid
 * @param {Number} r Index - end
 */
function merge( a, p, q, r ) {
    var i,
        j,
        k,
        n1 = q - p,
        n2 = r - q,
        lft = [], // left subarray
        rgt = []; // right subarray

    // Fill left array
    for (i = 0; i < n1; i += 1) {
        lft[i] = a[p + i];
    }

    // Fill right array
    for (j = 0; j < n2; j += 1) {
        rgt[j] = a[q + j];
    }

    // Sentinel values
    lft[n1] = Number.POSITIVE_INFINITY;
    rgt[n2] = Number.POSITIVE_INFINITY;
    i = 0;
    j = 0;

    for (k = p; k < r; k += 1) {
        if (lft[i] <= rgt[j]) {
            a[k] = lft[i];
            i += 1;
        } else {
            a[k] = rgt[j];
            j += 1;
        }
    }
}

/**
 * Merge sort
 *
 * @see Cormen et al., Introduction to Algorithms 2009, 34.
 * @param {Array} a
 * @param {Number} p
 * @param {Number} r
 * @returns {Array} Sorted array
 */
function mergeSort( a, p, r ) {
    var q;

    if (p < r - 1) {
        q = Math.floor( (p + r) / 2 );
        mergeSort( a, p, q );
        mergeSort( a, q + 1, r );
        merge( a, p, q, r );
    }

    return a;
}

// Export
module.exports = mergeSort;
