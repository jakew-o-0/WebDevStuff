<script setup>
import { ref } from 'vue';

const TaskContent = defineProps(['id', 'name', 'completed']);
const TaskId = ref(TaskContent.id);

</script>

<template>
    <div @click="$emit('UpdateTask', TaskId)" class="bg-slate-600 rounded-md shadow-lg shadow-black/70 w-fit max-w-screen-sm">
        <div class=" bg-blue-300 p-2 rounded-md">
            <h2 class=" font-semibold text-lg text-blue-400">ID</h2>
            <p>{{ TaskContent.id }}</p>
            <div class=" mt-3 mb-3">
                <h2 class=" font-semibold text-lg text-blue-400">Name</h2>
                <p>{{ TaskContent.name }}</p>
            </div>
        </div>
        <div class="flex justify-center">
            <h2 v-if="TaskContent.completed === true" class=" text-green-500">Completed</h2>
            <h2 v-else class=" text-red-500">Completed</h2>
        </div>
    </div>
</template>
