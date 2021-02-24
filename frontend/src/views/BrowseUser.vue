<template>
  <Browser :username="$route.params.username"/>
</template>

<script>
import Browser from "@/components/Browser.vue";

export default {
  name: "BrowseUser",
  components: {Browser}
}
</script>

<style scoped>

</style>
