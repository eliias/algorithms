'use strict';

/**
 * Creates a queue
 * @see Cormen et al., Introduction to Algorithms 2009, 234.
 * @constructor
 */
function Queue() {
    this.queue = [];
    this.tail = 0;
    this.head = 0;
}

/**
 * Enqueues x
 * @param {mixed} x
 */
Queue.prototype.enqueue = function( x ) {
    this.queue[this.tail] = x;
    if (this.tail === this.queue.length) {
        this.tail = 0;
    } else {
        this.tail += 1;
    }
};

/**
 * Dequeues an element
 * @returns {mixed}
 */
Queue.prototype.dequeue = function() {
    var x = this.queue[this.head];

    if (this.head === this.queue.length) {
        this.head = 0;
    } else {
        this.head += 1;
    }

    return x;
};

// Export
module.exports = Queue;
