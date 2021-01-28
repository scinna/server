<template>
  <div class="content">
    <p>Scinna is an open-source, self-hosted screenshot and picture sharing server designed to be used with Scinnapse, Scinna CLIent, and Scinnamon.</p>
    <p>The goal of this project is to provide a simple way to make it's own picture sharing server and to have a fully-featured client on all platforms.</p>
    <p>Come take a look at <a href="https://scinna.app/">our website</a> or <a href="https://github.com/scinna">our github</a>.</p>
  </div>
</template>

<script>
/** @TODO: Check if the user is an admin, and if so show up the build number & date of compile under this paragraph **/
export default {
  name: "About"
}
</script>

<style lang="scss" scoped>
@import '../assets/Colors.scss';

.content {
  width: 450px;
  padding: 2em;
  margin: auto;
  text-align: justify;

  background: $background-darker;
  box-shadow: 0 0 2px rgba(black, .5);
  border-radius: $border-radius;
  color: $text-lighter;

}

</style>
