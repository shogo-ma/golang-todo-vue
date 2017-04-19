var path = require('path');
var webpack = require('webpack');

module.exports = {
    entry: './assets/js/app.js',
    output: {
        path: path.resolve(__dirname, './assets/js'),
        filename: 'bundle.js'
    },
    plugins: [
        new webpack.ProvidePlugin({
            $: 'jquery'
        })
    ]
}
