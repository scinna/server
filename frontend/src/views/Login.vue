<template>
  <div id="content">
    <form @submit.prevent="login">
      <CustomInput type="text" label="Username" v-model="username" required/>
      <CustomInput type="password" label="Password" v-model="password" required/>

      <span class="message error" v-if="error.length > 0">{{error}}</span>

      <router-link class="message" to="/">Forgotten password?</router-link>

      <CustomInput type="submit" />
    </form>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import {Mutations} from "@/store/Mutations";
import CustomInput from "@/components/CustomInput.vue";
import {Authenticate} from "@/api/User";

export default Vue.extend({
  name: 'Login',
  components: {CustomInput},
  data: function() {
    return {
      username: '',
      password: '',
      error: '',
    }
  },
  methods: {
    login: function() {
      return Authenticate({ Username: this.username, Password: this.password})
      .then(resp => {
        this.$store.dispatch(Mutations.LOGIN_RESPONSE, resp.data);
      })
      .catch(err => {
        if (err.response && err.response.data && err.response.data.Message) {
          this.error = err.response.data.Message;
        } else {
          this.error = 'unknown';
        }
      });
    }
  }
});
</script>

<style lang="scss" scoped>
  @import '../assets/CenteredForm';
</style>
