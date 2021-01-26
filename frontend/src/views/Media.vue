<template>
  <div id="content">
    Displaying image {{ $route.params.id }}
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'Media',
  components: {},
});
</script>
