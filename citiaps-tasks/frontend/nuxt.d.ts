import type { RuntimeConfig } from 'nuxt'

declare global {
  namespace NodeJS {
    interface ProcessEnv {
      NUXT_PUBLIC_API_BASE?: string
      NODE_ENV?: 'development' | 'production' | 'test'
    }
  }

  interface Process {
    env: NodeJS.ProcessEnv
  }

  const process: Process
  const useRuntimeConfig: () => RuntimeConfig
  const $fetch: typeof import('ofetch').$fetch
  const defineNuxtConfig: typeof import('nuxt').defineNuxtConfig
  const defineNuxtPlugin: typeof import('nuxt').defineNuxtPlugin
  const defineNuxtMiddleware: typeof import('nuxt').defineNuxtMiddleware
  const defineNuxtLayout: typeof import('nuxt').defineNuxtLayout
  const useRouter: typeof import('#app').useRouter
  const useRoute: typeof import('#app').useRoute
  const useState: typeof import('#app').useState
  const useHead: typeof import('vue-meta').useHead
  const useFetch: typeof import('#app').useFetch
  const navigateTo: typeof import('#app').navigateTo
}

export {}
