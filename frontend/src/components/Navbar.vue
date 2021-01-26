<template>
  <nav :id="id">
    <button>hamburger</button>
    <img class="nav--logo" src="" alt="Scinna logo"/>
    <div class="nav--links">
      <img class="nav--mobile-logo" src="" alt="Scinna logo"/>
      <router-link to="/">Home</router-link> |
      <router-link to="/">Profile</router-link> |
      <router-link to="/">Logout</router-link> |
    </div>
  </nav>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'Navbar',
  props: [ 'id' ],
});
</script>

<style scoped lang="scss">
nav {

}
</style>
