<template>
  <input type="submit" :disabled="disabled" v-on:click.prevent="click"/>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'CustomSubmit',
  props: {
    disabled: {type: Boolean}
  },
  methods: {
    click: function (value: any) {
      this.$emit('click', value);
    }
  }
});
</script>

<style lang="scss" scoped>
@import "../assets/Colors.scss";

input[type="submit"] {
  border-radius: 5px;
  border: none;
  background: $accent-color;
  color: $background-darker;
  font-weight: bold;
  padding: .75em;

  &[disabled] {
    background: desaturate($accent-color, 15%);
    color: lighten($background-darker, 15%);
  }
}
</style>