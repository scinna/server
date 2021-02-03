<template>
  <div id="content">
    <Loader/>

    <FloatingActionButton icon="plus" @click="() => {  }"/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Loader from "@/components/Loader.vue";
import FloatingActionButton from "@/components/FloatingActionButton.vue";

export default Vue.extend({
  name: 'Home',
  components: {FloatingActionButton, Loader },
});
</script>
