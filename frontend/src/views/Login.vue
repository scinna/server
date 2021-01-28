<template>
  <div id="content">
    <form @submit.prevent="login">
      <CustomInput type="text" label="Username" :model="username" required/>
      <CustomInput type="password" label="Password" :model="password" required/>

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
    }
  },
  methods: {
    login: function() {
      console.log(this.username, this.password)
      return
      /*
      Authenticate({ Username: this.username, Password: this.password})
      .then(resp => {
        this.$store.dispatch(Mutations.LOGIN_RESPONSE, resp.data);
      })
      .catch(err => {
        console.error("Err: ", err);
      });
       */
    }
  }
});
</script>

<style lang="scss" scoped>
  @import '../assets/CenteredForm';

  .forgot_password {

  }
</style>
