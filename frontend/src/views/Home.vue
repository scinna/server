<template>
  <div id="content">
    <Browser v-if="user !== null" :username="user.Name" />
    <Loader v-else/>

    <Uploader v-if="IsUploaderVisible" :hide="() => this.IsUploaderVisible = false"/>
    <FloatingActionButton icon="plus" @click="() => this.IsUploaderVisible = true"/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import FloatingActionButton from "@/components/FloatingActionButton.vue";
import Uploader from "@/components/Uploader.vue";
import Browser from "@/components/Browser.vue";
import {mapState} from "vuex";
import Loader from "@/components/Loader.vue";

export default Vue.extend({
  name: 'Home',
  components: {Loader, Browser, Uploader, FloatingActionButton},
  computed: {
    ...mapState({
      user: state => state.Account.User,
    })
  },
  data: function() {
    return {
      IsUploaderVisible: false,
    }
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
