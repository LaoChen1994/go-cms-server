import Path from "path";
import { ProgressPlugin, DllReferencePlugin } from "webpack";
import HtmlWebpackPlugin from "html-webpack-plugin";
import getEntries from "../utils/getEntries";
import getResolves from "../utils/getResolves";
import { ICreateConfiguration } from "./index";
import {
  getDefaultTsRules,
  getDefaultFileRules,
  getDefaultSassRulesAndPlugins,
} from "../utils/default";

const createBaseConfig: ICreateConfiguration = async (env, config) => {
  const entries = getEntries();
  const resolves = await getResolves();
  const isDev = env.NODE_ENV === 'development'

  config.entry = entries;
  config.output = {
    path: Path.resolve(__dirname, "../../dist"),
    filename: "[name]_[contenthash].js",
  };
  config.resolve = resolves;
  config.target = ["web", "es5"];

  const { rules: tsRules } = getDefaultTsRules();
  const { rules: lessRules, plugins: lessPlugins } = getDefaultSassRulesAndPlugins(isDev);
  const { rules: fileRules } = getDefaultFileRules();

    config.module!.rules?.push(...tsRules);
    config.module!.rules?.push(...lessRules);
    config.module!.rules?.push(...fileRules);

    config.plugins!.push(...lessPlugins);

    config.plugins!.push(
      new ProgressPlugin({
        activeModules: false,
        entries: true,
        modules: true,
        modulesCount: 5000,
        profile: false,
        dependencies: true,
        dependenciesCount: 10000,
        percentBy: null,
      }),
    );

    config.plugins!.push(
      new HtmlWebpackPlugin({
        template: Path.resolve(__dirname, "../../public/index.html"),
      }),
    );

    if (!isDev) {
        config.plugins!.push(
          new DllReferencePlugin({
            context: Path.resolve(__dirname, "../../dist"),
            // eslint-disable-next-line global-require
            manifest: require(Path.join(__dirname, "../../dist/manifest.json")),
            sourceType: "umd",
          }),
        );
    }

    return config;
};

export default createBaseConfig;
