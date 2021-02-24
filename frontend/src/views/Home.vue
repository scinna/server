<template>
  <div id="content">
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'Home',
  mounted() {
    if (this.$store.state.Account.User)
    this.$router.push({
      name: 'Browse user',
      params: {
        username: this.$store.state.Account.User.Name,
        collection: '',
      }
    })
  }
});
</script>

<style lang="scss" scoped>
#content {
  display: flex;
  flex-direction: column;
  justify-content: stretch;
}
</style>
