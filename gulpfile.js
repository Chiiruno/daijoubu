// Compiles Less to CSS
'use strict'

const gulp = require('gulp'),
	less = require('gulp-less'),
    cssmin = require('gulp-clean-css')

gulp.task("default", () =>
    gulp.src(['less/*.less', '!less/*.mix.less'])
        .pipe(less())
        .pipe(cssmin())
        .pipe(gulp.dest('www/css'))
)
