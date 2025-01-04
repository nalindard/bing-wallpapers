<script setup></script>

<template>
  <div class="app-container px-20 py-4 font-light overflow-y-auto">
    <!-- <h1 class="text-4xl font-medium">About</h1> -->
    <!-- <hr class="my-6" /> -->

    <h2 class="text-2xl font-semibold">Disclaimer</h2>
    <p class="py-4 text-justify text-sm">
      The images featured in this software are sourced daily from Microsoft Bing
      and are provided by Microsoft solely for wallpaper use. All images and
      photos displayed are the property of their respective creators and
      copyright holders. This application does not claim ownership of any
      content.
    </p>
    <p class="py-4 text-justify text-sm">
      If you are the rightful owner of any image displayed here and have
      concerns regarding its usage, please contact me. I will promptly remove
      the content upon request.
    </p>
    <p class="py-4 text-justify text-sm">
      This project is a personal endeavor, created out of my admiration for
      these wallpapers. As Bing now limits access to wallpapers from the past
      seven days, this application aims to preserve and enjoy them.
    </p>
    <p class="py-4 text-justify text-sm">
      The source code for this software is publicly available
      <a
        class="text-blue-400 font-semibold"
        target="_blank"
        href="https://github.com/nalindard/bing-wallpapers"
        >here</a
      >. Please note, I am not liable for any damages or issues that may arise
      from using this software. Use it responsibly and at your own risk.
    </p>
    <hr class="my-6" />

    <h2 class="text-2xl font-semibold">Data Source and Usage Notice</h2>
    <p class="py-4 text-justify text-sm">
      This application loads data from raw GitHub URLs. The primary intent of
      this project is for learning and experimentation purposes. The data is
      fetched directly from GitHub repositories using raw URLs, which is
      different from traditional web scraping methods.
    </p>
    <p class="py-4 text-justify text-sm">
      If you are the owner of any repository or file linked within this
      application and have concerns about its usage or visibility, please
      contact me. I am happy to address any issues, including removing the data
      or links if needed. This project remains one of my favorite learning
      endeavors, and I aim to respect the rights and wishes of all original
      content creators.
    </p>
    <p class="py-4 text-justify text-sm">
      Thank you for your understanding and support.
    </p>
    <hr class="my-6" />

    <h2 class="text-2xl font-semibold mt-7">Built With</h2>
    <p class="text-sm flex flex-wrap gap-2 mt-4">
      <a
        class="text-sm lowercase text-blue-400 first-letter:uppercase"
        target="_blank"
        href="https://vuejs.org/"
        >Vue.js</a
      >
      <a
        class="text-sm lowercase text-blue-400 first-letter:uppercase"
        target="_blank"
        href="https://wails.io/"
        >Wails</a
      >
      <a
        class="text-sm lowercase text-blue-400 first-letter:uppercase"
        target="_blank"
        href="https://pkg.go.dev/github.com/reujab/wallpaper"
        >Native APIs</a
      >
      <a
        class="text-sm lowercase text-blue-400 first-letter:uppercase"
        target="_blank"
        href="https://github.com/niumoo/bing-wallpaper"
        >Bing Data</a
      >
      <a
        class="text-sm lowercase text-blue-400 first-letter:uppercase"
        target="_blank"
        href="https://tailwindcss.com/"
        >Tailwind CSS</a
      >
      <a
        class="text-sm lowercase text-blue-400 first-letter:uppercase"
        target="_blank"
        href="https://go.dev/"
        >Golang</a
      >
    </p>
    <hr class="my-6" />

    <p class="text-sm">
      Connect with me on
      <a
        class="text-blue-400 first-letter:uppercase font-semibold"
        href="https://github.com/nalindard"
        target="_blank"
        >GitHub</a
      >.
    </p>
    <hr class="my-6" />
  </div>
</template>

<style scoped>
.app-container {
  height: 100vh; /* Make the container take up the full viewport height */
  overflow-y: auto; /* Enable vertical scrolling */
  box-sizing: border-box; /* Include padding in height calculations */
}

body,
html {
  margin: 0;
  padding: 0;
  height: 100%; /* Ensure the body and html take up the full height */
  overflow: hidden; /* Prevent scrolling on the body */
}
</style>
