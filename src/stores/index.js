import { createPinia } from 'pinia'
// 必须导入持久化插件
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
const pinia = createPinia()

// 必须注册插件！
pinia.use(piniaPluginPersistedstate)
export * from './modules/user'
export * from './modules/counter'
export default pinia
