<template>
  <div id="app">
    <Navbar id="navbar"/>
    <router-view/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Navbar from '@/components/Navbar.vue';
import {ServerProps} from "@/store/Server";
import {Mutations} from "@/store/Mutations";

export default Vue.extend({
  name: 'App',
  components: {
    Navbar,
  },
  mounted() {
    this.axios.get('/api/infos')
        .then((resp) => {
          const serverInfos = resp.data as ServerProps;
          this.$store.commit(Mutations.GOT_SERVER_INFOS, serverInfos);
        })

    this.$store.dispatch(Mutations.LOAD_USER_TOKEN)
        .then(() => {
          console.log("ok")
        })
        .catch(() => {
          console.log("err");
        });
  }
});
</script>

<style lang="scss">
@import 'assets/Colors.scss';
@import 'assets/Global.scss';

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

  background: $background;
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
