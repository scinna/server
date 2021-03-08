<template>
  <div id="content" class="centeredContent">
    <form @submit.prevent="register" class="centeredBox">
      <CustomInput type="text" label="Username" v-model="username" :required="true"/>
      <CustomInput type="email" label="Email" v-model="email" :required="true"/>
      <CustomInput type="password" label="Password" v-model="password" :required="true"/>
      <CustomInput type="password" label="Repeat password" v-model="password2" :required="true"/>

      <CustomInput type="text" v-if="!RegistrationAllowed" label="Invitation code" :model="inviteCode"
                   :required="true"/>

      <span class="message error" v-if="error.length > 0">{{ error }}</span>

      <CustomSubmit :disabled="status === 'pending'"/>
    </form>
  </div>
</template>

<script lang="ts">
import Vue          from 'vue';
import CustomInput  from "@/components/CustomInput.vue";
import CustomSubmit from "@/components/CustomSubmit.vue";
import {mapState}   from "vuex";

export default Vue.extend({
  name: 'Register',
  components: {CustomInput, CustomSubmit},
  data() {
    return {
      username: '',
      email: '',
      password: '',
      password2: '',
      inviteCode: '',
      status: 'none',
      error: '',
    }
  },
  computed: {
    ...mapState({
      RegistrationAllowed: (state: any) => state.Server.RegistrationAllowed,
    })
  },
  methods: {
    register: function () {
      this.status = 'pending';

      this.$http.post("/api/auth/register", {
        Username: this.username,
        Email: this.email,
        Password: this.password,
        InviteCode: this.inviteCode,
      })
          .then(resp => {
            //this.error = resp.data.Message;

            this.username = '';
            this.password = '';
            this.password2 = '';
            this.email = '';

            //this.error = "Succes";
            //this.status = 'success';
          })
          .catch(err => {
            this.error = err.response.data.Message;
            this.status = 'none';
            this.password = '';
            this.password2 = '';
          })
    }
  }
});
</script>

<style lang="scss" scoped>
@import "../assets/CenteredForm.scss";

</style>
