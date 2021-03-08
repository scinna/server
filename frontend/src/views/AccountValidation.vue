<template>
  <div id="content" class="centeredContent">
    <Loader v-if="status === 'pending'"/>
    <div v-else class="centeredBox">
      <div v-if="status === 'success'">
        <p>You have successfully validated your account <span class="glow">{{username}}</span>.</p>
        <p>You can now log in!</p>
      </div>
      <div v-else>{{status}}</div>
    </div>
  </div>
</template>
<script lang="ts">
import Vue from 'vue';
import Loader from "@/components/Loader.vue";

export default Vue.extend({
  name: 'AccountValidation',
  components: {Loader},
  data() {
    return {
      status: "pending",
      username: "",
    }
  },
  mounted() {
    this.$http.get("/api/auth/register/" + this.$route.params.account)
        .then(resp => {
          this.status = 'success';
          this.username = resp.data.username;
        })
        .catch(err => {
          if (err.response.data.Message) {
            this.status = err.response.data.Message;
          } else {
            this.status = "Unknown error"
          }
        })
  }
});
</script>

<style lang="scss" scoped>
@import '../assets/Global';

.glow {
  text-shadow: 0 0 5px #5ddfd6;
}

.centeredContent {
  text-align: center;
}
</style>
