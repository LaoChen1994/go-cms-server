const Path = require('path')

module.exports = {
  env: {
    browser: true,
    es2021: true,
  },
  extends: ['plugin:react/recommended', 'airbnb', "plugin:react/jsx-runtime"],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaFeatures: {
      jsx: true,
    },
    ecmaVersion: 'latest',
    sourceType: 'module',
  },
  plugins: ['react', '@typescript-eslint'],
  settings: {
    "import/extensions": [".tsx", ".ts", ".js", ".scss", ".sass", ".css"],
    'import/resolver': {
      node: {
        paths: ["src"],
        extensions: [".tsx", ".ts", ".js", ".scss", ".sass", ".css"],
      },
      alias: {
        map: [
          ["Api/", ["./src/api/"]],
          ["Common/", ["./src/common/"]],
          ["Constant/", ["./src/constant/"]],
          ["Hooks/", ["./src/hooks/"]],
          ["Pages", [Path.resolve(__dirname, "./src/pages")]],
          ["Utils/", ["./src/utils/"]],
          ["Config/", ["./src/config/"]],
        ],
        extensions: [".tsx", ".ts", ".js", ".scss", ".sass", ".css"],
      },
    },
  },
  rules: {
    'react/jsx-filename-extension': [
      1,
      {
        extensions: ['.tsx', '.jsx', '.js'],
      },
    ],
    semi: "off",
    quotes: "off",
    "no-use-before-define": "off",
    "@typescript-eslint/no-use-before-define": ["error"],
    "no-unused-vars": "off",
    "@typescript-eslint/no-unused-vars": ["error"],
    "import/no-import-module-exports": "off",
    "import/no-extraneous-dependencies": "off",
    "import/extensions": "off",
    "react/function-component-definition": "off",
    "no-param-reassign": "off",
    "import/no-dynamic-require": "off",
    "import/no-unresolved": "off",
    "react/jsx-props-no-spreading": "off",
  },
};
