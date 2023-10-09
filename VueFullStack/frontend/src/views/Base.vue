<script>
import axios from 'axios'
import { ref } from 'vue'
import ItemInstance from '../components/ItemInstance.vue'
import Modal from '../components/Modal.vue'

export default {
setup() {
    const Tasks = ref([{_id: "dummy id", name: "dummy name", completed: "dummy completed"}]);
    const ReqestErr = ref(false);
    const CreateTaskPressed = ref(false);

        function GetAllTasks() {
            axios.get('http://localhost:8000/task/')
            .then((res) => {
                if(res.data.status_code === undefined) {
                    ReqestErr.value = false;
                } else {
                    ReqestErr.value = true;
                }
                Tasks.value = res.data;
            })
            .catch((err) => {
                ReqestErr.value = true;
            });
        };

        function updateTask(id) {
            axios.put(
                'http://localhost:8000/task/' + id,
                {
                    completed: true,
                })
            .then((res) => {
                console.log(res);
            })
            .catch((err) => {
                console.error(err);
            });
        }

    return {
        Tasks,
        ReqestErr,
        CreateTaskPressed,
        GetAllTasks,
        updateTask,
    };
},





created() {
    this.GetAllTasks();
},

components: { ItemInstance, Modal }
}
</script>


<template>
    <Modal @Close-Modal="CreateTaskPressed = false" v-if="CreateTaskPressed === true"/>

    <main class=" flex justify-center">
        <div class="flex flex-col justify-center lg:max-w-3xl">

            <div class="flex justify-center">
                    <h1 class="text-6xl font-semibold m-10 ">ToDo App</h1>
            </div>
            <div class="flex justify-center">
                <button @click.self="GetAllTasks" class=" bg-blue-300 rounded-md p-1 hover:bg-blue-400 active:bg-blue-500 active:ring active:ring-blue-400 active:ring-offset-1 m-1 shadow-sm">Show all tasks</button>
                <button   class=" bg-blue-300 rounded-md p-1 hover:bg-blue-400 active:bg-blue-500 active:ring active:ring-blue-400 active:ring-offset-1 m-1">Get task by ID</button>
                <button @click.self="CreateTaskPressed = true" class=" bg-blue-300 rounded-md p-1 hover:bg-blue-400 active:bg-blue-500 active:ring active:ring-blue-400 active:ring-offset-1 m-1">Create Task</button>
            </div>

            <div v-if="ReqestErr === false"  class=" grid grid-flow-row grid-cols-2 justify-center m-10">
                <ItemInstance
                    @Update-Task="(id) => {updateTask(id);}"
                    v-for="task in Tasks"
                    :key="task._id"
                    :id="task._id" 
                    :name="task.name" 
                    :completed="task.completed"/>
            </div>
            <div class="flex justify-center" v-else>
                <h2 class=" text-xl text-slate-400 font-semibold m-5">No Tasks Found</h2>
            </div>
        </div>
    </main>
</template>


<style>
</style>