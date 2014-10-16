'use strict';

/**
 * Creates a stack
 * @see Cormen et al., Introduction to Algorithms 2009, 232.
 * @constructor
 */
function Stack() {
    this.stack = [];
    this.top = -1;
}

/**
 * Whether or not the stack is empty
 * @returns {Boolean} True if stack is empty, otherwise false
 */
Stack.prototype.empty = function() {
    if (this.top === -1) {
        return true;
    }

    return false;
};

/**
 * Adds an element to the stack
 * @param {mixed} x
 */
Stack.prototype.push = function( x ) {
    this.top += 1;
    this.stack[this.top] = x;
};

/**
 * Removes an element from stack
 */
Stack.prototype.pop = function() {
    if (this.empty()) {
        throw Error( 'Underflow' );
    }

    this.top -= 1;
    return this.stack[this.top + 1];
};

// Export
module.exports = Stack;
