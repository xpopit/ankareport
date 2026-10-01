import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
      'ankareport': path.resolve(__dirname, '../dist/ankareport.esm.js')
    },
  },
  build: {
    lib: {
      entry: path.resolve(__dirname, 'src/index.ts'),
      name: 'XerpReportStudio',
      fileName: (format) => `xerp-report-studio.${format}.js`
    },
    rollupOptions: {
      external: ['vue', 'vue-router', 'pinia', 'echarts', 'axios'],
      output: {
        globals: {
          vue: 'Vue',
          pinia: 'Pinia',
          'vue-router': 'VueRouter'
        }
      }
    }
  }
})
