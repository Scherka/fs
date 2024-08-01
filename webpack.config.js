const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
module.exports = {
  entry: '/static/src/index.ts',
  module: {
    rules: [
      {
        test: /\.ts?$/,
        use: 'ts-loader',
        exclude: /node_modules/,
      },
    ],
  },
  resolve: {
    alias: {
      // aliases used in the template
      Bundle: path.resolve(__dirname, 'bundle/bundle.js'),
    },
    extensions: ['.tsx', '.ts', '.js'],
  },
  output: {
    filename: `bundle.[contenthash].js`,
    path: path.resolve(__dirname, 'bundle'),
  },
  plugins: [new HtmlWebpackPlugin({
    template: './static/page.html', // your HTML template file
    filename: '../static/page.html', // output HTML file relative to output.path
  })],
  devServer: {
    static: path.join(__dirname, "static"),
    compress: true,
    port: 4000,
  },
};