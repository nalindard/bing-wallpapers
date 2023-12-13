<script setup>
import { ref } from 'vue';

const props = defineProps({
    title: String,
    showTitle: Boolean,
})

const main = ref()
const bar = ref()

const x_s = () =>
{
    main.value.style.fontSize = '6rem'
    bar.value.style.scale = '100%'
}
const x_o = () =>
{
    main.value.style.fontSize = '1.76rem'
    bar.value.style.scale = '112%'
}
const x_e = () =>
{
    main.value.style.fontSize = '2.25rem'
    bar.value.style.scale = '100%'
}
</script>

<template>
    <div class="w-full h-full overflow-y-hidden overflow-x-hidden pr-72">
        <h2 v-show="showTitle" ref="main" class="text-4xl pb-4 pl-12 transition-all duration-700 text-base-100">
            {{ title || 'undefined !' }}
        </h2>
        <div v-dragscroll ref="bar" v-on:dragscrollstart="x_s" v-on:dragscrollmove="x_o" v-on:dragscrollend="x_e"
            class="w-screen py-2 pl-2 bg-base-300 bg-opacity-40 overflow-hidden cursor-move flex justify-start items-center  transition-all duration-700 scroll">

            <slot>
                <div class="aspect-video h-32 bg-base-200 bg-opacity-25 mx-2 cursor-auto rounded"
                    v-for="x in Array.from(Array(8).keys())"></div>
            </slot>

        </div>
    </div>
</template>

<style scoped>
.scroll {
    background-image: linear-gradient(to right, rgba(255, 255, 255, 0.277), rgba(0, 0, 0, 0) 10%, rgba(0, 0, 0, 0) 90%, rgba(255, 255, 255, 0.277)) !important;
}
</style>