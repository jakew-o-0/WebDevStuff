<script>
export default{
setup() {
    const TaskContent = defineProps(['id', 'name', 'completed'])

    function showTasks() {
        console.log(TaskContent)
    }

    return {
        TaskContent,
        showTasks
    }
},

created() {
    showTasks
}
}
</script>

<template>
    <div>
        <h2>ID</h2>
        <p>{{ TaskContent.id }}</p>
        <h2>Name</h2>
        <p>{{ TaskContent.name }}</p>
        <h2>Completed</h2>
        <p>{{ TaskContent.completed }}</p>
    </div>
</template>
