<script setup>
import { computed } from 'vue';
import { RouterLink, RouterView, useRoute, useRouter } from 'vue-router'

const route = useRoute();
const router = useRouter();

function GetHomePath() {
  if(route.path === '/') {
    router.push("/home")
  } else {
    router.push('/')
  }
}
</script>


<template>
  
  <div class=" sticky top-0 flex mb-10 max-w-full h-max border-b border-slate-500 shadow-md bg-white z-10">
    <h1 class="grow font-semibold text-5xl ml-2 mb-2">Todo App</h1>
    <div class="m-auto">
      <router-link to="/login"  class=" bg-blue-300 rounded-md p-1 hover:bg-blue-400 active:bg-blue-500 active:ring active:ring-blue-400 active:ring-offset-1 m-1">Login</router-link>
      <button @click.self="GetHomePath" class=" bg-blue-300 rounded-md p-1 hover:bg-blue-400 active:bg-blue-500 active:ring active:ring-blue-400 active:ring-offset-1 m-1">Home</button>
    </div>
  </div>
  <div class="z-0">
    <RouterView />
  </div>
</template>

<style scoped>
</style>
