import { DefinePlugin } from 'webpack'
import HtmlWebpackPlugin from 'html-webpack-plugin'
import Path from 'path'
import { ICreateConfiguration } from './index'

const createDevConfig: ICreateConfiguration = async (env, config) => {
  config.devtool = "eval-cheap-source-map";
  config.watch = false

  config.devServer = {
    open: true,
    port: env.port,
    hot: true,
  }

  config.plugins?.push(
    new DefinePlugin({
      nodeEnv: "development",
    }),
  )

  // config.plugins?.push(new HotModuleReplacementPlugin())
  config.plugins?.push(new HtmlWebpackPlugin({
    template: Path.resolve(__dirname, "../../public/index.html"),
  }))

  return config
}

export default createDevConfig
