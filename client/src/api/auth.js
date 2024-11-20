import api from './axios'

export const authAPI = {
  register: (username, password) => {
    return api.post('/register', 
        { username, password })
  },
  login: (username, password) => {
    return api.post('/login', 
        { username, password })
  }
}