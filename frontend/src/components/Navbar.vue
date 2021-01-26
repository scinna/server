<template>
  <nav :id="id" class="nav">
    <font-awesome-icon class="nav--button" icon="bars"/>
    <img class="nav--logo" src="" alt="Scinna logo"/>
    <div class="nav--links">
      <img class="nav--links-logo" src="" alt="Scinna logo"/>
      <router-link class="link" to="/">Home</router-link>
      <router-link class="link" to="/">Profile</router-link>
      <router-link class="link" to="/">Logout</router-link>
    </div>
  </nav>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'Navbar',
  props: ['id'],
});
</script>

<style scoped lang="scss">
@import '../assets/Responsive.scss';
@import '../assets/Colors.scss';

.nav {
  background: red;
  padding: .5em;

  display: flex;
  flex-direction: row;

  align-items: center;

  &--button {
    font-size: 2em;
  }

  &--logo {
    margin: .25em;
  }

  &--links {
    z-index: 999;
    display: flex;
    flex: 1;

    justify-content: right;
    align-items: center;
    @media (max-width: $size-s) {
      position: absolute;
      flex-direction: column;
      background: $background-lighter;

      height: 100%;
      width: 200px;

      top: 0;
      left: 0;
    }

    &-logo {
      display: none;

      @media (max-width: $size-s) {
        display: block;
      }
    }

    .link {
      margin: 0 .5em;
      font-weight: bold;
      color: inherit;
      text-decoration: none;
    }
  }
}
</style>
