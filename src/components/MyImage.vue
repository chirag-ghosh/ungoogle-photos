<script setup lang="ts">
import { reactive, ref } from 'vue';


    type MetaData = {
        CreationTime: string, 
        Width: number,
        Height: number
    }

    defineProps<{link: string, metadata: MetaData, name: string}>()

    const isLoaded = ref<boolean>(false);

    function loadDone() {
      isLoaded.value = true;
    }
</script>

<template>
    <div class="image-box">
        <img :src="link" :alt="name" :onload="loadDone"/>
        <div v-if="isLoaded" class="time">{{new Date(metadata.CreationTime).toLocaleDateString('en-US', {month: 'short', day: 'numeric', year: 'numeric'})}}</div>
    </div>
</template>

<style scoped>
  .image-box {
    display: inline-block;
    position: relative;
    min-width: 30%;
    width: 30%;
    margin: 0 calc(10% / 6);
  }

  @media (max-width: 600px) {
    .image-box {
        min-width: 90%;
        width: 90%;
        margin: 0 calc(10% / 2);
    }
  }

  @media (min-width: 601px) and (max-width: 1000px) {
    .image-box {
        min-width: 45%;
        width: 45%;
        margin: 0 calc(10% / 4);
    }
  }

  .image-box img {
    min-width: 100%;
    width: 100%;
    background: transparent url('../assets/skeleton.gif');
  }

  .image-box img::before {
    content: "";
    display: block;
    padding-top: 100%;
    float: left;
  }

  .time {
    position: absolute;
    bottom: 12%;
    left: 50%;
    transform: translateX(-50%);
    font-size: 1.3rem;
    font-weight: 600;
  }
</style>