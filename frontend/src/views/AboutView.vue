<script setup>
import { onMounted, ref } from 'vue';
import { get, getJson } from '../utils';

const html = ref([])
const jsonData = ref([])


onMounted(async () => {
     let link = 'https://raw.githubusercontent.com/niumoo/bing-wallpaper/main/docs/2021-02.html'
     // jsonData.value = getJson(await get(link))
})

// async function get() {
//      const res = await fetch('https://raw.githubusercontent.com/niumoo/bing-wallpaper/main/docs/2021-02.html', { cache: 'force-cache' })
//      const htmlData = await res.text()
//      return htmlData
// }
// function getJson(htmlString) {
//      const parser = new DOMParser()
//      const html_ = parser.parseFromString(htmlString, 'text/html')
//      const nodeList = html_.querySelectorAll('.w3-third');
//      // html.value = [...nodeList]
//      [...nodeList].forEach(node => jsonData.value.push(nodeToData(node)))
//      // alert(data)
// }

// function nodeToData(node) {
//      return {
//           date: node.querySelector('p').innerHTML?.slice(0, 10),
//           img: node.querySelector('a').getAttribute('href'),
//           smallThumbnail: node.querySelector('.smallImg').getAttribute('src'),
//           bigThumbnail: node.querySelector('.bigImg').getAttribute('src'),
//      }
// }
// router.back()

</script>

<template>
     <h2 v-shared-element:title>About</h2>
     <!-- <p>{{ html }}</p> -->
     <!-- <p v-for="j in html">{{ j.innerHTML }}</p> -->
     <p v-for="j in jsonData">{{ j }}</p>
</template>
