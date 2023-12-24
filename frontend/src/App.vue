<script setup>
import { onMounted, ref, provide } from 'vue';
import { RouterView } from 'vue-router';
import MenuIcon from './components/icons/MenuIcon.vue';
import SideBar from './components/SideBar.vue';

const sideBar = ref(null)
const fullWidth = ref(!true)

function t() {
  sideBar.value.classList.toggle('-left-72')
  sideBar.value.classList.toggle('left-0')
  fullWidth.value = !fullWidth.value
}

onMounted(() => document.addEventListener('contextmenu', event => event.preventDefault()))

provide('sidebarOpen', fullWidth)
</script>

<template>
  <div class="w-full min-h-screen flex relative overflow-hidden">
    <button
      :class="fullWidth ? 'duration-150 border aspect-square rounded-full backdrop-blur hover:bg-opacity-40 bg-white bg-opacity-10 p-3' : 'hover:bg-white px-2 py-1'"
      v-wave class="fixed top-2 left-2 border-white border-opacity-25 rounded duration-150 hover:bg-opacity-10 z-50" @click="t">
      <MenuIcon class="fill-white opacity-75" />
    </button>

    <div ref="sideBar" class="absolute top-0 left-0 flex transition-all duration-150 min-h-full">

      <!-- Sidebar -->
      <div class="w-72">
        <SideBar />
      </div>

      <!-- Page view -->
      <div class="min-h-screen overflow-y-scroll duration-150 transition-all" :class="fullWidth ? 'f-w' : 'f-z'">
        <RouterView />
      </div>
    </div>
  </div>
</template>

<style scoped>
::-webkit-scrollbar {
  display: none;
}

.f-w {
  width: 100vw;
}

.f-z {
  width: calc(100vw - 288px);
}
</style>
