<template>
  <div class="browser--button" v-if="!isSpacer">
    <img src="" :alt="title"/>
    <span>{{title}}</span>
  </div>
  <div class="browser--spacer" v-else>

  </div>
</template>

<script>
export default {
  name: "BrowserButton",
  props: {
    title: { type: String },
    isSpacer: { type: Boolean },
  }
}
</script>

<style lang="scss" scoped>
.browser--button {
  background: lightgray;
  border-radius: 5px;

  width: 2em;
  height: 2em;
  margin: 0 .25em;
}

.browser--spacer {
  width: 1em;
  padding: 1em;
}
</style>
