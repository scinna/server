<template>
  <div class="content">
    <p>{{ $t('about.first_paragraph') }}</p>
    <p>{{ $t('about.second_paragraph')}}</p>
    <p>{{ $t('about.come_take_look')}} <a href="https://scinna.app/">{{ $t('about.website') }}</a> {{ $t('about.or') }} <a href="https://github.com/scinna">{{ $t('about.github') }}</a>.</p>
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
