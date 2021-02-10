<template>
  <div class="icon" v-on:dblclick="() => $emit('click')">
    <!--
      @TODO: Fix when there are fewer icons they expand so much the icon is floating in the middle
             Maybe use a grid for better flow ?
      -->
    <font-awesome-icon :icon="getIconFromVisibility()" class="visibility" />
    <img src="../assets/tmp_directory.png" :alt="getTitle" v-if="isCollection()"/>
    <img :src="media.Thumbnail" :alt="getTitle" v-else/>
    <span>{{ getTitle() }}</span>
  </div>
</template>

<script lang="ts">
import {Collection} from "@/types/Collection";
import {Media} from "@/types/Media";

export default {
  name: "Icon",
  props: ['collection', 'media'], // @TODO: Find how to type the vars >:(
  methods: {
    isCollection() {
      return this.collection !== undefined && this.collection !== null;
    },
    getTitle() {
      if (this.isCollection()) {
        return this.collection.Title;
      }

      return this.media.Title;
    },
    getIconFromVisibility() {
      switch (this.isCollection() ? this.collection.Visibility : this.media.Visibility) {
        case 1:
          return 'eye-slash';
        case 2:
          return 'lock'
        default:
          return 'globe'

      }
    }
  }
}
</script>

<style lang="scss" scoped>
.icon {
  position: relative;

  display: flex;
  flex-direction: column;
  align-items: center;

  padding: .25em;

  flex: 1 1 0;

  img {
    width: 4em;
    height: 4em;
  }

  .visibility {
    position: absolute;
    top: .5em;
    right: .5em;

    width: 1.3rem;
    height: 1.3rem;
  }
}
</style>
