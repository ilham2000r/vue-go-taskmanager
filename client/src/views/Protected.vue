<script setup>
import { onMounted,ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()

let countdown = ref(3)

const interval = setInterval(()=> {
    countdown.value --
    if (countdown.value < 0) {
        clearInterval(interval)
    }
},1000)


const fetchProtect = () => {
    try {
        setTimeout(() => {
            router.push({ name: "login"})
        }, 3000);
    } catch (err) {
        console.log(err);
    }
}

onMounted(fetchProtect)

</script>

<template>
    <div 
        class="fixed top-0 z-[-2] h-screen w-screen bg-neutral-950 
        bg-[radial-gradient(ellipse_80%_80%_at_50%_-20%,rgba(120,119,198,0.3),rgba(255,255,255,0))] ">
    </div>
    <div class="flex flex-col justify-center items-center h-screen">
        <div class="text-white text-lg font-light">You don't have permission. Please log in first.</div>
        <div class="text-white text-lg font-light">You will be redirected to login page in {{ countdown }}</div>
    </div>
</template>