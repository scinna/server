<template>
  <div id="browser" v-if="isLoaded">
    <div class="browser--bar">
      <BrowserButton title="Previous"/>
      <BrowserButton title="Next"/>
      <BrowserButton is-spacer/>
      <BrowserButton title="Next"/>

      <h1>{{ usernameLocal }}/{{!content.IsDefault ? content.Title : ''}}</h1>
    </div>
    <div class="browser--content">
      <Icon v-for="media in content.Medias" :key="media.MediaID" :is-directory="false" :name="media.Title" />
    </div>
  </div>
  <Loader v-else/>
</template>

<script>
import BrowserButton from "@/components/BrowserButton";
import Loader from "@/components/Loader";
import {FetchCollection} from "@/api/Collections";
import Icon from "@/components/Icon";

export default {
  name: "Browser",
  components: {Loader, BrowserButton, Icon},
  props: {
    username: { type: String },
  },
  data: function () {
    return {
      browserUsername: '',
      browserCollection: '',
      isLoaded: false,
      content: null
    }
  },
  mounted() {
    this.browserUsername = this.username.length > 0 ? this.username : this.$route.params.username;
    this.browserCollection = this.$route.params.collection ?? '';

    FetchCollection(this.browserUsername, this.browserCollection)
        .then(resp => {
          this.content = resp.data;
          this.isLoaded = true;
        })
        .catch(err => {
          console.log(err.response)
        })
  }
}
</script>

<style lang="scss" scoped>

.browser--bar {
  display: flex;
  flex-direction: row;

  padding: 1em;

  h1 {
    font-size: 1.2em;
  }
}

.browser--content {
  display: flex;
  flex-direction: row;
}

</style>
