<script setup>
import { useRouter, useRoute } from "vue-router";
defineProps({
  data: Object,
  name: String,
  active: Number,
});
const emit = defineEmits(["select"]);
const router = useRouter();
const route = useRoute();

const select = (v, t) => {
  if (
    route.name === "about" ||
    route.name === "image" ||
    route.name !== "home"
  ) {
    router.push({ name: "home" }); // Navigate to home
  }
  emit("select", { value: v, title: t });
};
</script>

<template>
  <label
    @click="select(data?.value, data?.title)"
    class="font-thin text-sm py-1 min-h-[1.75rem] h-7 pl-4 border border-transparent hover:bg-white hover:bg-opacity-20 w-full text-left duration-150 transition-all relative overflow-hidden"
    v-wave
  >
    <input
      type="radio"
      :name="name"
      class="inset-0 absolute border-none bg-transparent w-full h-full pt-1"
      :checked="data?.value == active"
    />
    <span class="inset-0 absolute pl-4 h-full pt-1">
      {{ data?.title || "Error" }}
    </span>
  </label>
</template>

<style scoped>
input[type="radio"] {
  appearance: none;
  -webkit-appearance: none;
  transition: all 0.2s;
}
input[type="radio"]:checked::before {
  content: "";
  display: block;
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  background: rgba(255, 255, 255, 0.2);
}

@keyframes appear {
  0% {
    transform: scale(0);
  }
  100% {
    transform: scale(1);
  }
}
</style>
