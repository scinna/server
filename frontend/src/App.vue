<template>
    <router-view id="app"/>
</template>

<script>

export default {
  name: 'App',
  mounted() {
    // Fetch server infos (Like the registration state, etc...)
    this.$store.dispatch('fetchServerConfig');

    // Check if the token is still valid
  }
}

</script>

<style lang="scss">
@import './assets/style.scss';

* { box-sizing: border-box; }

body {
  margin: 0;
  padding: 0;

  background: var(--dp0);
  color: darken(#fff, 10%);
}

#app {
  font-family: 'Montserrat';
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  font-size: 1.1em;
  text-align: center;

  position: absolute;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;

  display: flex;
  flex-direction: column;

  nav {
    background: var(--dp4);

    width: 100%;
    padding: 1em;
  }

  #content {
    flex: 1;
    display: flex;
    flex-direction: column;
  }

  input[type="text"], input[type="password"], input[type="email"] {
    background: none;
    border: none;
    border-bottom: 2px solid var(--primary);
    color: darken(#fff, 10%);

    &.error {
      border-bottom-color: var(--error);
    }
  }

  button, input[type="submit"] {
    background: var(--primary);
    border-radius: .5em;
    border: none;
    padding: .75em;
    width: 100px;

    font-weight: bold;

    &:hover {
      filter: brightness(85%);
      cursor: pointer;
    }

    &:disabled {
      filter: brightness(75%);
      cursor: default;
    }
  }
}
</style>
