<script setup>
import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import BackButton from '../components/BackButton.vue';
import DownloadIcon from '../components/icons/DownloadIcon.vue'
import SetWallIcon from '../components/icons/SetWallIcon.vue'

import { SetWallpaper } from '../../wailsjs/go/main/App'

const route = useRoute()
const router = useRouter()
// router.back()
const downloadBtn = ref(null)

const type = ref("")

onMounted(() => {
    // let id = route.params.id
    type.value = route.query?.type || 'peapix'
    // content.value = type
    // searchFor(id, type)

    // console.log(route.query);
})



async function setWallpaper() {
    // alert('setting', route.query?.fullUrl)
    if (type.value === 'peapix') {
        await SetWallpaper(route.query?.fullUrl)
    } else if (type.value === 'bing') {
        await SetWallpaper(getResolution())
    } else {
        return
    }
}

function getResolution() {
    let w = screen.width
    let url = route.query?.bigThumbnail
    let src

    // alert(w)
    // alert(url)

    if (w < 1920) {
        src = url.replace('w=384&h=216', `w=1920&h=1080`)
        return src
    }
    if (w < 2560) {
        src = url.replace('w=384&h=216', `w=2560&h=1440`)
        return src
    }

    // alert(src)
    // alert(url)

    return route.query?.img
}

function downloadImg(size) {
    if (type.value === 'bing') {   
        downloadBtn.value?.setAttribute('href', getResolution())
        downloadBtn.value?.click()
    }
    if (type.value == 'peapix'){
        downloadBtn.value?.setAttribute('href', route.query?.img)
        downloadBtn.value?.click()
    }
}

</script>

<template>
    <div class="w-full h-full relative overflow-hidden">

        <BackButton />

        <!-- Bavkdrop -->
        <!-- <div class="absolute inset-0 bg-cover bg-center blur"
            :style="`background-image: url(${route.query?.thumbUrl || route.query?.bigThumbnai});`"></div> -->

        <div class="absolute inset-0 backdrop:blur-lg">
            <img v-if="type === 'peapix'" v-shared-element:[route.query?.title] :src="route.query?.thumbUrl" class="w-full"
                alt="uhdimage">
            <img v-if="type === 'bing'" v-shared-element:[route.query?.img] :src="route.query?.bigThumbnail" class="w-full"
                alt="uhdimage">
            <!--<img v-shared-element:something :src="route.query?.imageUrl" alt="uhdimage">-->

            <!-- Real Image -->
            <!-- <img src="" alt="image" class="w-full"> -->
        </div>

        <div class="z-50 absolute left-0 w-full bottom-0">

            <div class="w-full flex justify-between">
                <span
                    class="flex flex-col mb-2 px-2 bg-white bg-opacity-20 hover:bg-opacity-50 duration-150 transition-all">
                    <!-- <label for="qfhd">Full HD</label>
                    <input type="radio" name="quality" id="qfhd" checked>
                    <label for="qwqhd">2K</label>
                    <input type="radio" name="quality" id="qwqhd">
                    <label for="quhd">4K</label>
                    <input type="radio" name="quality" id="quhd"> -->
                    <h4 class="text-sm">{{ route.query?.title || 'Title loading' }}</h4>
                    <h4 class="text-sm">{{ route.query?.copyright || '© Copyright loading' }}</h4>
                    <h4 class="text-sm">{{ route.query?.date }}</h4>
                </span>

                <span class="flex py-4">
                    <!-- <link rel="stylesheet" href="" > -->
                    <a ref="downloadBtn" href="" target="_blank" download></a>
                    <button @click="setWallpaper()"
                        class="bg-white bg-opacity-20 hover:bg-opacity-50 duration-150 transition-all border text-xl flex justify-between items-center px-2 mr-2">
                        <SetWallIcon class="fill-white mr-2" /> Set as wallpaper
                    </button>
                    <button @click="downloadImg('fhd')"
                        class="bg-white bg-opacity-20 hover:bg-opacity-50 duration-150 transition-all border text-xl flex justify-between items-center mr-2 px-2">
                        <DownloadIcon class="fill-white mr-2" /> FHD
                    </button>
                    <button @click="downloadImg('wqhd')"
                        class="bg-white bg-opacity-20 hover:bg-opacity-50 duration-150 transition-all border text-xl flex justify-between items-center mr-2 px-2">
                        <DownloadIcon class="fill-white mr-2" /> WQHD
                    </button>
                    <button @click="downloadImg('uhd')"
                        class="bg-white bg-opacity-20 hover:bg-opacity-50 duration-150 transition-all border text-xl flex justify-between items-center mr-2 px-2">
                        <DownloadIcon class="fill-white mr-2" /> UHD
                    </button>
                    <!-- <button @click="getResolution()" class="bg-emerald-500 text-7xl">screen</button> -->
                </span>
            </div>


            <!-- <img :src="route.query?.imgsrc" alt="" v-shared-element:[route.params?.id]> -->
            <!-- {{ route.query }} -->
        </div>
        <!-- <div class="absolute w-full h-2/4 linear bottom-0 left-0"></div> -->
    </div>
</template>

<style scoped>
.linear {
    background-image: linear-gradient(to bottom, rgba(0, 0, 0, 0), rgb(0, 0, 0))
}
</style>


<!-- https://cn.bing.com/th?id=OHR.IceSailingBalaton_EN-US2751943390_UHD.jpg&pid=hp&w=1920&h=1080&rs=1&c=4 -->
<!-- https://cn.bing.com/th?id=OHR.IceSailingBalaton_EN-US2751943390_UHD.jpg&pid=hp&w=2560&h=1440&rs=1&c=4 -->