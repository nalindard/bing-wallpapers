<script setup>
import { ref, onMounted, watch, inject } from 'vue';
import { useRouter } from 'vue-router';
import { get, getJson, getSrcLink } from '../utils';
import BingImage from '../components/BingImage.vue';
import { store as time } from '../store/index.js';

const router = useRouter()

const jsonData = ref([])

onMounted(async () => {
    if (!window.updated) {
        updateTimeStamp()
        window.updated = true
    }
    getImages()
})

watch(time, () => getImages())

function updateTimeStamp() {
    const d = new Date()
    time.thisYear = d.getFullYear()
    time.thismonth = (d.getMonth() + 1)
    time.year = time.thisYear
    time.month = time.thismonth
    time.monthTitle = time.monthsList[time.thismonth - 1]?.title
}

async function getImages() {
    jsonData.value = getJson(await get(getSrcLink()))
}

function proceedBing(data) {
    console.log(data);
    router.push({ name: 'image', params: { id: data.date }, query: { ...data, type: "bing" } },)
}

const sidebarOpen = inject('sidebarOpen')

</script>

<template>
    <div class="w-full h-[98vh] overflow-hidden pl-4 pr-4 pt-3">
        <h2 :class="sidebarOpen ? 'pl-12' : ''" class="text-4xl pb-4 flex justify-between transition-all duration-150"
            v-shared-element:title>Monthly Wallpapers <span class="">{{ time.year }} / {{ time.monthTitle }}</span> </h2>
        <div class="overflow-y-scroll h-full pb-14 mb-20">
            <div class="w-full flex flex-wrap gap-2 justify-between items-center pb-4">
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