<template>
  <button @click="$emit('click')">
    <font-awesome-icon :icon="icon"/>
  </button>
</template>

<script lang="ts">
import Vue, {PropType} from 'vue';

export default Vue.extend({
  name: "FloatingActionButton",
  props: {
    icon: { type: Function as PropType<string> },
    click: { type: Function }
  }
});
</script>

<style lang="scss" scoped>
@import "../assets/Colors.scss";

  button {
    position: absolute;
    right: 1em;
    bottom: 1em;
    width: 3em;
    height: 3em;

    border: 2px solid darken($accent-color, 10%);
    border-radius: 50%;

    color: $background-darker;
    background: $accent-color;
  }
</style>
