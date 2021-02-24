<template>
  <div id="browser" v-if="isLoaded">
    <div class="browser--bar">
      <BrowserButton style="grid-area: b" title="Previous" icon="chevron-left" @click="previous" disabled/>
      <BrowserButton style="grid-area: c" title="Next" icon="chevron-right" @click="next" disabled/>
      <BrowserButton style="grid-area: d" :disabled="!user || username !== user.Name" title="Create folder" icon="folder-plus" @click="() => {}"/>

      <h1 style="grid-area: a">/{{ browserUsername }}/{{ !content.IsDefault ? content.Title : '' }}</h1>
    </div>
    <div class="browser--content">
      <Icon v-for="collection in content.Collections"
            :key="collection.CollectionID + collection.Title + collection.Visibility" :collection="collection"
            v-on:click="() => browse(collection)"/>
      <Icon v-for="media in content.Medias" :key="media.MediaID + media.Title + media.Visibility" :media="media"
            v-on:click="() => showMedia(media)"/>
    </div>

    <Uploader v-if="isUploaderVisible" :hide="() => this.isUploaderVisible = false"/>
    <FloatingActionButton icon="plus" @click="() => this.isUploaderVisible = true"/>
  </div>
  <Loader v-else/>
</template>

<script lang="ts">
import BrowserButton         from "@/components/BrowserButton.vue";
import Loader                from "@/components/Loader.vue";
import {Browse as ApiBrowse} from "@/api/Collections";
import Icon                  from "@/components/Icon.vue";
import Uploader              from "@/components/Uploader.vue";
import FloatingActionButton  from "@/components/FloatingActionButton.vue";
import {Collection}          from "@/types/Collection";
import {Media}               from "@/types/Media";
import {mapState}            from "vuex";
import Vue, {PropType}       from "vue";

type Data = {
  browserUsername: string;
  browserCollection: string;
  isLoaded: boolean;
  content: any;
  isUploaderVisible: boolean;
}

export default Vue.extend({
  name: "Browser",
  components: {Loader, BrowserButton, Icon, Uploader, FloatingActionButton},
  props: {
    username: {type: Function as PropType<string>},
  },
  computed: {
    ...mapState({
      user: state => state.Account.User,
    })
  },
  data: function (): Data {
    return {
      browserUsername: '',
      browserCollection: '',
      isLoaded: false,
      content: null,
      isUploaderVisible: false,
    }
  },
  methods: {
    previous() {
      console.log("Previous");
    },
    next() {
      console.log("Next");
    },
    browse(collection: Collection) {
      this.$router.push({ name: 'Browse user', params: { username: this.browserUsername, collection: collection.Title } })
    },
    showMedia(media: Media) {
      console.log("showing", media.Title)
    },
  },
  mounted() {
    this.browserUsername = this.username.length > 0 ? this.username : this.$route.params.username;

    ApiBrowse(this.browserUsername, this.$route.params.collection ?? '')
        .then(resp => {
          this.content = resp.data;
          this.isLoaded = true;
        })
        .catch(err => {
          console.log(err.response)
        })
  }
});
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

  grid-template-areas: "a a a" "b c d";
  @media (min-width: $size-xs) {
    grid-template-areas: "b c d a";
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
