<template>
  <div id="content">
    <Loader/>

    <Uploader v-if="IsUploaderVisible" :hide="() => this.IsUploaderVisible = false"/>
    <FloatingActionButton icon="plus" @click="() => this.IsUploaderVisible = true"/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Loader from "@/components/Loader.vue";
import FloatingActionButton from "@/components/FloatingActionButton.vue";
import Uploader from "@/components/Uploader.vue";

export default Vue.extend({
  name: 'Home',
  components: {Uploader, FloatingActionButton, Loader },
  data: function() {
    return {
      IsUploaderVisible: false,
    }
  }
});
</script>
