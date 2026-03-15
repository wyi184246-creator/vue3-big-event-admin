import { defineStore } from 'pinia'
import { ref } from 'vue'
import { userGetProfileService } from '@/api/user'
export const useUserStore = defineStore(
  'big-user',
  () => {
    const token = ref('')
    const setToken = (newToken) => {
      token.value = newToken
    }
    const removeToken = () => {
      token.value = ''
    }
    const user = ref({})
    const getUser = async () => {
      const res = await userGetProfileService()
      user.value = res.data.data
    }
    const setUser = (newUser) => {
      user.value = newUser
    }
    return {
      token,
      setToken,
      removeToken,
      user,
      getUser,
      setUser
    }
  },
  {
    persist: true
  }
)
