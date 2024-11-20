<script setup>
import TopForm from '../components/TopForm.vue'
import TaskList from '../components/TaskList.vue'
import Navbar from '@/layout/Navbar.vue';
</script>

<template>
    <Navbar />
        <div 
            class="fixed top-0 z-[-2] h-screen w-screen bg-neutral-950 
            bg-[radial-gradient(ellipse_80%_80%_at_50%_-20%,rgba(120,119,198,0.3),rgba(255,255,255,0))] ">
        </div>
        
        <div class="flex justify-center items-center ">
            <div class="bg-white w-11/12 rounded-md my-4">
                <TopForm /> 
                <hr class="my-5 mx-5" />
                <TaskList />
            </div>
        </div>
    
</template>