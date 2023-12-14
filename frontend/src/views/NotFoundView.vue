<script setup>
import { ref, onMounted } from 'vue';
const err = ref()

onMounted(() => err.value = window?.err)

</script>

<template>
    <div class="pl-20 py-2">
        <h2>You are not supposed to see this unless somthing gone wrong !</h2>
        <p>If there is an erro, it will be listed down below</p>
        <p>{{ err }}</p>
    </div>
</template>