var gulp  = require('gulp'),
    rev = require("gulp-rev"),
    uglify = require("gulp-uglify"),
    concat = require('gulp-concat'),
    minifyCss = require('gulp-minify-css'),
    gulpif = require("gulp-if"),
    args = require("yargs").argv,
    clean = require('gulp-clean'),
    sourcemaps = require('gulp-sourcemaps');

var staticOutput = "dist";
var paths = {
		css: {
            base: {
                src: [
                    "bower_components/bootstrap/dist/css/bootstrap.css"
                ],
                dst: "base.css"
            }
        },
		js: {
            base: {
                src: [
                    "bower_components/bootstrap/dist/js/bootstrap.js",
                    "bower_components/jquery/dist/jquery.js"
                ],
                dst: "base.js"
            }
		}
}

gulp.task('buildcss', function() {
    return gulp.src(paths['css']['base']['src'])
    .pipe(sourcemaps.init())
    .pipe(concat(paths['css']['base']['dst']))  // concat
    .pipe(gulpif(!args.debug, minifyCss({compatibility: 'ie8'})))    // minify
    .pipe(gulpif(!args.debug, rev()))            // versionate
    .pipe(sourcemaps.write())
    .pipe(gulp.dest(staticOutput))
    .pipe(rev.manifest(staticOutput+'/rev-manifest.json', {base: process.cwd()+'/'+staticOutput, merge: true}))
    .pipe(gulp.dest(staticOutput));
});

gulp.task('buildjs', function() {
    return gulp.src(paths['js']['base']['src'])
    .pipe(sourcemaps.init())
    .pipe(concat(paths['js']['base']['dst']))   // concat
    .pipe(gulpif(!args.debug, uglify()))                             // minify
    .pipe(gulpif(!args.debug, rev()))                                // versionate
    .pipe(sourcemaps.write())
    .pipe(gulp.dest(staticOutput))
    .pipe(rev.manifest(staticOutput+'/rev-manifest.json', {base: process.cwd()+'/'+staticOutput, merge: true}))
    .pipe(gulp.dest(staticOutput));
});

gulp.task('clean', function () {
    return gulp.src(staticOutput, {read: false})
        .pipe(clean());
});

gulp.task('buildassets', ["buildcss", "buildjs"]);
