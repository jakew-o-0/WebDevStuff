<script setup>
</script>

<template>
    <div class="flex bg-gray-100 rounded shadow-md mx-auto w-fit">
        <div class="flex-row m-5 justify-start">
            <h1 class="font-Header font-semibold text-5xl flex justify-center">Task Name</h1>
            <div class="flex justify-around">
                <h2 class="font-Body text-lg">Class name</h2>
                <h2 class="font-Body text-lg">Subject</h2>
                <h2 class="font-Body text-lg">Due:</h2>
            </div>
            <div class="flex mt-5">
                <button class="bg-indigo-700 text-gray-100  flex items-center p-1 mx-2 my-2 rounded hover:bg-indigo-600 active:ring active:ring-indigo-600 font-Body text-xl">View Info</button>
                <button class="bg-emerald-300 flex items-center p-1 mx-2 my-2 rounded hover:bg-emerald-400 active:ring active:ring-emerald-400 font-Body text-xl">Submit Work</button>
            </div>
        </div>

        <div class="flex-row justify-center rounded mx-5 my-auto">
            <h1 class="font-Header font-semibold text-4xl text-indigo-700">Class Score: </h1>
            <h1 class="font-Header font-semibold text-xl text-indigo-700">Current Multiplier: </h1>
        </div>
    </div>
</template>
