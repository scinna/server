<template>
  <div>
    <label :for="id" v-if="label && label.length > 0">{{ label }}: </label>
    <input :id="id" :type="type" :required="required" :value="value" v-on:input="updateValue($event.target.value)"/>
  </div>
</template>

<script lang="ts">
import Vue, {PropType} from 'vue';

export default Vue.extend({
  name   : "CustomInput",
  props  : {
    id      : {type: Function as PropType<string>},
    label   : {type: Function as PropType<string>},
    type    : {type: Function as PropType<string>},
    value   : {type: Function as PropType<string>},
    required: {type: Function as PropType<boolean>},
  },
  methods: {
    updateValue: function (value: any) {
      this.$emit('input', value);
    }
  }
});
</script>

<style lang="scss" scoped>
@import '../assets/Colors.scss';

div {
  margin: .25em;
  padding: .25em;
  color: $text;

  display: flex;
  flex-direction: column;

  input {
    background: $background-lighter;
    padding: .25em;
    color: $text;
    border-radius: 5px;
    border: 2px solid $background;
    margin-top: .5em;

    &[type="text"], &[type="password"] {
    }

    &[type="submit"] {
      background: $accent-color;
      color: $background-darker;
      font-weight: bold;
      padding: .75em;

      &:disabled {
        background: darken($accent-color, 10%);
        color: lighten($background-darker, 10%);
      }
    }
  }
}

</style>
