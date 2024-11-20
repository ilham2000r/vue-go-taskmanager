<script setup>
import Detail from '@/components/Detail.vue';
import Navbar from '@/layout/Navbar.vue';
</script>

<template>
    <Navbar />
    <div>
        <div class="fixed top-0 z-[-2] h-screen w-screen bg-neutral-950 
        bg-[radial-gradient(ellipse_80%_80%_at_50%_-20%,rgba(120,119,198,0.3),rgba(255,255,255,0))]
        flex justify-center items-center">
            <Detail class="w-11/12 h-5/6 "/>
        </div>
    </div>
</template>