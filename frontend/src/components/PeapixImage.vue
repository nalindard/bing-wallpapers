<script setup>
import { defineEmits, defineProps, ref } from 'vue'
import { sharedElementMixin } from 'v-shared-element'

const props = defineProps({
    data: Object
})

// mixins: [sharedElementMixin]

const emit = defineEmits(['proceed'])

// const $keepSharedElementAlive = ref(sharedElementMixin)

</script>

<template>
    <div @click="emit('proceed', data)" class="w-72 mr-2 relative overflow-hidden bg-white bg-opacity-5 border border-white border-opacity-40 shadow" style="min-width: 277px !important;">
        <!-- <img class="w-48" v-shared-element:[data?.title]="{ $keepSharedElementAlive }" :src="data?.thumbUrl" -->
        <img class="w-full aspect-video" v-shared-element:[data?.title] :src="data?.thumbUrl"
            :alt="data?.date">
        <!-- <h4>{{ data?.title }}</h4> -->
        <!-- <h4>{{ data?.copyright }}</h4> -->
        <h4 class="absolute bottom-1 right-1 backdrop-blur px-2 text-sm shadow">{{ data?.date }}</h4>
    </div>
</template>


<!-- { 
"title": "ポインセチア", 
"copyright": "© Yarygin/Getty Image", 
"fullUrl": "https://img.peapix.com/eefd4a29e93247f883df5dc488266959_1080.jpg", 
"thumbUrl": "https://img.peapix.com/eefd4a29e93247f883df5dc488266959_480.jpg", 
"imageUrl": "https://img.peapix.com/eefd4a29e93247f883df5dc488266959.jpg", 
"date": "2023-12-12", 
"pageUrl": "https://peapix.com/bing/45488" 
} -->