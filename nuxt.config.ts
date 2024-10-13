// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: {
    enabled: true,
    timeline: {
      enabled: true,
    },
  },
  runtimeConfig: {
    databaseUrl: '',
    public: {
      url: '',
    },
    google: {
      clientId: '',
      clientSecret: '',
    },
    stripe: {
      publishableKey: '',
      secretKey: '',
      webhookSecret: '',
    },
  },
  modules: ['@nuxt/ui', '@nuxtjs/tailwindcss', '@nuxt/icon'],
  ui: {
    global: true,
  },
})