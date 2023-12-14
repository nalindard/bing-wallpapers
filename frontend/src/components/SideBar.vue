<script setup>
import { onMounted, reactive, ref, computed } from 'vue';
import { RouterLink, useRouter } from 'vue-router';
import SelectBtn from './SelectBtn.vue';
import AccentPicker from '../components/AccentPicker.vue'
import { store as time } from '../store/index.js';
import SettingsIcon from './icons/SettingsIcon.vue'
import HomeIcon from './icons/HomeIcon.vue'
import ReloadIcon from './icons/ReloadIcon.vue'

// const time = reactive({
//     month: "02",
//     monthTitle: "01 January",
//     year: 2021,
//     firstYear: 2021,
//     firstmonth: 4,
//     thisYear: 2023,
//     thismonth: 12,
//     monthsList: [
//         { title: "01 January", value: "01", enabled: true },
//         { title: "02 February", value: "02", enabled: true },
//         { title: "03 March", value: "03", enabled: true },
//         { title: "04 April", value: "04", enabled: true },
//         { title: "05 May", value: "05", enabled: true },
//         { title: "06 June", value: "06", enabled: true },
//         { title: "07 July", value: "07", enabled: true },
//         { title: "08 August", value: "08", enabled: true },
//         { title: "09 September", value: "09", enabled: true },
//         { title: "10 October", value: "10", enabled: true },
//         { title: "11 November", value: "11", enabled: true },
//         { title: "12 December", value: "12", enabled: true },
//     ]
// })

const router = useRouter()

const validMonths = computed(() => {
    if (time.year === time.firstYear) {
        return time.monthsList.slice(time.firstmonth - 1)
    } else if (time.year === time.thisYear) {
        return time.monthsList.slice(0, time.thismonth)
    } else {
        return time.monthsList
    }
})

const years = ref([])

const f = ({ value, title }) => {
    time.month = value
    time.monthTitle = title
}
const ff = ({ value }) => {
    let year = parseInt(value)

    if (year == time.firstYear) {
        if (time.month < time.firstmonth) {
            time.month = time.firstmonth
        }
    } else if (year == time.thisYear) {
        if (time.month > time.thismonth) {
            time.month = time.thismonth
        }
    }

    time.year = parseInt(value)
    // time.ye = title
}


onMounted(() => {
    const d = new Date()
    time.thisYear = d.getFullYear()
    time.thismonth = (d.getMonth() + 1)
    for (let y = 2021; y <= time.thisYear; y++) {
        years.value.push({ title: y, value: y })
    }
})

function reload() {
    // alert('')
    router.push({name: 'home'}).then(
        window.location.reload())
}

</script>

<template>
    <div class="px-4 pt-12 flex flex-col items-start w-full h-screen bg-black bg-opacity-5">
        <!-- <RouterLink :to="{ name: 'home' }">Home</RouterLink> -->
        <!-- <RouterLink :to="{ name: 'image' }">Image</RouterLink> -->
        <!-- <RouterLink :to="{ name: 'about' }">About</RouterLink> -->

        <!-- Current value -->
        <div>
            <!-- <h2 class="text-3xl">{{ time.year }} - {{ time.monthTitle.toString().slice(3) }}</h2> -->
            <h2 class="text-3xl pb-2">{{ time.year }} - {{ time.getMonth() }}</h2>
        </div>

        <!-- <AccentPicker /> -->

        <!-- Year picker -->
        <h2 class="mt-2 text-xl">Year</h2>
        <hr class="bg-white border-none bg-opacity-20 h-[1px] w-full">
        <div class="flex flex-col items-start h-1/4 overflow-y-scroll overflow-x-hidden w-full pt-2">
            <SelectBtn @select="ff" v-for="y in years" :key="y.value" :data="y" :name="'year'" :active="time.year" />
        </div>

        <!-- Month picker -->
        <h2 class="text-xl">Month</h2>
        <hr class="bg-white border-none bg-opacity-20 h-[1px] w-full">
        <div class="flex flex-col h-3/4 overflow-y-scroll overflow-x-hidden w-full pt-2">
            <SelectBtn @select="f" v-for="m in validMonths" :key="m.value" :data="m" :name="'month'" :active="time.month" />
        </div>


        <hr class="bg-white border-none bg-opacity-20 h-[1px] w-full">
        <div class="flex py-2 gap-2">
            <RouterLink v-wave class="p-1 rounded-full" :to="{ name: 'home' }">
                <HomeIcon class="fill-white" />
            </RouterLink>
            <RouterLink v-wave class="p-1 rounded-full" :to="{ name: 'about' }">
                <SettingsIcon class="fill-white" />
            </RouterLink>
            <button v-wave @click="reload" class="p-1 rounded-full">
                <ReloadIcon class="fill-white" />
            </button>
        </div>
    </div>
</template>

<style scoped>::-webkit-scrollbar {
    visibility: hidden;
}</style>