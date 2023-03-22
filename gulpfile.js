'use strict';

const gulp = require('gulp');
const sass = require('gulp-sass')(require('sass'));

function buildStyles() {
  return gulp.src('./assets/scss/*.scss')
    .pipe(sass().on('error', sass.logError))
    .pipe(gulp.dest('./static/css'));
};

function buildProdcutionSytles() {
  return gulp.src('./assets/scss/*.scss')
      .pipe(sass({outputStyle:'compressed'}).on('error', sass.logError))
      .pipe(gulp.dest('./static/css'));
}

exports.buildStyles = buildStyles;
exports.buildProdStyles = buildProdcutionSytles;

exports.watch = function () {
  gulp.watch('./assets/scss/*.scss', gulp.series('buildStyles'));
};