<template>
  <div id="content">
    <!--
        @TODO: Create a custom form component that will use the loader during the request and auto-map the values in it
      -->
    <form @submit.prevent="login">
      <CustomInput type="text" :label="$t('login.username')" v-model="username" required/>
      <CustomInput type="password" :label="$t('login.password')" v-model="password" required/>

      <span class="message error" v-if="error.length > 0">{{ $t('login.errors.' + error) }}</span>

      <router-link class="message" to="/">{{ $t('login.forgotten_password') }}</router-link>

      <CustomInput type="submit" :value="$t('login.submit')" />
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
        this.$store.dispatch(Mutations.LOGIN_RESPONSE, resp.data).then(() => this.$router.push("/"))
      })
      .catch(err => {
        if (err.response && err.response.data && err.response.data.Message) {
          this.error = err.response.data.Message;
        } else {
          this.error = 'unknown';
        }

        this.password = '';
      });
    }
  }
});
</script>

<style lang="scss" scoped>
  @import '../assets/CenteredForm';
</style>
