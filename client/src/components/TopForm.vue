<script setup>
import { computed, ref } from 'vue';
import { useTaskStore } from '@/stores/task'
import { toast } from 'vue3-toastify';

const taskStore = useTaskStore()

const formData = ref({
    taskName: '',
    description: '',
    dueDate: '',
    priority: 'medium'
})
 
const handleSubmit = async () => {
    try {
        if (!isFormValid.value) {
            toast('Please fill in all required fields.')
            return;
        }
        // console.log('Form Data:', formData.value)
        console.log(formData.value)
        console.log('Before creating task:', taskStore.tasks)
        await taskStore.createTask(formData.value)
        toast('Task Created')
        console.log('After creating task:', taskStore.tasks)
        // clear input when finished
        formData.value = {
            taskName: '',
            description: '',
            dueDate: '',
            priority: 'medium'
        }       
    } catch (err) {
        console.log(err);
        
    }
}

const isFormValid = computed(() => {
    return formData.value.taskName.trim() !== '' && 
           formData.value.dueDate !== '' && 
           formData.value.priority !== '';
})

</script>

<template>
    <div>
        <p class=" m-5 ml-7 text-2xl font-semibold">
            Create Your task
        </p>
        <form @submit.prevent="handleSubmit" class="flex flex-col m-5 gap-2 justify-between">
            <div class="">
                <div class="">
                    <p class="font-light">Task Name</p>
                    <input
                        v-model="formData.taskName"
                        class="my-2 w-full border-2 shadow-sm rounded-md p-1" 
                        type="text"
                        >
                </div>
                <div>
                    <p class="font-light">Description</p>
                    <textarea
                        v-model="formData.description"
                        class="my-2 w-full border-2 shadow-sm rounded-md resize-none h-24" >
                    </textarea>
                </div>
            </div>
            <div class="flex ">
                <div class="flex-1">
                    <p class="font-light">Due Date</p>
                    <input 
                        v-model="formData.dueDate"
                        class="my-2 border-2 shadow-sm rounded-md text-zinc-400 p-1" 
                        type="date"
                        >
                </div>
                <div class="">
                    <p class="font-light">Priority</p>
                    <select v-model="formData.priority" class="my-2 border-2 px-4 py-1">
                        <option value="high">High</option>
                        <option value="medium">Medium</option>
                        <option value="low">Low</option>
                    </select>
                </div>
            </div>
            <div class="flex justify-center" >
                <button 
                    @disable="!isFormValid"
                    class="bg-sky-500 w-full mx-5 p-2 rounded-md hover:bg-sky-600 hover:text-white duration-200">Create</button>
            </div>
        </form>

    </div>
</template>