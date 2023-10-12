<script setup>
import axios from 'axios';
import { ref } from 'vue';

const emit = defineEmits(['CloseModal']);
const taskName = ref("");
const errorMessage = ref("");

function PostTask() {
    axios.post(
        'http://127.0.0.1:8000/task/',
        {
            name: taskName.value,
            completed: false,
        },
    )
    .then((res) => {
        console.log(res);
        if(res.status === 201) {
            emit('CloseModal')
        }
        else {
            errorMessage.value = res.statusText;
        }

    })
    .catch((err) => {
        console.error(err);
    });
}
</script>

<template>
    <div @click.self="$emit('CloseModal')" class="absolute top-0 left-0 grid h-full w-full bg-slate-500/70 z-50">
        <form class="flex flex-col m-auto w-fit h-fit p-2 bg-white rounded-md" @submit.prevent="" action="#">
            <p>Task Name:</p>
            <input v-model="taskName" class="border border-blue-300 rounded-md" type="text">
            <div class="flex justify-center">
                <button type="button" @click.self="$emit('CloseModal')" class=" bg-blue-300 rounded-md w-fit pl-1 pr-1 m-2 hover:bg-blue-400 active:ring active:ring-blue-400 active:ring-offset-1" >Cancel</button>
                <button type="button" @click.self="PostTask" class=" bg-blue-300 rounded-md w-fit pl-1 pr-1 m-2 hover:bg-blue-400 active:ring active:ring-blue-400 active:ring-offset-1">Create Task</button>
            </div>
        </form>
    </div>
</template>