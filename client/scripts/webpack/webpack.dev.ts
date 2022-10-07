import HtmlWebpackPlugin from 'html-webpack-plugin'
import Path from 'path'
// eslint-disable-next-line import/no-cycle
import { ICreateConfiguration } from '.'

const createDevConfig: ICreateConfiguration = async (env, config) => {
  config.devtool = "eval-cheap-source-map";
  config.watch = false

  config.devServer = {
    open: true,
    port: env.port,
    hot: true,
    historyApiFallback: true,
  }

  config.plugins?.push(new HtmlWebpackPlugin({
    template: Path.resolve(__dirname, "../../public/index.html"),
  }))

  return config
}

export default createDevConfig
