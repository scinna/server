<template>
  <div>
    <label :for="id" v-if="label && label.length > 0">{{ label }}: </label>
    <input :id="id" :type="type" :required="required" :value="value" v-on:input="updateValue($event.target.value)"/>
  </div>
</template>

<script>
export default {
  name   : "CustomInput",
  props  : {
    id      : {type: String},
    label   : {type: String},
    type    : {type: String},
    value   : {type: String},
    required: {type: Boolean},
  },
  methods: {
    updateValue: function (value) {
      this.$emit('input', value);
    }
  }
}
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
    border: none;
    padding: .25em;
    color: $text;
    border-radius: 5px;

    &[type="text"], &[type="password"] {
    }

    &[type="submit"] {
      padding: .75em;
    }
  }
}

</style>