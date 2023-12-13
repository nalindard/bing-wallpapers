<script setup>
import { onMounted, ref } from 'vue';
import HelloWorld from './components/HelloWorld.vue'
import { RouterLink, RouterView } from 'vue-router';
import MenuIcon from './components/icons/MenuIcon.vue';
import SideBar from './components/SideBar.vue';
import BackButton from './components/BackButton.vue';

const sideBar = ref(null)
const fullWidth = ref(!true)

function t() {
  sideBar.value.classList.toggle('-left-72')
  sideBar.value.classList.toggle('left-0')
  fullWidth.value = !fullWidth.value
}

</script>

<template>
  <div class="w-full min-h-screen flex relative overflow-hidden">
    <button v-wave class="fixed top-2 left-2 px-2 py-1 rounded duration-75 hover:bg-white hover:bg-opacity-10 z-50"
      @click="t">
      <MenuIcon class="fill-white opacity-75" />
    </button>

    <!-- <BackButton /> -->

    <div ref="sideBar" class="absolute top-0 left-0 flex transition-all duration-300 min-h-full">

      <!-- Side -->
      <div class="w-72">
        <!-- <RouterLink :to="{ name: 'home' }">Home</RouterLink>
        <RouterLink :to="{ name: 'image' }">Image</RouterLink>
        <RouterLink :to="{ name: 'about' }">About</RouterLink> -->
        <SideBar />
      </div>

      <!-- Page -->
      <div class="min-h-screen overflow-y-scroll duration-150 transition-all" :class="fullWidth ? 'f-w': 'f-z'">
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
}</style>
