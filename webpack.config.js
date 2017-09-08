var webpack = require("webpack");
var CopyWebpackPlugin = require("copy-webpack-plugin");
var ExtractTextPlugin = require("extract-text-webpack-plugin");

module.exports = {
    entry: {
        admin: "./assets/js/admin.js",
        frontend: "./assets/js/frontend.js",
        "admin-style": "./assets/scss/admin.scss",
        "frontend-style": "./assets/scss/frontend.scss"
    },

    output: {
        filename: "[name].js",
        path: __dirname + "/public/assets"
    },
    plugins: [
        new webpack.ProvidePlugin({
            $: 'jquery',
            jQuery: 'jquery',
            'window.jQuery': 'jquery',
            Popper: ['popper.js', 'default'],
            Util: "exports-loader?Util!bootstrap/js/dist/util",
            Collapse: "exports-loader?Collapse!bootstrap/js/dist/collapse",
            Alert: "exports-loader?Alert!bootstrap/js/dist/alert",
            Dropdown: "exports-loader?Dropdown!bootstrap/js/dist/dropdown",
            Tooltip: "exports-loader?Tooltip!bootstrap/js/dist/tooltip",
            Tab: "exports-loader?Tab!bootstrap/js/dist/tab",
        }),
        new ExtractTextPlugin({
            filename: "[name].css"
        }),
        new CopyWebpackPlugin(
            [
                {
                    from: "./assets",
                    to: ""
                }
            ],
            {
                ignore: ["scss/**/*", "js/**/*"]
            }
        )
    ],
    module: {
        rules: [
            {
                test: /\.jsx?$/,
                loader: "babel-loader",
                options: {
                    presets: ["env"]
                },
                exclude: /node_modules/
            },
            {
                test: /\.scss$/,
                use: ExtractTextPlugin.extract({
                    fallback: "style-loader",
                    use: [
                        {
                            loader: "css-loader",
                            options: {
                                sourceMap: true
                            }
                        },
                        {
                            loader: "sass-loader",
                            options: {
                                sourceMap: true
                            }
                        }
                    ]
                })
            },
            {
                test: /\.css$/,
                use: ["style-loader", "css-loader"]
            },
            {
                test: /\.woff(\?v=\d+\.\d+\.\d+)?$/,
                use: "url-loader?limit=10000&mimetype=application/font-woff"
            },
            {
                test: /\.woff2(\?v=\d+\.\d+\.\d+)?$/,
                use: "url-loader?limit=10000&mimetype=application/font-woff"
            },
            {
                test: /\.ttf(\?v=\d+\.\d+\.\d+)?$/,
                use: "url-loader?limit=10000&mimetype=application/octet-stream"
            },
            {
                test: /\.eot(\?v=\d+\.\d+\.\d+)?$/,
                use: "file-loader"
            },
            {
                test: /\.svg(\?v=\d+\.\d+\.\d+)?$/,
                use: "url-loader?limit=10000&mimetype=image/svg+xml"
            },
            {
                test: require.resolve("jquery"),
                use: "expose-loader?jQuery!expose-loader?$"
            }
        ]
    }
};
