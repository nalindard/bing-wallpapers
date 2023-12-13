<script setup>
defineProps({ data: Object })
const emit = defineEmits(['proceed'])
</script>

<template>
    <div @click="emit('proceed', data)" class="img aspect-video rounded overflow-hidden bg-white bg-opacity-5 border border-white border-opacity-40 shadow-lg relative">
        <img loading="lazy" v-shared-element:[data?.img] :src="data?.bigThumbnail" alt="thumbnail" class="absolute inset-0">
        <h2 class="absolute bottom-1 right-1 backdrop-blur px-2 text-sm shadow">{{ data?.date || 'Date' }}</h2>
        <!-- <p>{{ data || 'Error' }}</p> -->
    </div>
</template>

<style scoped>
.img{
    min-width: 288px;
    max-width: 33%;
    flex-grow: 1;
}
</style>


<!-- {
    "date": "2021-02-28",
    "img": "https://cn.bing.com/th?id=OHR.TwinsDenning_EN-US9910127756_UHD.jpg",
    "smallThumbnail": "https://cn.bing.com/th?id=OHR.TwinsDenning_EN-US9910127756_UHD.jpg&pid=hp&w=50",
    "bigThumbnail": "https://cn.bing.com/th?id=OHR.TwinsDenning_EN-US9910127756_UHD.jpg&pid=hp&w=384&h=216&rs=1&c=4" 
} -->