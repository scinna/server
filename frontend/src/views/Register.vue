<template>
  <div id="content">
    <form @submit.prevent="register">
      <CustomInput type="text" label="Username" :model="Request.username" required/>
      <CustomInput type="email" label="Email" :model="Request.email" required/>
      <CustomInput type="password" label="Password" :model="Request.password" required/>
      <CustomInput type="password" label="Repeat password" :model="Request.password2" required/>

      <CustomInput type="text" v-if="!RegistrationAllowed" label="Invitation code" :model="Request.inviteCode" required/>

      <span class="message error" v-if="error.length > 0">{{ error }}</span>

      <CustomInput type="submit" />
    </form>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import CustomInput from "@/components/CustomInput.vue";
import {mapState} from "vuex";

export default Vue.extend({
  name: 'Register',
  components: { CustomInput },
  data() {
    return {
      Request: {
        username: '',
        email: '',
        password: '',
        password2: '',
        inviteCode: '',
      },
      error: '',
    }
  },
  computed: {
    ...mapState('Server', [ 'RegistrationAllowed' ])
  },
  methods: {
    register: () => {
      console.log("lol");
    }
  }
});
</script>

<style lang="scss" scoped>
  @import "../assets/CenteredForm.scss";
</style>
