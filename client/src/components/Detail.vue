<script setup>
import { useRoute,RouterLink,useRouter } from 'vue-router'
import { ref, onMounted } from 'vue'
import {  useTaskStore } from '../stores/task'
import { toast } from 'vue3-toastify';
import dayjs from 'dayjs';

const route = useRoute()
const routerIns = useRouter()
const taskStore = useTaskStore()
const taskDetail = ref({
  taskName: '',
  description: '',
  CreateAt: '',
  updatedAt: '',
  dueDate: '',
  priority: '',
});

const dayLeft = ref(0);

const fetchTaskDetail = async () => {
  try {
    const taskId = route.params.id; 
    await taskStore.listTaskById(taskId); 
    taskDetail.value = taskStore.currentTask;
    console.log('taskStore.currentTask:', taskStore.currentTask);
    
    // cal Day Left
    if (taskDetail.value.dueDate) {
      const currentDate = new Date(); 
      const dueDate = new Date(taskDetail.value.dueDate); 
      const timeDiff = dueDate.getTime() - currentDate.getTime(); 
      dayLeft.value = Math.ceil(timeDiff / (1000 * 60 * 60 * 24)); 
    }

    // convert date 
    if (taskDetail.value.dueDate) {
    const date = new Date(taskDetail.value.dueDate); 
    taskDetail.value.dueDate = date.toISOString().split('T')[0];
    }
  } catch (err) {
    console.log('Error fetching task detail:', err);
  }
  
}

const handleUpdate = async () =>{
    try {
        const updateData = {
            ID: taskDetail.value.ID,
            taskName: taskDetail.value.taskName,
            description: taskDetail.value.description,
            dueDate: taskDetail.value.dueDate,
            priority: taskDetail.value.priority
        };
        await taskStore.updateTask(updateData)
        routerIns.push({ name : "indexView" })
        toast('Update Success')
        console.log('Task updated successfully');
    } catch (err) {
        console.log("err to updated", err);
        
    }
}

const updateStatus = async (taskId) => {
    try {
        console.log(taskDetail.value.taskStatus);
        
        const taskData = {
            ID: taskId,
            status: !taskDetail.value.taskStatus 
        };
        await taskStore.updateStatus(taskData); 
        routerIns.push({ name : "indexView" })
        toast('Status Updated')
        console.log(`Task ${taskId} status updated to ${!taskDetail.value.taskStatus}`)
    } catch (err) {
        console.log('Error updating task status: ', err);
    }
}

onMounted(fetchTaskDetail)


</script>

<template>
    <div>
        <div class="shadow-md bg-white rounded-md my-10 m-5 p-5">
            <!-- Header -->
            <div class="flex justify-between">
                <!-- left header -->
                <div class="m-2 ml-5">
                    <p class="font-bold text-2xl ">{{ taskDetail.taskName  }}</p>
                </div>
                <!-- right header -->
                <div class="m-2 mr-5">
                    <p class="text-end">Day Left</p>
                    <p :class="{
                        'text-red-600 ': dayLeft <= 3,
                        'font-bold text-4xl text-end mt-2': true}">{{ dayLeft }}</p>
                </div>
            </div>
            <!-- middle content -->
             <div>
                <!-- Header of detail -->
                 <div>
                    <div class="flex flex-col m-5 gap-2 justify-between">
                        <div class="flex mb-3">
                            <div class="flex-1">
                                <p class="font-light mb-2">Created At</p>
                                <p >{{  dayjs(taskDetail.CreatedAt).format('DD-MM-YYYY') }}</p>
                            </div>
                            <div class="">
                                <p class="font-light mb-2">Last Update</p>
                                <p >{{ dayjs(taskDetail.UpdatedAt).format('DD-MM-YYYY') }}</p>
                            </div>
                        </div>
                        <div>
                            <div class="">
                                <p class="font-light">Task Name</p>
                                <input
                                    v-model="taskDetail.taskName"
                                    class="my-2 w-full border-2 shadow-sm rounded-md p-1" 
                                    type="text"
                                    >
                            </div>
                            <div class="">
                                <p class="font-light">Description</p>
                                <textarea
                                    v-model="taskDetail.description"
                                    class="my-2 p-2 w-full border-2 shadow-sm rounded-md resize-none h-24" >
                                </textarea>
                            </div>
                        </div>
                        <div class="flex ">
                            <div class="flex-1">
                                <p class="font-light">Due Date</p>
                                <input 
                                    v-model="taskDetail.dueDate"
                                    class="my-2 border-2 shadow-sm rounded-md text-zinc-400 p-1" 
                                    type="date"
                                    >
                            </div>
                            <div class="flex-1 flex flex-col justify-center items-center">
                                <p class="font-light">Task Status</p>
                                <div 
                                    @click="updateStatus(taskDetail.ID, taskDetail.status)"
                                    class="w-7/12 p-2 px-3 text-center shadow-md mt-2 hover:cursor-pointer"
                                    :class="{
                                    'bg-orange-200 rounded-full text-orange-700  hover:bg-orange-700 hover:text-white': !taskDetail.taskStatus, 
                                    'bg-green-300 rounded-full text-green-700  hover:bg-green-700 hover:text-white ': taskDetail.taskStatus   
                                    }">
                                    <p class="block sm:inline md:hidden lg:hidden">{{ taskDetail.taskStatus ? 'C' : 'IP'  }}</p>
                                    <p class="hidden sm:hidden md:inline lg:inline">{{ taskDetail.taskStatus ? 'Complete' : 'In process'  }}</p>
                                </div>
                            </div>
                            <div class="">
                                <p class="font-light">Priority</p>
                                <select 
                                    v-model="taskDetail.priority"
                                    class="my-2 border-2 px-4 py-1">
                                    <option value="high">High</option>
                                    <option value="medium">Medium</option>
                                    <option value="low">Low</option>
                                </select>
                            </div>
                        </div>
                    </div>
                 </div>
                 <!-- Button -->
                 <div class="flex gap-2 mx-5">
                    <button 
                        @click="handleUpdate"
                        class="flex-1 text-center bg-blue-500 p-2 rounded-md shadow-md hover:bg-blue-700 hover:text-white duration-200">Update</button>
                    <RouterLink
                        :to="{ name: 'indexView'}"
                        class="flex-1 text-center bg-zinc-300 p-2 rounded-md shadow-md hover:bg-zinc-500 hover:text-white duration-200">Back</RouterLink>
                 </div>
             </div>
        </div>
    </div>
</template>