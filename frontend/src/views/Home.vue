<template>
  <div id="content">
    <Browser v-if="user !== null" :username="user.Name" />
    <Loader v-else/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Browser from "@/components/Browser.vue";
import {mapState} from "vuex";
import Loader from "@/components/Loader.vue";

export default Vue.extend({
  name: 'Home',
  components: {Loader, Browser},
  computed: {
    ...mapState({
      user: state => state.Account.User,
    })
  },
});
</script>

<style lang="scss" scoped>
#content {
  display: flex;
  flex-direction: column;
  justify-content: stretch;
}
</style>
