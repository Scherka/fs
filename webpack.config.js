const path = require('path');

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
    filename: 'main.js',
    path: path.resolve(__dirname, 'static'),
  },
  devServer: {
    static: path.join(__dirname, "static"),
    compress: true,
    port: 4000,
  },
};