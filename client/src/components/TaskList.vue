<script setup>
import { ref, onMounted, computed } from 'vue';
import { useTaskStore } from '@/stores/task';
import { RouterLink } from 'vue-router';
import { Trash } from 'lucide-vue-next'
import { toast } from 'vue3-toastify';
import { differenceInDays } from 'date-fns';
import dayjs from 'dayjs';

const taskStore = useTaskStore()
const searchQuery = ref('')
const selectedPriority = ref('all')
const currentDate = new Date();

const fetchTasks = async () => {
    try {
        await taskStore.listTasks();
        console.log('Tasks fetched:', taskStore.tasks);
    } catch (err) {
        console.log('Error fetching tasks: ', err);
    }
}

const handleDelete = async (taskId) => {
    const isDeleted = confirm('Are you sure to delete task ?')
    if (!isDeleted) {
        return
    }
    try {
        await taskStore.deleteTask(taskId);
        fetchTasks()
        console.log('Delete Success');
        toast('Task Deleted')
        
    } catch (err) {
        console.log('Error fetching tasks: ', err);
    }
}

onMounted(fetchTasks)

const filteredTasks = computed(()=>{
    return taskStore.tasks.filter(task => {
        const matchesSearch = task.taskName.toLowerCase().includes(searchQuery.value.toLowerCase())
        const matchesPriority = selectedPriority.value === 'all' || task.priority === selectedPriority.value
        return matchesSearch && matchesPriority;
    })
})


</script>

<template>
    <div>
        <div class="mx-5 px-4 mb-2">
            <div class="flex justify-end gap-5 mb-5">
                <div>
                    <input 
                        v-model="searchQuery"
                        class="rounded-md border-2 p-2"
                        type="text"
                        placeholder="Search your task">
                </div>
                <div>
                    <select 
                        v-model="selectedPriority"
                        class="rounded-md border-2 p-2"
                    >
                        <option value="all">All</option>
                        <option value="high">High</option>
                        <option value="medium">Medium</option>
                        <option value="low">Low</option>
                    </select>
                </div>
            </div>
            <div class="flex">
                <div class="flex justify-between w-11/12">
                    <!-- No -->
                    <div class="w-1/12 text-center">
                        <p>No</p>
                    </div>
                    <!-- Status -->
                    <div class="w-2/12 text-center">
                        <p>Status</p>
                    </div>
                    <!-- Name -->
                    <div class="w-4/12 text-center ">
                        <p>Task Name</p>
                    </div>
                    <!-- Priority -->
                    <div class="w-2/12 text-center">
                        <p>Priority</p>
                    </div>
                    <!-- Due date -->
                    <div class="w-3/12 text-center">
                        <p>Due Date</p>
                    </div>
                </div>
                <div class="w-1/12 text-center">
                    <p></p>
                </div>
            </div>
        </div>
        <div 
            v-for="(task, index) in filteredTasks" 
            :key="task.id"
            :task="task"
            :index="index"  
            :class="{
            'border-r-8 border-r-rose-500 hover:bg-rose-200': differenceInDays(task.dueDate, currentDate) <=3 , 
            'mx-5 p-4 my-2 border-2 rounded-md shadow-lg hover:scale-105 hover:bg-zinc-100 duration-300 flex': true }"
            >
            <RouterLink
                :to="{ name : 'detailTask', params: {id: task.ID }}"
                class="flex justify-between w-11/12">
                <!-- No -->
                <div class="w-1/12 flex items-center justify-center">
                    <p class="font-light">{{ index+1 }}</p>
                </div>
                <!-- Status -->
                <div 
                    class="w-2/12 shadow-md flex items-center justify-center"
                    :class="{
                    'bg-orange-200 rounded-full text-orange-700  ': !task.taskStatus, 
                    'bg-green-300 rounded-full text-green-700  ': task.taskStatus   
                    }">
                    <p class="block sm:inline md:hidden lg:hidden">{{ task.taskStatus ? 'C' : 'IP'  }}</p>
                    <p class="hidden sm:hidden md:inline lg:inline">{{ task.taskStatus ? 'Complete' : 'In Process'  }}</p>
                </div>
                <!-- Name -->
                <div class="w-4/12 flex items-center ml-2 truncate overflow-hidden whitespace-nowrap">
                    <p class="font-light">{{ task.taskName }}</p>
                </div>
                <!-- Priority -->
                <div 
                    :class="{
                        'bg-red-300': task.priority === 'high',  
                        'bg-orange-300': task.priority === 'medium',  
                        'bg-green-300': task.priority === 'low'
                    }"
                    class="w-2/12 flex items-center justify-center text-center shadow-md rounded-full">
                    <p class="font-light">{{ task.priority }}</p>
                </div>
                <!-- Due date -->
                <div class="w-3/12 flex items-center justify-center text-center">
                    <p class="font-light">{{ dayjs(task.dueDate).format('DD-MM-YYYY') }}</p>
                </div>
                <!-- Manage -->
            </RouterLink>
            <div 
                @click="handleDelete(task.ID)"
                class="w-1/12 flex justify-center p-2 rounded-md hover:bg-red-500 hover:text-white hover:cursor-pointer">
                <Trash 
                    class="rounded-md hover:bg-red-500 "
                />
            </div>
                
        </div>
        <div 
            v-if="filteredTasks.length === 0" class="text-center py-10 text-gray-500">
                No tasks found
        </div> 
        
    </div>
</template>