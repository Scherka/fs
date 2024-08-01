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
    extensions: ['.tsx', '.ts', '.js'],
  },
  output: {
    filename: `bundle.[contenthash].js`,
    path: path.resolve(__dirname, 'static'),
  },
  plugins: [new HtmlWebpackPlugin({
    template: '/static/template.html', // your HTML template file
    filename: '../bundle/bundle.html', // output HTML file relative to output.path
  })],
  devServer: {
    static: path.join(__dirname, "static"),
    compress: true,
    port: 4000,
  },
};