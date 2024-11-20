import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'
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
    // ตรวจสอบว่ามี hash (fragment) หรือไม่
    if (savedPosition) {
      return savedPosition; // ถ้าผู้ใช้กดปุ่ม back/forward ให้เลื่อนไปตำแหน่งเดิม
    } else if (to.hash) {
      return {
        el: to.hash, // เลื่อนไปยังตำแหน่งที่มี id ตรงกับ hash
        behavior: 'smooth' // เลื่อนแบบนุ่มนวล
      };
    } else {
      return { top: 0 }; // ถ้าไม่มี hash ให้เลื่อนไปบนสุดของหน้า
    }
  },
  // Navigation Guard
})

// router.beforeEach((to, from, next) => {
//   const authStore = useAuthStore()
  
//   if (to.meta.requiresAuth && !authStore.isAuthenticated) {
//     next('/login')
//   } else {
//     next()
//   }
// })

router.beforeEach((to, from, next) => {
  const isAuthenticated = !!localStorage.getItem('token'); // ตรวจสอบว่ามี token หรือไม่

  if (to.meta.requiresAuth && !isAuthenticated) {
    // หาก route ต้องการ auth และไม่มี token
    next({ name: 'protected' }); // ส่งผู้ใช้ไปหน้า login
  } else {
    // อนุญาตให้เข้าถึง
    next();
  }
});




export default router