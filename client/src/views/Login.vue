<script setup>
import { ref } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { RouterLink, useRouter } from 'vue-router'
import { toast } from 'vue3-toastify';

const authStore = useAuthStore()
const router = useRouter()

const formUser = ref({
    username: '',
    password: ''
})

const handleLogin = async () => {
    try {
        console.log(formUser.value.username)
        await authStore.login(formUser.value.username, formUser.value.password)
        // clear input when finished
        formUser.value = {
            username: '',
            password: ''
        }
        router.push({ name : "indexView" })
        toast('Welcome back')
        console.log('Login Success');
    } catch (err) {
        toast("Username or Password invalid")
        console.log(err);
    }
}

</script>

<template>
    <div 
        class="fixed top-0 z-[-2] h-screen w-screen bg-neutral-950 
        bg-[radial-gradient(ellipse_80%_80%_at_50%_-20%,rgba(120,119,198,0.3),rgba(255,255,255,0))] ">
    </div>
    <div class="flex justify-center items-center h-screen">
        <div class="shadow-md bg-white rounded-md m-5 p-10">
            <P class="text-3xl font-semibold mb-5">
                Login
            </P>
            <div class="flex gap-4">
                <div>
                    <p>Username</p>
                    <input
                        v-model="formUser.username"
                        class="rounded-md border-2 shadow-md p-1 mt-2" 
                        type="text">
                </div>
                <div>
                    <p>Password</p>
                    <input
                        v-model="formUser.password"
                        class="rounded-md border-2 shadow-md p-1 mt-2" 
                        type="text">
                </div>
            </div>
            <div class="flex justify-end mt-4">
                <button 
                    @click="handleLogin"
                    class="p-2 px-5 bg-blue-500 shadow-md rounded-md hover:bg-blue-700 hover:text-white duration-200">Login</button>
            </div>
            <div class="flex justify-end mt-2 gap-2">
                <p>Don't have an account yet?</p>
                <router-link 
                :to="{name : 'register'}"    
                    class="hover:text-blue-700 duration-300">
                    <p>Register</p>
                </router-link>
            </div>
        </div>
    </div>
    
</template>