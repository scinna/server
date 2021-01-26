<template>
  <div id="content">
    <form @submit.prevent="login">
      <label for="username">Username:</label>
      <input type="text" id="username" v-model="username"/>
      <label for="password">Password:</label>
      <input type="password" id="password" v-model="password"/>

      <input type="submit" />
    </form>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import {Mutations} from "@/store/Mutations";

export default Vue.extend({
  name: 'Login',
  data: function() {
    return {
      username: '',
      password: '',
    }
  },
  methods: {
    login: function() {
      this.$store.dispatch(Mutations.LOGIN_REQUEST, { username: this.username, password: this.password })
        .then((resp) => {
          console.log(resp);
        })
        .catch((err) => {
          console.log(err);
        });
    }
  }
});
</script>

<style lang="scss" scoped>
  @import '../assets/CenteredForm';
</style>
