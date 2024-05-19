<script setup>
defineProps({ data: Object })
const emit = defineEmits(['proceed'])
</script>

<template>
    <div @click="emit('proceed', data)"
        class="img aspect-video overflow-hidden bg-white bg-opacity-5  shadow relative">
        <img loading="lazy" v-shared-element:[data?.img] :src="data?.bigThumbnail" alt="thumbnail"
            class="absolute w-full inset-0">
        <h2 class="absolute bottom-1 right-1 backdrop-blur px-2 text-sm shadow">{{ data?.date || 'Date' }}</h2>
    </div>
</template>

<style scoped>
/* .img {
    min-width: 288px;
    max-width: 50%;
    flex-grow: 1;
} */
</style>