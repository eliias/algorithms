'use strict';

var should = require( 'should' ),
    fun = require( '../src/fun' );

describe( 'fun', function() {

    describe( 'round-robin-tournament', function() {
        it( 'creates buckets with pairs', function() {
            const users = [0, 1, 2, 3, 4, 5];
            const sessions = fun.roundRobinTournament(users);

            sessions.should.be.an.Array();
            sessions.should.have.length(users.length - 1);
            sessions.should.deepEqual([
                [
                    [0,5],[1,4],[2,3]
                ],
                [
                    [0,4],[5,3],[1,2]
                ],
                [
                    [0,3],[4,2],[5,1]
                ],
                [
                    [0,2],[3,1],[4,5]
                ],
                [
                    [0,1],[2,5],[3,4]
                ]
            ]);
        } );
    } );

} );
