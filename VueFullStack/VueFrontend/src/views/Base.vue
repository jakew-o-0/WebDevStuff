<script>
import axios from 'axios'
import { ref } from 'vue'
import ItemInstance from '../components/ItemInstance.vue'

export default {
setup() {
    const Tasks = ref([]);
    const isTasksUpdated = ref(false);
    function GetAllTasks() {
        axios.get('http://0.0.0.0:8000/task/')
            .then((res) => {
            Tasks.data = res.data;
            isTasksUpdated.data = true;
        })
            .catch((err) => {
            console.error(err);
        });
    }
    return {
        Tasks,
        isTasksUpdated,
        GetAllTasks
    };
},

components: { ItemInstance }
}
</script>


<template>
    <div class="flex justify-center">
            <h1 class="text-6xl font-semibold m-10">ToDo App</h1>
    </div>
    <div class="flex justify-center">
        <button @click.self="GetAllTasks" class=" bg-blue-300 rounded-md p-1 hover:bg-blue-400 active:bg-blue-500 active:ring active:ring-blue-400 active:ring-offset-1 m-1">Show all tasks</button>
        <button  class=" bg-blue-300 rounded-md p-1 hover:bg-blue-400 active:bg-blue-500 active:ring active:ring-blue-400 active:ring-offset-1 m-1">Get task by ID</button>
        <button class=" bg-blue-300 rounded-md p-1 hover:bg-blue-400 active:bg-blue-500 active:ring active:ring-blue-400 active:ring-offset-1 m-1">Create Task</button>
    </div>
        <div  class=" grid grid-flow-row auto-rows-max">
            <ItemInstance
                v-for="task in Tasks" 
                :key="task._id" :id="task._id" 
                :name="task.name" 
                :completed="task.completed"/>
        </div>
</template>


<style>
</style>