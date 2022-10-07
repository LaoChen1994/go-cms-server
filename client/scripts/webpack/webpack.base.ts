import Path from "path";
import { ProgressPlugin, DllReferencePlugin, DefinePlugin } from "webpack";
import HtmlWebpackPlugin from "html-webpack-plugin";
import getEntries from "../utils/getEntries";
import getResolves from "../utils/getResolves";
// eslint-disable-next-line import/no-cycle
import { ICreateConfiguration } from "./index";
import {
  getDefaultTsRules,
  getDefaultFileRules,
  getDefaultSassRulesAndPlugins, getDefaultLessRulesAndPlugins,
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
  const { rules: sassRule, plugins: sassRulePlugins } = getDefaultSassRulesAndPlugins(isDev);
  const { rules: lessRule, plugins: lessRulePlugins } = getDefaultLessRulesAndPlugins(isDev);
  const { rules: fileRules } = getDefaultFileRules();

    config.module!.rules?.push(...tsRules);
    config.module!.rules?.push(...sassRule);
    config.module?.rules?.push(...lessRule)
    config.module!.rules?.push(...fileRules);

    config.plugins!.push(...sassRulePlugins);
    config.plugins!.push(...lessRulePlugins)

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

    config.plugins!.push(new DefinePlugin({
      'process.env.START_ENV': JSON.stringify(isDev ? 'development' : 'production'),
    }))

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
