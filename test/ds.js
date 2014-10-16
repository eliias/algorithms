'use strict';

var should = require( 'should' ),
    ds = require( '../src/ds' );

describe( 'data structures', function() {

    describe( 'linked list', function() {
        it( 'search', function() {
            var n = new ds.LinkedList.Node( 4 ),
                l = new ds.LinkedList();
            l.insert( n );
            l.search( 4 ).should.be.exactly( n );
        } );

        it( 'insert', function() {
            var n = new ds.LinkedList.Node( 4 ),
                l = new ds.LinkedList();
            l.insert( n );
            l.search( 4 ).should.be.exactly( n );
        } );

        it( 'delete', function() {
            var n = new ds.LinkedList.Node( 4 ),
                l = new ds.LinkedList();
            l.insert( n );
            l.search( 4 ).should.be.exactly( n );
            l.delete( n );
            should( l.search( 4 ) ).be.exactly( null );
        } );
    } );

    describe( 'stack', function() {
        it( 'empty', function() {
            var s = new ds.Stack();
            s.empty().should.eql( true );
        } );

        it( 'push', function() {
            var s = new ds.Stack();
            s.push( 1 );
            s.empty().should.not.eql( true );
        } );

        it( 'pop', function() {
            var s = new ds.Stack();
            s.push( 1 );
            s.pop().should.eql( 1 );
            s.empty().should.eql( true );
        } );
    } );

} );
