<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
const token = localStorage.getItem('token')

const router = useRouter()

const handleLogout = async () => {
    const isConfirmed = confirm("Are you sure to logout ?")
        if (!isConfirmed) return
    try {
        await authStore.logout()
        console.log('Token:', token)
        router.push({ name : "login" })
    } catch (err) {
        console.log("Logout failed : ",err);
        
    }
}
</script>

<template>
    <div 
        @click="handleLogout"
        class="bg-zinc-800 p-2 px-6 rounded-lg shadow-md hover:scale-105">
        <button class="font-bold text-2xl  text-transparent bg-clip-text bg-gradient-to-r from-rose-500 to-blue-500  ">Logout</button>
    </div>
</template>