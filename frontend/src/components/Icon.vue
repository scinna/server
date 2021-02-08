<template>
  <div class="icon" @click="() => $emit('click')">
    <img src="../assets/tmp_directory.png" :alt="Name" v-if="IsDirectory"/>
    <!-- @TODO: Handle picture thumbnail + private picture thumbnails-->
    <!-- @TODO: save the thumbnail server-side so it does not render it everytime -->
    <img src="../assets/tmp_directory.png" :alt="Name" v-else/>
    <span>{{Name}}</span>
  </div>
</template>

<script>
export default {
  name: "Icon",
  props: {
    Name: { type: String },
    IsDirectory: { type: Boolean },
    Visibility: { type: Number },
  }
}
</script>

<style lang="scss" scoped>
  .icon {
    display: flex;
    flex-direction: column;
    align-items: center;

    padding: .25em;

    img {
      width: 4em;
      height: 4em;
    }
  }
</style>
