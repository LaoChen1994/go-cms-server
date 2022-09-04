import { DllPlugin, Configuration, ProgressPlugin } from 'webpack'
import Path from 'path'

const config: Configuration = {
  mode: "production",
  entry: {
    dll: [
      'lodash',
      'react',
      'react-dom',
      'react-router',
      'mobx',
      'mobx-react',
    ],
  },
  output: {
    filename: 'dll.bundle.js',
    path: Path.resolve(__dirname, "../../dist"),
  },
  plugins: [
    new DllPlugin({
      context: Path.resolve(__dirname, "../../dist"),
      name: "[name]_[fullhash]",
      path: Path.join(__dirname, "../../dist/manifest.json"),
    }),
    new ProgressPlugin({
      activeModules: false,
      entries: true,
      modules: true,
      modulesCount: 5000,
      profile: false,
      dependencies: true,
      dependenciesCount: 10000,
      percentBy: "modules",
    }),
  ],
}

module.exports = config
