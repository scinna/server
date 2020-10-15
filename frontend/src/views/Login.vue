<template>
  <form @submit.prevent="login" v-if="!IsRegistration">
    <h1>Login</h1>
    <label for="username">Username</label>
    <input id="username" required v-model="username" type="text" />
    <label for="password">Password</label>
    <input id="password" required v-model="password" type="password" />
    <button type="submit">Login</button>
  </form>
  <form @submit.prevent="register" v-else>
    <h1>Register</h1>
    <label for="username">Username</label>
    <input id="username" required v-model="username" type="text" />
    <label for="password">Password</label>
    <input id="password" required v-model="password" type="password" />
    <label for="password2">Confirm password</label>
    <input id="password2" required v-model="confirm_password" type="password" />
    <button type="submit">Register</button>
  </form>
</template>

<script>
import {ACTION_LOGIN, ACTION_REGISTER} from "@/store/actions";

export default {
  name: 'Login',
  data: function () {
    return {
      IsRegistration: false,
      username: "",
      password: "",
      confirm_password: "",
      email: "",
      invite_code: ""
    }
  },
  methods: {
    login: function () {
      const payload = {
        Username: this.username,
        Password: this.password,
      };

      this.$store.dispatch(ACTION_LOGIN, payload).then(() => {
        this.$router.push('/')
      })
    },
    register: function() {
      const payload = {
        Username: this.username,
        Password: this.password,
        Email: this.email,
        InviteCode: this.invite_code,
      };

      this.$store.dispatch(ACTION_REGISTER, payload).then(() => {

      })
    }
  }
}
</script>
