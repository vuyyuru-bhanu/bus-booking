const path = require('path');
const { merge } = require('webpack-merge');
const commonConfig = require('./webpack.common.js');

module.exports = (env, argv) => {
  const isProduction = argv.mode === 'production';

  return merge(commonConfig, {
    mode: isProduction ? 'production' : 'development',
    output: {
      filename: isProduction ? '[name].[contenthash].js' : '[name].js',
      path: path.resolve(__dirname, 'dist'),
    },
    module: {
      rules: [
        {
          test: /\.css$/,
          use: ['style-loader', 'css-loader'],
        },
        {
          test: /\.(png|svg|jpg|jpeg|gif)$/,
          use: [
            {
              loader: 'file-loader',
              options: {
                name: '[path][name].[ext]',
                context: 'src',
              },
            },
          ],
        },
      ],
    },
    plugins: [
      // Add any plugins you need here
    ],
  });
};
