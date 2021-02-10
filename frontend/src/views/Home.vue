<template>
  <div id="content">
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Browser from "@/components/Browser.vue";
import {mapState} from "vuex";
import Loader from "@/components/Loader.vue";

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
