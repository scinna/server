<template>
  <div id="browser" v-if="isLoaded">
    <div class="browser--bar">
      <BrowserButton style="grid-area: b" title="Previous" icon="chevron-left" @click="previous" disabled />
      <BrowserButton style="grid-area: c" title="Next" icon="chevron-right" @click="next"/>

      <h1 style="grid-area: a">/{{ browserUsername }}/{{!content.IsDefault ? content.Title : ''}}</h1>
    </div>
    <div class="browser--content">
      <Icon v-for="collection in content.Collections" :key="collection.CollectionID + collection.Title + collection.Visibility" :collection="collection" />
      <Icon v-for="media in content.Medias" :key="media.MediaID + media.Title + media.Visibility" :media="media"/>
    </div>

    <Uploader v-if="IsUploaderVisible" :hide="() => this.IsUploaderVisible = false"/>
    <FloatingActionButton icon="plus" @click="() => this.IsUploaderVisible = true"/>
  </div>
  <Loader v-else/>
</template>

<script>
import BrowserButton from "@/components/BrowserButton";
import Loader from "@/components/Loader";
import {Browse as ApiBrowse} from "@/api/Collections";
import Icon from "@/components/Icon";
import Uploader from "@/components/Uploader";
import FloatingActionButton from "@/components/FloatingActionButton";

export default {
  name: "Browser",
  components: {Loader, BrowserButton, Icon, Uploader, FloatingActionButton},
  props: {
    username: { type: String },
  },
  data: function () {
    return {
      browserUsername: '',
      browserCollection: '',
      isLoaded: false,
      content: null,
      IsUploaderVisible: false,
    }
  },
  methods: {
    previous() {
      console.log("Previous");
    },
    next() {
      console.log("Next");
    }
  },
  mounted() {
    this.browserUsername = this.username.length > 0 ? this.username : this.$route.params.username;
    this.browserCollection = this.$route.params.collection ?? '';

    ApiBrowse(this.browserUsername, this.browserCollection)
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
@import "@/assets/Responsive.scss";
@import "@/assets/Colors.scss";

.browser--bar {
  display: grid;
  border-top: 1px solid $background-lighter;
  background: $background-darker;
  padding: 1em;
  gap: 1em;

  grid-template-areas: "a a" "b c";
  @media (min-width: $size-xs) {
    grid-template-areas: "b c a";
    justify-content: left;
  }

  align-items: center;
  justify-content: center;

  h1 {
    font-size: 1.2em;
    flex: 1;
    text-align: center;

    @media (min-width: $size-xs) {
      width: 100%;
    }
  }
}

.browser--content {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
}

</style>
