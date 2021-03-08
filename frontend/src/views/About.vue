<template>
  <div id="content" class="centeredContent">
    <div class="centeredBox">
      <p>{{ $t('about.first_paragraph') }}</p>
      <p>{{ $t('about.second_paragraph')}}</p>
      <p>{{ $t('about.come_take_look')}} <a href="https://scinna.app/">{{ $t('about.website') }}</a> {{ $t('about.or') }} <a href="https://github.com/scinna">{{ $t('about.github') }}</a>.</p>
    </div>
  </div>
</template>

<script>
/** @TODO: Check if the user is an admin, and if so show up the build number & date of compile under this paragraph **/
export default {
  name: "About"
}
</script>

<style lang="scss" scoped>
@import '../assets/Global.scss';
</style>
