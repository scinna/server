module.exports = {
    publicPath: '/app/',

    devServer: {
        "public": "i.scinna.drx"
    },

    pluginOptions: {
      i18n: {
        locale: 'en',
        fallbackLocale: 'en',
        localeDir: 'locales',
        enableInSFC: true
      }
    }
}
