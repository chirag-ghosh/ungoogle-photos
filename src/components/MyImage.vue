<script setup lang="ts">

    type MetaData = {
        CreationTime: string, 
        Width: number,
        Height: number
    }

    defineProps<{link: string, metadata: MetaData, name: string}>()
</script>

<template>
    <div class="image-box">
        <img :src="link" :alt="name"/>
        <div class="time">{{new Date(metadata.CreationTime).toLocaleDateString('en-US', {month: 'short', day: 'numeric', year: 'numeric'})}}</div>
    </div>
</template>

<style scoped>
  .image-box {
    display: inline-block;
    position: relative;
    width: 30%;
    margin: 0 calc(10% / 6);
  }

  @media (max-width: 600px) {
    .image-box {
        width: 90%;
        margin: 0 calc(10% / 2);
    }
  }

  @media (min-width: 601px) and (max-width: 1000px) {
    .image-box {
        width: 45%;
        margin: 0 calc(10% / 4);
    }
  }

  .image-box img {
    width: 100%;
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