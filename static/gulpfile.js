var gulp = require('gulp'),
	jshint = require('gulp-jshint'),
	uglify = require('gulp-uglify'),
	uglifyCss = require('gulp-uglifycss'),
	concat = require('gulp-concat'),
	concatCss = require('gulp-concat-css'),
	notify = require('gulp-notify'),
	flatten = require('gulp-flatten'),
	clean = require('gulp-clean'),
	install = require('gulp-install'),
	runSequence = require('run-sequence');

var commonScriptPaths = [
	'assets/js/vendor/purple.js',
	'assets/js/vendor/highlight.pack.js',
	'assets/bower_components/marked/marked.min.js',
	'assets/bower_components/angular/angular.js',
	'assets/bower_components/angular-ui-router/release/angular-ui-router.min.js',
	'assets/bower_components/pace/pace.js',
	'assets/bower_components/moment/moment.js',
	'assets/bower_components/angular-moment/angular-moment.js',
	'apps/shared/**/*.js'
];

var frontScriptPaths = commonScriptPaths.concat([
	'apps/front/app.module.js',
	'apps/front/app.router.js',
	'apps/front/**/*.js',
	'assets/bower_components/angular-disqus/angular-disqus.js',
	'assets/bower_components/angularjs-socialshare/src/js/angular-socialshare.js'
]);

var adminScriptPaths = commonScriptPaths.concat([
	'assets/bower_components/ace-builds/src/ace.js',
	'assets/bower_components/ace-builds/src/mode-markdown.js',
	'assets/js/theme-purple.js',
	'assets/bower_components/angular-ui-ace/ui-ace.js',
	'assets/bower_components/ng-tags-input/ng-tags-input.js',
	'assets/bower_components/d3/d3.js',
	'assets/bower_components/n3-line-chart/build/line-chart.js',
	'assets/bower_components/topojson/topojson.js',
	'assets/bower_components/angular-animate/angular-animate.js',
	'assets/bower_components/angular-toastr/dist/angular-toastr.tpls.js',
	'apps/admin/app.module.js',
	'apps/admin/app.router.js',
	'apps/admin/**/*.js'
]);

var commonCssPaths = [
	'assets/css/purple.css',
	'assets/css/color-brewer.min.css',
	'assets/bower_components/fontawesome/css/font-awesome.css',
	'assets/bower_components/pace/themes/purple/pace-theme-minimal.css',
	'assets/css/style.css'
];

var frontCssPaths = commonCssPaths.concat([]);

var adminCssPaths = [
	'assets/bower_components/ngAnimate/css/ng-animation.css',
	'assets/bower_components/angular-toastr/dist/angular-toastr.css',
	'assets/css/tags-input.css',
	'assets/css/countries.css'
	].concat(commonCssPaths);

gulp.task('clean', function (cb) {
	return gulp.src('dist', {read: false})
		.pipe(clean());
});

gulp.task('front-styles', function() {

	return gulp.src(frontCssPaths)
		.pipe(concatCss('front.min.css', {rebaseUrls:false}))
		// .pipe(uglifyCss())
		.pipe(gulp.dest('dist/css'))
		.pipe(notify({message: 'Front style task complete.'}));
});

gulp.task('admin-styles', function() {

	return gulp.src(adminCssPaths)
		.pipe(concatCss('admin.min.css', {rebaseUrls:false}))
		// .pipe(uglifyCss())
		.pipe(gulp.dest('dist/css'))
		.pipe(notify({message: 'Admin style task complete.'}));
});

gulp.task('styles', ['front-styles', 'admin-styles']);

gulp.task('front-scripts', function () {
	
	return gulp.src(frontScriptPaths)
		.pipe(concat('front.min.js'))
		// .pipe(uglify())
		.pipe(gulp.dest('dist/js'))
		.pipe(notify({message: 'Front scripts task complete.'}));
});

gulp.task('admin-scripts', function () {
	
	return gulp.src(adminScriptPaths)
		.pipe(concat('admin.min.js'))
		// .pipe(uglify())
		.pipe(gulp.dest('dist/js'))
		.pipe(notify({message: 'Admin scripts task complete.'}));
});

gulp.task('scripts', ['front-scripts', 'admin-scripts']);

gulp.task('copyScripts', function () {

});

gulp.task('copyFonts', function () {
	var copyPaths = [
		'assets/fonts/*.*',
		'assets/bower_components/fontawesome/fonts/*.*'
	];

	return gulp.src(copyPaths)
		.pipe(gulp.dest('dist/fonts'));
});

gulp.task('copy', ['copyScripts', 'copyFonts']);

gulp.task('watch', function () {
	gulp.watch(commonScriptPaths.concat(frontScriptPaths).concat(adminScriptPaths), ['scripts']);
	gulp.watch(commonCssPaths.concat(frontCssPaths).concat(adminCssPaths), ['styles']);
});

gulp.task('bower', function () {
	console.log('Start update bower packages...');

	return gulp.src('assets/bower.json')
		.pipe(install());
});

gulp.task('default', function () {
	runSequence(
		'clean',
		'bower',
		'styles',
		'scripts',
		'copy',
		'watch',
		function () {
			console.log('Start listening file changes...');
		}
	);
});