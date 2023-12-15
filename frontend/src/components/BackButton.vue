<script setup>
import { inject } from 'vue';
import { useRouter } from 'vue-router';
import BackIcon from './icons/BackIcon.vue';

const router = useRouter()

const sidebarOpen = inject('sidebarOpen')

</script>

<template>
    <button v-wave @click="router.push({ name: 'home' })"
        :class="sidebarOpen ? 'top-14' : 'top-2'"
        class="absolute left-2 p-3 rounded-full duration-150 border border-white border-opacity-25 aspect-square backdrop-blur hover:bg-opacity-40 bg-white bg-opacity-10 z-50">
        <BackIcon class="fill-white opacity-75"/>
    </button>
</template>
