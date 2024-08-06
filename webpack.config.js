const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
module.exports = {
  entry: path.resolve('static','src','index.ts'),
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
    path: path.resolve(__dirname, 'static','bundle'),
  },
  plugins: [new HtmlWebpackPlugin({
    template: path.resolve('static','template.html'), 
    filename: 'index.html',
  })],
};