<script lang="ts">
  import axios from 'axios'
  import { reactive } from 'vue'

  type ImgurImage = {
    id:            string;
    title:         string;
    description:   string;
    datetime:      number;
    type:          string;
    animated:      boolean;
    width:         number;
    height:        number;
    size:          number;
    views:         number;
    bandwidth:     number;
    vote:          null;
    favorite:      boolean;
    nsfw:          null;
    section:       null;
    account_url:   null;
    account_id:    null;
    is_ad:         boolean;
    in_most_viral: boolean;
    has_sound:     boolean;
    tags:          any[];
    ad_type:       number;
    ad_url:        string;
    edited:        string;
    in_gallery:    boolean;
    link:          string;
  }

  const images = reactive<ImgurImage[]>([])
  export default {
    setup() {
      return{
        images
      }
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
        .catch((err) => {
          console.log(err)
        });
    }
  }
</script>

<template>
  <div class="title">
    Chirag's Gallery
  </div>
  <div class="image-grid">
    <img v-for="image in images" :key="image.id" :src="image.link" :alt="image.title" />
  </div>
</template>

<style scoped>
  .title {
    font-size: 3rem;
    font-weight: 500;
  }

  .image-grid {
    width: 90%;
    margin: 1rem auto;
  }

  .image-grid img {
    width: 30%;
    margin: calc(10% / 6);
  }
</style>