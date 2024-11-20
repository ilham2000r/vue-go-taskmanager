import { defineStore } from 'pinia'
import { authAPI } from '../api/auth'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    token: localStorage.getItem('token'),
    isAuthenticated: false,
    error: null
  }),
  
  actions: {
    async register(username, password) {
      try {
        const response = await authAPI.register(username, password)
        this.token = response.data.token
        this.user = response.data.user
        this.isAuthenticated = true
        localStorage.setItem('token', this.token)
      } catch (error) {
        this.error = error.response?.data?.message || 'Registration failed'
        throw error
      }
    },
    
    async login(username, password) {
      try {
        const response = await authAPI.login(username, password)
        this.token = response.data.token
        this.user = response.data.user
        this.isAuthenticated = true
        localStorage.setItem('token', this.token)
      } catch (error) {
        this.error = error.response?.data?.message || 'Login failed'
        throw error
      }
    },
    
    logout() {
      this.user = null
      this.token = null
      this.isAuthenticated = false
      localStorage.removeItem('token')
    }
  }
})