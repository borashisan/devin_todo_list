export default defineNuxtConfig({
  devtools: { enabled: true },
  ssr: true,
  modules: ['@nuxtjs/tailwindcss'],
  devServer: {
    host: '0.0.0.0',
    port: 3000
  }
})
