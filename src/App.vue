<script lang="ts">
  import axios from 'axios'
  import { reactive } from 'vue'
  import MyImage from './components/MyImage.vue'

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
        return {
            images
        };
    },
    mounted() {
        axios.get("https://api.imgur.com/3/album/lKLDuiF", {
            headers: {
                Authorization: "Client-ID ab142a02ee43a2c"
            }
        })
            .then((response) => {
            images.push(...response.data.data.images);
        })
            .catch((err) => {
            console.log(err);
        });
    },
    components: { MyImage }
}
</script>

<template>
  <div class="header">
    <div class="title">Chirag's Gallery</div>
    <a class="ig-link" href="https://www.instagram.com/unfocusedclicks/" target="_blank">@unfocusedclicks</a>
  </div>
  <div class="image-grid">
    <MyImage v-for="image in images" :key="image.id" :link="image.link" :name="image.title" :metadata="JSON.parse(image.description)"/>
  </div>
</template>

<style scoped>

  .header {
    margin: 2rem auto;
    width: 90%;
    display: flex;
    justify-content: space-between;
  }

  .title {
    font-size: 3rem;
    font-weight: bold;
  }

  .ig-link {
    font-size: 1.8rem;
    color: #444;
  }

  .image-grid {
    width: 90%;
    margin: 1rem auto;
  }
</style>