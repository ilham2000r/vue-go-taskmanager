import api from './axios'

export const taskAPI = {
  createTask: (taskData) => { 
    return api.post('/api/tasks', 
      taskData,
        { headers: { Authorization: `Bearer ${localStorage.getItem('token')}` } })
  },
  updateTask: (taskData) => {
    return api.put(`/api/tasks/${taskData.ID}`, 
      taskData,
      { headers: { Authorization: `Bearer ${localStorage.getItem('token')}` } } )
  },
  updateStatus: (taskData) => {
    return api.patch(`/api/tasks/${taskData.ID}/status`, 
      { status: taskData.status },
      { headers: { Authorization: `Bearer ${localStorage.getItem('token')}` } })
  },
  deleteTask: (taskId) => {
    return api.delete(`/api/tasks/${taskId}`,
      { headers: { Authorization: `Bearer ${localStorage.getItem('token')}` } })
  },
  listTasks: () => {
    return api.get('/api/tasks',
      { headers: { Authorization: `Bearer ${localStorage.getItem('token')}` } }
    )
  },
  getTaskById: (taskId) => {
    return api.get(`/api/tasks/${taskId}`,
      { headers: { Authorization: `Bearer ${localStorage.getItem('token')}` } }
    )
  },
  // searchTasks: (searchParams) => {
  //   return api.get('/api/tasks/search', 
  //       { params: searchParams })
  // }
}