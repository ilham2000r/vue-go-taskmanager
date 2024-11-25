import { createRouter, createWebHistory } from 'vue-router'
import IndexView from '../views/IndexView.vue'
import DetailTask from '@/views/DetailTask.vue'
import Login from '@/views/Login.vue'
import Register from '@/views/Register.vue'
import Protected from '@/views/Protected.vue'

const router = createRouter({
  history: createWebHistory(),
  routes : [
      {
      path: '/',
      name: 'indexView',
      component: IndexView,
      meta: { requiresAuth: true }
    },
    {
      path: '/detail/:id',
      name: 'detailTask',
      component: DetailTask,
      meta: { requiresAuth: true }
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
    },
    {
      path: '/register',
      name: 'register',
      component: Register,
    },
    { 
      path: '/protected',
      name: 'protected', 
      component: Protected,
    },
  ],
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition; 
    } else if (to.hash) {
      return {
        el: to.hash, 
        behavior: 'smooth' 
      };
    } else {
      return { top: 0 }; 
    }
  },
})

// Navigation Guard
router.beforeEach((to, from, next) => {
  const isAuthenticated = !!localStorage.getItem('token'); 
  if (to.meta.requiresAuth && !isAuthenticated) {
    next({ name: 'protected' }); 
  } else {
    next();
  }
});

export default router