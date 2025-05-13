import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    // vueDevTools(),
    AutoImport({
      resolvers: [ElementPlusResolver()],
    }),
    Components({
      resolvers: [ElementPlusResolver({
        importStyle: 'sass'
      })]
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: `@use "@/assets/styles/element-plus.scss" as *;`,
      }
    }
  },
  build: {
    assetsDir: '.',
    rollupOptions: {
      output: {
        entryFileNames: `static/assets/js[name].[hash].js`, // форматируем имена entry файлов
        chunkFileNames: `static/assets/js[name].[hash].js`, // форматируем имена chunk файлов
        assetFileNames: `static/assets/[ext][name].[hash].[ext]` // форматируем имена asset файлов
      }
    }
  }
})
