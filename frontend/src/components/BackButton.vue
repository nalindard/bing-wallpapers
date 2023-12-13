<script setup>
import { useRouter, useRoute } from 'vue-router';
import BackIcon from './icons/BackIcon.vue';

const router = useRouter()
</script>

<template>
    <button v-wave @click="router.push({ name: 'home' })"
        class="absolute top-12 left-2 px-2 py-1 rounded-full duration-150 border aspect-square bg-white bg-opacity-20 hover:bg-opacity-40 hover:bg-red-400 hover:border-red-400 z-50">
        <BackIcon class="fill-white opacity-75"/>
    </button>
</template>
