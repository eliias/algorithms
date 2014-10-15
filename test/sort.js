'use strict';

var sort = require( '../src/sort' );

describe( 'sorting algorithms', function() {

    describe( 'insertion sort', function() {
        it( 'should sort A[1..n]', function() {
            var a = [5, 2, 4, 6, 1, 3];
            sort.insertionSort( a ).should.eql( [1, 2, 3, 4, 5, 6] );
        } );
    } );

    describe( 'selection sort', function() {
        it( 'should sort A[1..n]', function() {
            var a = [5, 2, 4, 6, 1, 3];
            sort.selectionSort( a ).should.eql( [1, 2, 3, 4, 5, 6] );
        } );
    } );

    describe( 'merge sort', function() {
        it( 'should sort A[1..n]', function() {
            var a = [5, 2, 4, 7, 1, 3, 2, 6];
            sort.mergeSort( a, 0, a.length ).should.eql( [1, 2, 2, 3, 4, 5, 6, 7] );
        } );
    } );

} );
