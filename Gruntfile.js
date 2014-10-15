module.exports = function( grunt ) {
    'use strict';

    // Measure task time
    require( 'time-grunt' )( grunt );
    require( 'load-grunt-tasks' )( grunt );

    // Project configuration.
    grunt.initConfig( {
        // Read package
        pkg: grunt.file.readJSON( 'package.json' ),

        // Config
        config: {
            src: './src'
        },

        // JSHint
        jshint: {
            options: {
                jshintrc: true
            },
            test: [
                './Gruntfile.js',
                './src/**/*.js',
                './test/**/*.js'
            ]
        },

        // Tests
        mochaTest: {
            test: {
                options: {
                    require: 'should',
                    reporter: 'dot',
                    ui: 'bdd'
                },
                src: ['./test/**/*.js']
            }
        }

    } );

    // Default task(s).
    grunt.registerTask( 'test', ['mochaTest'] );
    grunt.registerTask( 'default', ['jshint', 'test'] );

};
