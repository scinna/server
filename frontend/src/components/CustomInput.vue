<template>
  <div>
    <label :for="id" v-if="label && label.length > 0">{{ label }}: </label>
    <input v-if="!disabled" :id="id" :type="type" :required="required" :value="value" v-on:input="updateValue($event.target.value)" />
    <input v-else :id="id" :type="type" :required="required" :value="value" v-on:input="updateValue($event.target.value)" disabled />
    <span class="error" v-if="violations[id]">{{violations[id]}}</span>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name   : "CustomInput",
  props  : {
    id      : {type: String},
    label   : {type: String},
    type    : {type: String},
    value   : {type: String},
    required: {type: Boolean},
    disabled: {type: Boolean},
    violations: {type: Array},
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
