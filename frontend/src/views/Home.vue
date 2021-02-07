<template>
  <div id="content">
    <Browser :username="username" />

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

export default Vue.extend({
  name: 'Home',
  components: {Browser, Uploader, FloatingActionButton},
  computed: {
    ...mapState({
      username: state => state.Account.User.Name,
    })
  },
  data: function() {
    return {
      IsUploaderVisible: false,
    }
  },
});
</script>
