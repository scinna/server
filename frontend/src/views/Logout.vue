<template>
  <div class="content">
    <p>You successfully logged out of your scinna account !</p>
    <p>Redirecting on login page...</p>
  </div>
</template>

<script>
import {Mutations} from "@/store/Mutations";

export default {
  name: "Logout",
  mounted() {
    this.$store.dispatch(Mutations.LOGOUT).then(() => setTimeout(() => this.$router.push('/login'), 2000))
  }
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
