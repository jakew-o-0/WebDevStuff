<script setup>
import axios from 'axios';
import { ref } from 'vue';

const props = defineProps(['isModalOpen']);
const taskName = ref("");

function PostTask() {
    axios({
        method: 'POST',
        url: 'http://localhost/task/',
        data: {
            name: taskName.value,
            completed: false
        }
    })
}
</script>

<template>
    <div class="grid h-screen w-screen bg-slate-500/70">
        <form class="flex flex-col m-auto w-fit h-fit p-2 bg-white rounded-md" action="">
            <p>Task Name:</p>
            <input v-model="taskName" class="border border-blue-300 rounded-md" type="text">
            <div class="flex justify-center">
                <button @click.self="props.isModalOpen = false" class=" bg-blue-300 rounded-md w-fit pl-1 pr-1 m-2 hover:bg-blue-400 active:ring active:ring-blue-400 active:ring-offset-1" >Cancel</button>
                <button @click.self="PostTask" class=" bg-blue-300 rounded-md w-fit pl-1 pr-1 m-2 hover:bg-blue-400 active:ring active:ring-blue-400 active:ring-offset-1">Create Task</button>
            </div>
        </form>
    </div>
</template>