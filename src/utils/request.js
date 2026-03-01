import axios from 'axios'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/stores'
import router from '@/router'
const baseURL = 'http://big-event-vue-api-t.itheima.net'
const instance = axios.create({
  baseURL,
  timeout: 10000
})
instance.interceptors.request.use(
  (config) => {
    const userStore = useUserStore()
    if (userStore.token) {
      config.headers.Authorization = userStore.token
    }
    return config
  },
  (err) => {
    return Promise.reject(err)
  }
)
instance.interceptors.response.use(
  (res) => {
    if (res.data.status === 0) {
      return res
    }
    ElMessage.error(res.data.message || '请求失败')
    return Promise.reject(res.data || '请求失败')
  },
  (err) => {
    if (err.response?.status === 401) {
      router.push('/login')
    }

    ElMessage.error(err.response.data.message || '请求失败')

    return Promise.reject(err)
  }
)
export default instance
export { baseURL }
