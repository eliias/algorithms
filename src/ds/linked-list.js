'use strict';

/**
 * Creates a linked list
 * @see Cormen et al., Introduction to Algorithms 2009, 236.
 * @constructor
 */
function LinkedList() {
    this.head = null;
}

/**
 * Searches for parameter k within the linked list
 * @param {Number} k
 * @returns {LinkedList.Node}
 */
LinkedList.prototype.search = function( k ) {
    var x = this.head;
    while (x !== null && x.key !== k) {
        x = x.next;
    }
    return x;
};

/**
 * Splices a node onto the front of the linked list
 * @param {LinkedList.Node} node
 */
LinkedList.prototype.insert = function( node ) {
    node.next = this.head;
    if (this.head !== null) {
        this.head.prev = node;
    }
    this.head = node;
    node.prev = null;
};

/**
 * Removes a node from the list
 * @param node
 */
LinkedList.prototype.delete = function( node ) {
    if (node.prev !== null) {
        node.prev.next = node.next;
    } else {
        this.head = node.next;
    }
    if (node.next !== null) {
        node.next.prev = node.prev;
    }
};

/**
 * A node of the linked list
 * @constructor
 * @param {Number} k
 */
LinkedList.Node = function( k ) {
    this.next = null;
    this.prev = null;
    this.key = k;
};

// Export
module.exports = LinkedList;
