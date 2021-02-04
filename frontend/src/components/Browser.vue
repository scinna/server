<template>
  <div id="browser">
    Browsing {{ username }}
  </div>
</template>

<script>
export default {
  name: "Browser",
  props: {
    username: { type: String }
  }
}
</script>

<style lang="scss" scoped>

</style>
