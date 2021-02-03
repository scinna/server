<template>
  <nav :id="id" class="nav">
    <div v-if="menuOpened" id="global_background" :class="menuOpened ? 'shown' : ''" @click="hideMenu"/>

    <font-awesome-icon class="nav--button" icon="bars" @click="showMenu"/>

    <router-link to="/">
      <img class="nav--logo" src="../assets/logo.png" alt="Scinna logo"/>
    </router-link>

    <div :class="'nav--links ' + (menuOpened ? 'shown' : '')">
      <img class="nav--links-logo" src="../assets/logo.png" alt="Scinna logo"/>

      <router-link class="link" to="/" v-if="isLoggedIn">Profile</router-link>
      <router-link class="link" to="/login" v-if="!isLoggedIn">Login</router-link>
      <router-link class="link" to="/register" v-if="!isLoggedIn">Register</router-link>
      <router-link class="link" to="/about">About</router-link>
      <router-link class="link" to="/logout" v-if="isLoggedIn">Logout</router-link>
    </div>
  </nav>
</template>

<script lang="ts">
import Vue from 'vue';
import {mapGetters} from "vuex";

export default Vue.extend({
  name: 'Navbar',
  props: ['id'],
  data() {
    return {
      menuOpened: false,
    }
  },
  computed: {
    ...mapGetters(['isLoggedIn'])
  },
  methods: {
    showMenu: function () {
      this.menuOpened = true;
    },
    hideMenu: function () {
      this.menuOpened = false;
    },
  },
  mounted() {
    window.addEventListener('resize', () => {
      if (window.innerWidth > 600)
        this.menuOpened = false;
    });
  }
});
</script>

<style scoped lang="scss">
@import '../assets/Responsive.scss';
@import '../assets/Colors.scss';

.nav {
  background: $background-darker;
  color: $text-lighter;
  padding: .5em;

  display: flex;
  flex-direction: row;

  align-items: center;

  &--button {
    display: none;
    position: absolute;
    left: 1em;
    height: 100%;

    font-size: 2em;
    margin-right: 1em;

    cursor: pointer;
  }

  &--logo {
    margin: .25em .25em .25em 1.5em;
    height: 2em;
  }

  &--links {
    background: $background-darker;
    z-index: 999;
    display: flex;
    flex: 1;

    justify-content: right;
    align-items: center;

    &-logo {
      display: none;
    }

    .link {
      margin: 0 .5em;
      font-weight: bold;
      color: inherit;
      text-decoration: none;
    }
  }
}

@media (max-width: $size-s) {
  .nav {
    justify-content: center;

    &--button {
      display: block;
    }

    &--links {
      background: $background;
      position: absolute;
      flex-direction: column;

      height: 100%;
      width: 200px;

      top: 0;
      left: 0;
      transform: translateX(-100%);
      transition: transform .3s ease-in-out;

      &-logo {
        display: block;
        height: 2em;
        margin: 1em 0;
      }

      .link {
        margin: .75em 0;
        padding: .5em;
        width: 100%;
        text-align: center;
        background: $background-darker;

        &:nth-last-child {
          margin-top: auto;
        }
      }

      &.shown {
        transform: translateX(0);
      }
    }
  }
}

#global_background {
  display: none;
  position: absolute;

  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;

  background: rgba(black, .5);

  z-index: 998;

  &.shown {
    @media (max-width: $size-s) {
      display: block;
    }
  }
}
</style>
