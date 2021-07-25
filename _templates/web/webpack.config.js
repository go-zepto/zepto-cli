'use strict';

var path = require('path');
var webpack = require('webpack');
const { WebpackManifestPlugin } = require('webpack-manifest-plugin');

var CompressionPlugin = require("compression-webpack-plugin");

var host = process.env.HOST || 'localhost'
var devServerPort = 3808;

var production = process.env.NODE_ENV === 'production';

const ExtractCssChunks = require("extract-css-chunks-webpack-plugin");

var entry = require("./assets/entry.json");

class CleanUpExtractCssChunks {
    shouldPickStatChild(child) {
        return child.name.indexOf('extract-css-chunks-webpack-plugin') !== 0;
    }
    apply(compiler) {
        compiler.hooks.done.tap('CleanUpExtractCssChunks', (stats) => {
            const children = stats.compilation.children;
            if (Array.isArray(children)) {
                // eslint-disable-next-line no-param-reassign
                stats.compilation.children = children
                    .filter(child => this.shouldPickStatChild(child));
            }
        });
    }
}

// Resolve relative path for entry.json
Object.keys(entry).forEach(key => {
    entry[key] = entry[key].map(path => './assets/' + path);
});

var config = {
    //stats: { children: false },
    mode: production ? "production" : "development",
    entry: entry,
    module: {
        rules: [
            { test: /\.es6$/, use: "babel-loader" },
            { test: /\.jsx$/, use: "babel-loader" },
            //{ test: /react-select\/src/, use: "babel-loader" },
            { test: /\.(jpe?g|png|gif)$/i, use: "file-loader" },
            {
                test: /\.woff($|\?)|\.woff2($|\?)|\.ttf($|\?)|\.eot($|\?)|\.svg($|\?)|\.otf($|\?)/,
                //use: production ? 'file-loader' : 'url-loader'
                use: 'file-loader'
            },
            {
                test: /\.(sass|scss|css)$/,
                use: [
                    {
                        loader: ExtractCssChunks.loader,
                        options: {
                            hot: production ? false : true,
                            // Force reload all
                            //reloadAll: true,
                        }
                    },
                    {
                        loader: "css-loader",
                        options: {
                            //minimize: true,
                            sourceMap: true
                        }
                    },
                    {
                        loader: "sass-loader"
                    }
                ]
            },
        ]
    },

    output: {
        path: path.join(__dirname, ".", 'public', 'build'),
        publicPath: '/build/',

        filename: production ? '[name]-[chunkhash].js' : '[name].js',
        chunkFilename: production ? '[name]-[chunkhash].js' : '[name].js',
    },

    resolve: {
        modules: [path.resolve(__dirname, "..", "build"), path.resolve(__dirname, ".", "node_modules")],
        extensions: [".es6", ".jsx", ".sass", ".css", ".js"],
        alias: {
            '~': path.resolve(__dirname, ".", "build"),
        }
    },

    plugins: [
        new ExtractCssChunks(
            {
                filename: production ? "[name]-[chunkhash].css" : "[name].css",
                chunkfilename: production ? "[name]-[id].css" : "[name].css",
            }
        ),
        new CleanUpExtractCssChunks(),
        new WebpackManifestPlugin({
            writeToFileEmit: true,
            publicPath: production ? "/public/build/" : 'http://' + host + ':' + devServerPort + '/public/build/',
        }),
        new webpack.ContextReplacementPlugin(/moment[/\\]locale$/, /ru|en/),
    ],
    optimization: {
        minimize: production,
        splitChunks: {
            cacheGroups: {
                default: false,
                vendors: {
                    test: /[\\/]node_modules[\\/].*js/,
                    priority: 1,
                    name: "vendor",
                    chunks: "initial",
                    enforce: true
                },
            },
        },
    }
};

if (production) {
    config.plugins.push(
        new webpack.DefinePlugin({
            'process.env': { NODE_ENV: JSON.stringify('production') }
        }),
        new CompressionPlugin({
            algorithm: "gzip",
            test: /\.js$|\.css$/,
            threshold: 4096,
            minRatio: 0.8
        })
    );
    config.output.publicPath = '/public/build/';
} else {
    config.plugins.push(
        new webpack.NamedModulesPlugin(),
    )

    config.devServer = {
        stats: 'minimal',
        port: devServerPort,
        disableHostCheck: true,
        headers: { 'Access-Control-Allow-Origin': '*' },
    };

    config.output.publicPath = 'http://' + host + ':' + devServerPort + '/public/build/';
    // Source maps
    config.devtool = 'source-map';
}

module.exports = config
