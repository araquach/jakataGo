let mix = require('laravel-mix');

/*
 |--------------------------------------------------------------------------
 | Mix Asset Management
 |--------------------------------------------------------------------------
 |
 | Mix provides a clean, fluent API for defining some Webpack build steps
 | for your Laravel application. By default, we are compiling the Sass
 | file for your application, as well as bundling up your JS files.
 |
 */

// mix.js('src/app.js', 'dist/').sass('src/app.scss', 'dist/');

mix.sass('assets/sass/app.scss', 'public/css', {
    implementation: require('node-sass')
});

mix.scripts([
    'assets/js/app.js'
], 'public/js/app.js');

// mix.js('assets/js/app.js', 'public/js/app.js');
