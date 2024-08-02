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
      {
        test: /\.css$/i,
        use: ["style-loader", "css-loader"],
      }
    ],
  },
  resolve: {
    extensions: ['.tsx', '.ts', '.js'],
  },
  output: {
    filename: `bundle.[contenthash].js`,
    path: path.resolve(__dirname, 'bundle'),
  },
  plugins: [new HtmlWebpackPlugin({
    template: '/static/template.html', // your HTML template file
    filename: '../bundle/index.html', // output HTML file relative to output.path
  })],
};