import { Configuration, ids } from 'webpack'
import TerserPlugin from 'terser-webpack-plugin'
import CssMinimizerPlugin from 'css-minimizer-webpack-plugin'
import { ICreateConfiguration } from './index'

const createProdConfig: ICreateConfiguration = async (_, config) => {
  config.cache = true;

  config.optimization = {
    runtimeChunk: "single",
    splitChunks: {
      cacheGroups: {
        vendor: {
          test: /[\\/]node_modules[\\/]/,
          name: 'vendors',
          chunks: 'all',
        },
      },
    },
    moduleIds: "deterministic",
    minimize: true,
    minimizer: [
      new TerserPlugin({
        parallel: true,
        exclude: /node_modules/,
      }),
      new CssMinimizerPlugin(),
    ],
  } as Configuration['optimization']

  config.plugins?.push(
    new ids.DeterministicModuleIdsPlugin({
      maxLength: 5,
    }),
  )

  return config
}

export default createProdConfig
