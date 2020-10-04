<template>
  <div id="app" class="login">
    <nav>
      <img src="../assets/logo.png" alt="Scinna Logo"/>
    </nav>
    <div id="content">
      <form v-if="IsInRegistrationMode" @submit="register">
        <label for="username">Username:</label>
        <input id="username" type="text" v-model="Username"/>

        <label for="email">Email:</label>
        <input id="email" type="email" v-model="Email" />

        <label for="password">Password:</label>
        <input id="password" type="password" v-model="Password" />

        <label for="password2">Repeat:</label>
        <input id="password2" type="password" v-model="Password2" />

        <label for="invite" v-if="!RegistrationAllowed">Invite code: </label>
        <input id="invite" type="text" v-model="InviteCode" v-if="!RegistrationAllowed"/>

        <a href="#" @click="changeMode">Already have an account?</a>

        <input type="submit" :disabled="!canRegister" value="Register"/>
      </form>
      <form v-else @submit="login">
        <label for="username">Username:</label>
        <input id="username" type="text" v-model="Username"/>
        <label for="password">Password:</label>
        <input id="password" type="password" v-model="Password" />

        <span>{{ ErrorMessage }}</span>
        <a href="#" @click="changeMode">Not registered yet?</a>

        <input type="submit" :disabled="!canLogin" value="Login"/>
      </form>
    </div>
  </div>
</template>

<script>
import {mapState} from "vuex";

export default {
  name: 'Login',
  data: function() {
    return {
      Username: localStorage.getItem('scinna-last-username') || '',
      Email: '',
      Password: '',
      Password2: '',
      InviteCode: '',
      IsInRegistrationMode: false,
    }
  },
  computed: {
    ...mapState({
      RegistrationAllowed: state => state.Server.RegistrationAllowed,
      ErrorMessage: state => state.ErrorMessage,
    }),
    canLogin() {
      return this.Username.length > 0 && this.Password.length > 0
    },
    canRegister() {
      return this.canLogin && this.Email.length > 0 && this.Password2 === this.Password && (this.RegistrationAllowed || this.InviteCode.length > 0);
    }
  },
  methods: {
    login(e) {
      e.preventDefault();
      this.$store.dispatch('login', {router: this.$router, data: { Username: this.Username, Password: this.Password, }});
      return false;
    },

    register(e) {
      e.preventDefault();
      this.$store.dispatch('register', { Username: this.Username, Password: this.Password, Email: this.Email, InviteCode: this.InviteCode })
      return false;
    },

    changeMode() {
      this.Username = this.Password = this.Password2 = this.InviteCode = this.Email = "";
      this.IsInRegistrationMode = !this.IsInRegistrationMode;
      this.$store.commit('clearError');
    },
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/style';

  .login {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    nav img {
      height: 3em;
    }

    #content {
      justify-content: center;
      align-items: center;
    }

    form {
      @include card;

      display: flex;
      flex-direction: column;
      align-items: center;
      min-width: 400px;

      a {
        display: block;
        color: var(--secondary);
        margin: 1em;
      }

      label, input {
        display: block;
        margin: .5em;
      }

      button {
        margin-top: 1em;
      }

      span {
        color: var(--error);
      }
    }
  }
</style>