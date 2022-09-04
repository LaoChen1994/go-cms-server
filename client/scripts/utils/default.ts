import { ModuleOptions, RuleSetUseItem, Configuration } from "webpack";
import MiniCssExtractPlugin from "mini-css-extract-plugin";
import { TypedCssModulesPlugin } from "typed-css-modules-webpack-plugin";

type Rules = ModuleOptions["rules"];
type Plugins = Configuration["plugins"];

export const getDefaultTsRules = (useEsbuild: boolean = false) => {
  const rules: Rules = [];

  if (useEsbuild) return { rules };

  rules.push({
    test: /\.tsx?/,
    exclude: /node_modules|bower_components/,
    use: [
      {
        loader: "thread-loader",
        options: {
          workerParallelJobs: 50,
          workerNodeArgs: ["--max-old-space-size=1024"],
          poolParallelJobs: 50,
        },
      },
      {
        loader: "babel-loader",
        options: {
          presets: [
            [
              "@babel/preset-env",
              {
                useBuiltIns: "entry",
                corejs: "3",
                targets: { chrome: "58", ie: "11" },
              },
            ],
            [
              "@babel/preset-react",
              {
                runtime: "automatic",
              },
            ],
            [
              "@babel/preset-typescript",
              {
                isTSX: true,
                allExtensions: true,
              },
            ],
          ],
          plugins: [
            "@babel/plugin-transform-runtime",
            "@babel/plugin-syntax-dynamic-import",
          ],
          cacheDirectory: true,
        },
      },
    ],
  });

  return { rules };
};

const getCssRule = (props: {
  useModule: boolean;
  loaders?: number;
  isDev?: boolean;
}): RuleSetUseItem[] => {
  const { useModule, loaders = 1, isDev } = props;
  return [
    isDev ? "style-loader" : MiniCssExtractPlugin.loader,
    {
      loader: "css-loader",
      options: {
        modules: useModule,
        importLoaders: loaders,
        esModule: true,
      },
    },
    {
      loader: "postcss-loader",
      options: {
        postcssOptions: {
          plugins: ["postcss-preset-env"],
        },
      },
    },
  ];
};

// @todo sass-loader 与 css module冲突问题
// @todo 待修复
export const getDefaultSassRulesAndPlugins = (isDev: boolean) => {
  const rules: Rules = [];
  const plugins: Plugins = [];

  rules.push({
    test: /\.module.css$/,
    use: [
      ...getCssRule({
        useModule: true,
        isDev,
      }),
    ],
  });

  rules.push({
    test: /\.css$/,
    exclude: /\.module\.css$/,
    use: [
      ...getCssRule({
        useModule: false,
        isDev,
      }),
    ],
  });

  rules.push({
    test: /\.module\.s(a|c)ss/i,
    use: [
      ...getCssRule({
        useModule: true,
        loaders: 2,
        isDev,
      }),
      "sass-loader",
    ],
  });

  rules.push({
    test: /\.s(a|c)ss/,
    exclude: /\.module\.s(a|c)ss/,
    use: [
      ...getCssRule({
        useModule: false,
        isDev,
        loaders: 2,
      }),
      "sass-loader",
    ],
  });

  if (isDev) {
    plugins.push(new MiniCssExtractPlugin());
  }

  plugins.push(
    new TypedCssModulesPlugin({
      globPattern: "src/**/*.scss",
      camelCase: true,
    }),
  );

  return { rules, plugins };
};

export const getDefaultLessRulesAndPlugins = (isDev: boolean) => {
  const rules: Rules = [];
  const plugins: Plugins = [];

  rules.push({
    test: /\.module.css$/,
    exclude: /node_modules/,
    use: [
      ...getCssRule({
        useModule: true,
        isDev,
      }),
    ],
  });

  rules.push({
    test: /\.css$/,
    exclude: [/\.module\.css$/, /node_modules/],
    use: [
      ...getCssRule({
        useModule: false,
        isDev,
      }),
    ],
  });

  rules.push({
    test: /\.module\.less/i,
    exclude: /node_modules/,
    use: [
      ...getCssRule({
        useModule: true,
        isDev,
        loaders: 2,
      }),
      "less-loader",
    ],
  });

  rules.push({
    test: /\.less/,
    exclude: [/\.module\.less/, /node_modules/],
    use: [
      ...getCssRule({
        useModule: false,
        loaders: 2,
        isDev,
      }),
      "less-loader",
    ],
  });

  if (isDev) {
    plugins.push(new MiniCssExtractPlugin());
  }

  plugins.push(
    new TypedCssModulesPlugin({
      globPattern: "src/**/*.less",
      camelCase: true,
    }),
  );

  return { rules, plugins };
};

export const getDefaultFileRules = () => {
  const rules: Rules = [];

  rules.push({
    test: /\.(png|jpg|gif)$/,
    use: {
      loader: "url-loader",
      options: {
        limit: 8192,
        fallback: "file-loader",
      },
    },
  });

  rules.push({
    test: /\.(ttf|eot|woff|woff2)$/,
    use: {
      loader: "file-loader",
      options: {
        name: "fonts/[name].[ext]",
      },
    },
  });

  return { rules };
};
