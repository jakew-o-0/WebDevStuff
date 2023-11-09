<script setup>
</script> 

<template>
    <!-- Task related buttons -->
    <div class="sticky top-24 mx-auto flex-row bg-gray-100 shadow-md rounded w-fit">
        <button class="bg-indigo-700 text-gray-100 text-center p-1 mx-auto my-2 rounded hover:bg-indigo-600 active:ring active:ring-indigo-600 font-Body text-xl">
            Search Task
        </button>
        <button class="bg-indigo-700 text-gray-100 text-center p-1 mx-2 my-2 rounded hover:bg-indigo-600 active:ring active:ring-indigo-600 font-Body text-xl">
            Delete Due Tasks
        </button>
        <button class="bg-indigo-700 text-gray-100 text-center p-1 mx-2 my-2 rounded hover:bg-indigo-600 active:ring active:ring-indigo-600 font-Body text-xl">
            Add Task
        </button>
    </div>

    <div class="my-10">

        <!-- needs to be changed to a component -->
        <div class="bg-gray-100 flex-row rounded shadow-md mx-auto w-80">

            <h1 class="font-Header font-semibold text-5xl flex justify-center">Task Name</h1>

            <div class="flex justify-around">
                <h2 class="font-Body text-lg">Class name</h2>
                <h2 class="font-Body text-lg">Subject</h2>
                <h2 class="font-Body text-lg">Due:</h2>
            </div>

            <div class="flex justify-center mt-5">
                <button class="grow bg-indigo-700 text-gray-100 text-center p-1 mx-2 my-2 rounded hover:bg-indigo-600 active:ring active:ring-indigo-600 font-Body text-xl">
                    View Info
                </button>
                <button class="bg-emerald-300 p-1 grow mx-2 my-2 rounded hover:bg-emerald-400 active:ring active:ring-emerald-400 font-Body text-xl">
                    Edit Task
                </button>
            </div>
        </div>

    </div>
</template>
