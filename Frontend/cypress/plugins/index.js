const injectDevServer = require("@cypress/webpack-dev-server");
const findWebpack = require("find-webpack");
const { startDevServer } = require("@cypress/webpack-dev-server");

module.exports = (on, config) => {
  on("dev-server:start", (options) => {
    const webpackConfig = findWebpack.getWebpackConfig();

    return startDevServer({
      options,
      webpackConfig,
    });
  });

  return config;
};
