<script lang="ts">
  import axios from 'axios'
  const images: string[] = [];

  export default {
    data() {
      return images
    },
    mounted() {
      axios.get("https://api.imgur.com/3/album/lKLDuiF", {
        headers: {
          Authorization: "Client-ID ab142a02ee43a2c"
        }
      })
        .then((response) => {
          images.push(...response.data.data.images)
        })
        .catch((err) => console.log(err));
    }
  }
</script>

<template>
  <div>
    Hello
  </div>
  <ul>
    <li v-for="image in images" :key="image.id">
      <img src={{image.link}} alt={{image.title}} />
    </li>
  </ul>
</template>