<template>
  <div :class="'browser--button ' + (this.disabled ? 'disabled' : '')" @click="() => { !this.disabled && this.$emit('click') }">
    <font-awesome-icon :icon="icon" />
    <!--<span>{{title}}</span>-->
  </div>
</template>

<script lang="ts">
import Vue, {PropType} from 'vue';

export default Vue.extend({
  name: "BrowserButton",
  props: {
    title: { type: Function as PropType<string> },
    icon: { type: Function as PropType<string> },
    disabled: { type: Function as PropType<boolean> }
  }
});
</script>

<style lang="scss" scoped>
@import "@/assets/Colors.scss";

.browser--button {
  $size: 3em;

  flex: 0 0 $size;

  background: $background-lighter;
  border-radius: 5px;

  color: $accent-color;

  width: $size;
  height: $size;
  margin: 0 .25em;

  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  &.disabled {
    background: $background;
  }

  &:not(.disabled):hover {
    background: lighten($background-lighter, 5%);
  }

  svg {
    height: #{$size * .65};
  }
}

.browser--spacer {
  width: 1em;
  padding: 1em;
}
</style>
