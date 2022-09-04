import { Configuration } from "webpack";
import Path from "path";
import createBaseConfig from "./webpack.base";
import createDevConfig from "./webpack.dev";
import createProdConfig from "./webpack.prod";

type WebpackConfigutation = Configuration & {
  devServer?: Record<string, any>
}

export type ICreateConfiguration = (
    env: Record<string, string>,
    argv: WebpackConfigutation
) => Promise<WebpackConfigutation>;

const createConfig: ICreateConfiguration = async (env) => {
  const config: WebpackConfigutation = {
    module: {
      rules: [],
    },
    plugins: [],
  };
  const nodeEnv = env.NODE_ENV;

  config.context = Path.resolve(__dirname, "../../");
  config.mode = nodeEnv === "development" ? "development" : "production";

  await createBaseConfig(env, config);
  await (
    nodeEnv === "development" ? createDevConfig : createProdConfig
  )(env, config);

  return config;
};

module.exports = createConfig;
