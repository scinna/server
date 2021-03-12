module.exports = {
  module: {
      rules: [
          {
              test: /\.((c|sa|sc)ss)$/,
              loader: "css-loader",
              options: {
                  modules: true,
              }
          }
      ]
  }
};