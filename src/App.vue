<script setup lang="ts">
  import axios from 'axios'

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

  const images: ImgurImage[] = [];

  function fetchImages() {
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

  fetchImages()
</script>

<template>
  <div>
    Hello
  </div>
  <ul>
    <li v-for="image in images" :key="image.id">
      <img :src="image.link" :alt="image.title" />
    </li>
  </ul>
</template>