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
import axios from "axios";

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
  methods:{
    submitForm(){
      const api = 'http://localhost:8000/users'

      let formData = new FormData();

      formData.append('first_name', this.first_name);
      formData.append('last_name', this.last_name);
      formData.append('email', this.email);
      formData.append('password', this.password);
      formData.append('role', this.role);

      axios.post(api, formData, {})

    }
  }
}
</script>

<style scoped>

</style>