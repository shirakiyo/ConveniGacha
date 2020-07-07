import colors from 'vuetify/es5/util/colors'

export default {
  mode: 'spa',
 	head: {
    titleTemplate: '%s - ' + process.env.npm_package_name,
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: process.env.npm_package_description || '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
 	components: true,
 	buildModules: [
    '@nuxtjs/vuetify',
  ],
 	modules: [
    // Doc: https://axios.nuxtjs.org/usage
    '@nuxtjs/axios',
    '@nuxtjs/pwa',
 	],
 	webfontloader: {
    google: {
      families: ["Noto+Sans+JP"]
    }
  },
 	axios: {},
 	vuetify: {
    customVariables: ['~/assets/variables.scss'],
    theme: {
      themes: {
        light: {
          base: "#000",
        }
      }
    }
  },
 	build: {
  }
}
