<template>
  <div id="app">
    <Navbar id="navbar"/>
    <router-view/>
  </div>
</template>

<script lang="ts">
import Navbar from '@/components/Navbar.vue';

export default {
  name: 'App',
  components: {
    Navbar,
  },
}

</script>

<style lang="scss">
* {
  box-sizing: border-box;
}

html {
  height: 100%;
}

body {
  width: 100%;
  height: 100%;

  margin: 0;
  padding: 0;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;

  background: #2c3e50;
  height: 100%;

  color: white;

  display: flex;
  flex-direction: column;

  #navbar {
    flex: 0 0 3em;
  }

  #content {
    flex: 1;
  }
}
</style>
