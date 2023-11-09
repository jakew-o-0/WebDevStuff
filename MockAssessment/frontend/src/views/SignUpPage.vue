<script setup>
import { RouterLink, RouterView, useRoute, useRouter } from 'vue-router'
const router = useRouter();

function GoToLogin() {
    router.push('/login/');
}
</script>

<template>
    <h1 class="flex justify-center font-Header text-5xl mt-10 mb-2 text-stone-800">SignUp</h1>

    <!-- Login form stuff -->
    <form action="" method="post" class="rounded bg-gray-100 mx-auto w-fit shadow-md">

        <input type="text" id="username" placeholder="Username:" required class="m-5 mb-5 p-1 border-2 border-indigo-700 outline-0 rounded font-Body text-lg">
        <br>
        <input type="email" id="email" placeholder="Email: " required class="mx-5 mb-10 p-1 border-2 border-indigo-700 outline-0 rounded font-Body text-lg" >
        <br>

        <input type="password" id="password" placeholder="Password: " required class="mx-5 mb-5 p-1 border-2 border-indigo-700 outline-0 rounded font-Body text-lg">
        <br>
        <input type="password" id="ConfPassword" placeholder="Confirm Password: " required class="mx-5 mb-10 p-1 border-2 border-indigo-700 outline-0 rounded font-Body text-lg">
        <br>

        <div class="flex-row mx-5 mb-10">
            <h3 class="font-extrabold font-Header text-xl">Account Type:</h3>
            <input type="radio" id="Student" name="AccountType">
            <label for="Student">Student</label>
            <br>
            <input type="radio" id="Teacher" name="AccountType">
            <label for="Teacher">Teacher</label>
        </div>
        <input type="text" id="SchoolCode" placeholder="School Code:" required class="mx-5 mb-10 p-1 border-2 border-indigo-700 outline-0 rounded font-Body text-lg">
        <br>

        <div class="flex justify-center pb-5 px-5">
            <input type="submit" value="SignUp" class=" grow bg-emerald-300 text-stone-800 font-Body p-1 mr-1 rounded hover:bg-violet-600 active:ring active:ring-violet-600 "> <br>
            <input type="button" value="Login" @click.self="GoToLogin" class="bg-indigo-700 text-gray-100 font-Body p-1 ml-1 rounded hover:bg-indigo-600 active:ring active:ring-indigo-600 "> <br>
        </div>
    </form>

</template>
