<template>
  <div class="icon" @click="() => $emit('click')">
    <img src="../assets/tmp_directory.png" :alt="Collection.Title" v-if="Collection !== undefined && Collection !== null"/>
    <!-- @TODO: Handle picture thumbnail + private picture thumbnails-->
    <!-- @TODO: save the thumbnail server-side so it does not render it everytime -->
    <img :src="Media.Thumbnail" :alt="Media.Title" v-else/>
    <span>{{ Collection !== undefined && Collection !== null ? Collection.Title : Media.Title }}</span>
  </div>
</template>

<script>
import {Collection} from "@/types/Collection";
import {Media} from "@/types/Media";

export default {
  name: "Icon",
  props: ['Collection', 'Media'] // @TODO: Find how to type the vars >:(
}
</script>

<style lang="scss" scoped>
.icon {
  display: flex;
  flex-direction: column;
  align-items: center;

  padding: .25em;

  flex: 1 1 0;

  img {
    width: 4em;
    height: 4em;
  }
}
</style>
