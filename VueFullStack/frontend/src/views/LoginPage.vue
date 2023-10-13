<script>
</script>

<template>
<main class=" absolute top-0 left-0 grid h-screen w-screen justify-items-center bg-slate-300">

    <div class="flex-col justify-center w-fit h-fit m-auto rounded-md bg-white shadow-xl">
        <h1 class=" bg-blue-400 rounded-md mb-2 text-center text-lg font-semibold">Login</h1>

        <form @submit.prevent class="m-2">
            <p>User Name:</p>
            <input class="border border-blue-300 rounded-md" type="text">
            <p>Password:</p>
            <input class="border border-blue-300 rounded-md" type="password">
            <div class="flex justify-center">
                <button class=" bg-blue-300 rounded-md w-fit pl-1 pr-1 mt-1 ml-1 mr-1 hover:bg-blue-400 active:ring active:ring-blue-400 active:ring-offset-1">Login</button>
                <button class=" bg-blue-300 rounded-md w-fit pl-1 pr-1 mt-1 ml-1 mr-1 hover:bg-blue-400 active:ring active:ring-blue-400 active:ring-offset-1">SignUp</button>
            </div>
        </form>
    </div>

</main>
</template>