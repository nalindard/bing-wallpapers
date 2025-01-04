<script setup>
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import BackButton from "../components/BackButton.vue";
import DownloadIcon from "../components/icons/DownloadIcon.vue";
import SetWallIcon from "../components/icons/SetWallIcon.vue";
import { SetWallpaper } from "../../wailsjs/go/main/App";

const route = useRoute();
const router = useRouter();
const downloadBtn = ref(null);
const setWallBtn = ref(null);
const setWallBtnDisabled = ref(false);
const showOgImage = ref(false);

onMounted(() => {
  setTimeout(() => {
    showOgImage.value = true;
  }, 400);
});

async function setWallpaper() {
  try {
    setWallBtnDisabled.value = true;
    await SetWallpaper(getResolution());
  } catch (error) {
    window.err = error;
    router.push({ name: "404" });
  }
}

function getResolution() {
  let w = screen.width;
  let url = route.query?.bigThumbnail;
  let src;

  if (w < 1920) {
    src = url.replace("w=384&h=216", `w=1920&h=1080`);
    return src;
  }
  if (w < 2560) {
    src = url.replace("w=384&h=216", `w=2560&h=1440`);
    return src;
  }

  return route.query?.img;
}

function downloadImg(size) {
  try {
    let url = route.query?.bigThumbnail;
    let src;

    if (size == "fhd") {
      src = url.replace("w=384&h=216", `w=1920&h=1080`);
    } else if (size == "wqhd") {
      src = url.replace("w=384&h=216", `w=2560&h=1440`);
    } else {
      src = route.query?.img;
    }

    downloadBtn.value?.setAttribute("href", src);
    downloadBtn.value?.setAttribute("download", `${getTitle()}.jpg`);
    downloadBtn.value?.click();
  } catch (error) {
    window.err = error;
    router.push({ name: "404" });
  }
}

function getTitle(str = "") {
  let title = str?.replace("https://cn.bing.com/th?id=OHR.", "");
  title = title.slice(0, title.indexOf("_"));
  return title;
}

function camelCaseToWords(s) {
  const result = s.replace(/([A-Z])/g, " $1");
  return result.charAt(0).toUpperCase() + result.slice(1);
}
</script>

<template>
  <div class="w-full h-full relative rounded overflow-hidden">
    <BackButton />
    <div class="absolute inset-0 backdrop-blur-lg p-2 pt-0 overflow-hidden">
      <!-- Shared Element Image -->
      <img
        v-shared-element:[route.query?.img]
        :src="route.query?.bigThumbnail"
        class="absolute inset-0 w-full h-full object-cover"
        alt="uhdimage"
      />

      <!-- Real Image -->
      <img
        v-show="showOgImage"
        :src="getResolution()"
        alt="image"
        class="absolute inset-0 w-full h-full object-cover"
      />
    </div>

    <div class="z-50 absolute left-0 w-full bottom-0">
      <div class="w-full flex justify-between">
        <span
          class="flex flex-col ml-2 mb-2 px-2 py-1 border border-white border-opacity-25 bg-white bg-opacity-10 backdrop-blur duration-150 transition-all rounded"
        >
          <h4 class="text-lg">
            Title:
            {{
              camelCaseToWords(getTitle(route.query?.bigThumbnail)) ||
              "Title loading"
            }}
          </h4>
          <h4 class="text-base font-light">Date: {{ route.query?.date }}</h4>
          <h4 class="text-xs font-thin">
            {{ route.query?.copyright || "© Copyright loading" }}
          </h4>
        </span>

        <span class="flex items-end py-2">
          <a ref="downloadBtn" href="" target="_blank" download="bing.jpg"></a>
          <button
            ref="setWallBtn"
            :disabled="setWallBtnDisabled"
            @click="setWallpaper()"
            class="backdrop-blur hover:bg-opacity-25 bg-white bg-opacity-10 duration-150 transition-all border border-white border-opacity-25 flex justify-between items-center px-2 mr-2 text-sm font-light py-2 rounded"
          >
            <SetWallIcon class="fill-white mr-2" /> Set as wallpaper
          </button>
          <button
            @click="downloadImg('fhd')"
            class="backdrop-blur hover:bg-opacity-25 bg-white bg-opacity-10 duration-150 transition-all border border-white border-opacity-25 flex justify-between items-center mr-1 px-2 text-sm font-light py-2 rounded"
          >
            <DownloadIcon class="fill-white mr-2" /> Full HD
          </button>
          <button
            @click="downloadImg('wqhd')"
            class="backdrop-blur hover:bg-opacity-25 bg-white bg-opacity-10 duration-150 transition-all border border-white border-opacity-25 flex justify-between items-center mr-1 px-2 text-sm font-light py-2 rounded"
          >
            <DownloadIcon class="fill-white mr-2" /> 2K WQHD
          </button>
          <button
            @click="downloadImg('uhd')"
            class="backdrop-blur hover:bg-opacity-25 bg-white bg-opacity-10 duration-150 transition-all border border-white border-opacity-25 flex justify-between items-center mr-2 px-2 text-sm font-light py-2 rounded"
          >
            <DownloadIcon class="fill-white mr-2" /> 4K UHD
          </button>
        </span>
      </div>
    </div>
  </div>
</template>
