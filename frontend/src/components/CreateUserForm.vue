<template>
  <form @submit="submitForm">
    <label>First Name</label>
    <input type="text" required v-model="first_name">

    <label>Last Name</label>
    <input type="text" required v-model="last_name">

    <label>Email</label>
    <input type="email" required v-model="email">

    <label>Password</label>
    <input type="password" required v-model="password">

    <label>Role</label>
    <select v-model="role" required>

    </select>
    <input type="submit" value="submit">
  </form>
</template>

<script>
export default {
  name: "CreateUserForm",

  data(){
    return{
      first_name: '',
      last_name: '',
      email: '',
      password: '',
      role: ''
    }
  },
  methods: {

    submitForm(){

    }
  }
}
</script>

<style scoped>

</style>