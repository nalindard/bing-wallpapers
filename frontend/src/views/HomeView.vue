<script setup>
import { ref, onMounted, watch, inject } from 'vue';
import PeapixImage from '../components/PeapixImage.vue';
import { useRoute, useRouter } from 'vue-router';
import VerticalScroller from '../components/VerticalScroller.vue';
import { get, getJson, getSrcLink } from '../utils';
import BingImage from '../components/BingImage.vue';
import { store as time } from '../store/index.js';

const route = useRoute()
const router = useRouter()

const data = ref([])
const jsonData = ref([])

// async function f() {
//     if (!window.postlist) {
//         const res = await fetch('https://peapix.com/bing/feed?country=jp', { cache: 'force-cache' })
//         data.value = await res.json()

//         // let res = await fetch(posts.postsUrl)
//         // data.value = await res.json()
//         window.postlist = data.value
//     } else {
//         data.value = window.postlist
//     }
// }


onMounted(async () => {
    if (!window.updated) {
        updateTimeStamp()
        // let link = `https://raw.githubusercontent.com/niumoo/bing-wallpaper/main/docs/2021-02.html`
        window.updated = true
    }
    // f()
    getImages()
})

watch(time, () => getImages())

function updateTimeStamp() {
    const d = new Date()
    time.thisYear = d.getFullYear()
    time.thismonth = (d.getMonth() + 1)
    time.year = time.thisYear
    time.month = time.thismonth
}
async function getImages() {
    jsonData.value = getJson(await get(getSrcLink()))
}
function proceedPeapix(data) {
    // router.push({ name: 'image',})
    console.log(data);
    router.push({ name: 'image', params: { id: data.date }, query: { ...data, type: "peapix" } },)
}
function proceedBing(data) {
    // router.push({ name: 'image',})
    console.log(data);
    router.push({ name: 'image', params: { id: data.date }, query: { ...data, type: "bing" } },)
}

const sidebarOpen = inject('sidebarOpen')

</script>

<template>
    <div class="w-full h-screen overflow-hidden pl-4 pr-4 pt-3">
        <!-- <h2 v-shared-element:title class="text-left text-7xl">Home</h2> -->
        <!-- <p>{{ data }}</p> -->
        <!-- <keep-alive> -->
        <!-- <div class="w-full h-full overflow-y grid grid-cols-7 justify-center items-center">
                <PeapixImage v-for="img in data" :data="img" @proceed="proceed" />
            </div> -->

        <!--<div class="w-full h-72">
            <VerticalScroller :title="'Week of bing'" :showTitle="true">
                <PeapixImage v-for="img in data" :data="img" @proceed="proceedPeapix" />
            </VerticalScroller>
        </div> -->
        <!-- </keep-alive> -->

        <!--  -->
        <h2 :class="sidebarOpen ? 'pl-12' : ''" class="text-4xl pb-4 flex justify-between transition-all duration-150" v-shared-element:title>Monthly Wallpapers <span class="">{{ time.year }}-{{ time.getMonth() }}</span> </h2>
        <div class="overflow-y-scroll h-full pb-72">
            <div class="w-full overflow-y-scroll flex flex-wrap gap-2 justify-between items-center pb-4">
                <BingImage v-for="j in jsonData" :data="j" :key="j?.date" @proceed="proceedBing" />
            </div>
        </div>
    </div>
</template>

<style scoped>
::-webkit-scrollbar {
    visibility: hidden;
    display: none;
}
</style>