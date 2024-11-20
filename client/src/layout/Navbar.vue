<script setup>
import { RouterLink } from 'vue-router';
import Logout from '@/components/Logout.vue';

</script>

<template>
    <div class="flex m-5 justify-between">
        <RouterLink 
        :to="{name:'indexView'}"
        class="">
            <p class="text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-blue-500 to-green-500 ">Task's</p>
        </RouterLink>
        <Logout />
    </div>
    
    
    
</template>