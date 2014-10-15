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

} );
