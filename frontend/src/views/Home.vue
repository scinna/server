<template>
  <div id="content">
    <Loader/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Loader from "@/components/Loader.vue";

export default Vue.extend({
  name: 'Home',
  components: { Loader },
});
</script>
