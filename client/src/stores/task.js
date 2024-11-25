import { defineStore } from 'pinia'
import { taskAPI } from '../api/task'

export const useTaskStore = defineStore('task', {
  state: () => ({
    tasks: [],
    currentTask: null,
    loading: false,
    error: null
  }),
  
  actions: {
    async createTask(taskData) {
      this.loading = true
      try {
        const response = await taskAPI.createTask(taskData)
        this.tasks.push(response.data)
        console.log(response.data)
      } catch (error) {
        this.error = error.response?.data?.message || 'Failed to create task'
        throw error
      } finally {
        this.loading = false
      }
    },
    
    async updateTask(taskData) {
      console.log('Task ID:', taskData.ID);
      this.loading = true
      try {
        const response = await taskAPI.updateTask(taskData)
        const index = this.tasks.findIndex(task => task.ID === taskData.ID)
        if (index !== -1) {
          this.tasks[index] = response.data
        }
      } catch (error) {
        this.error = error.response?.data?.message || 'Failed to update task'
        throw error
      } finally {
        this.loading = false
      }
    },
    
    async updateStatus(taskData) {
      try {
        console.log('API call with taskData:', taskData);
        await taskAPI.updateStatus(taskData)
        const index = this.tasks.findIndex(task => task.ID === taskData.ID)
        if (index !== -1) {
          this.tasks[index].status = taskData.status
        }
      } catch (error) {
        this.error = error.response?.data?.message || 'Failed to update status'
        throw error
      }
    },
    
    async deleteTask(taskId) {
      try {
        await taskAPI.deleteTask(taskId)
        this.tasks = this.tasks.filter(task => task.id !== taskId)
      } catch (error) {
        this.error = error.response?.data?.message || 'Failed to delete task'
        throw error
      }
    },
    
    async listTasks() {
      this.loading = true
      try {
        // console.log('Call API');
        const response = await taskAPI.listTasks()
        // console.log('API Response', response);
        this.tasks = response.data
      } catch (error) {
        this.error = error.response?.data?.message || 'Failed to fetch tasks'
        throw error
      } finally {
        this.loading = false
      }
    },
    
    async listTaskById(taskId) {
      this.loading = true
      try {
        const response = await taskAPI.getTaskById(taskId)
        this.currentTask = response.data
      } catch (error) {
        this.error = error.response?.data?.message || 'Failed to fetch task'
        throw error
      } finally {
        this.loading = false
      }
    },

  }
})