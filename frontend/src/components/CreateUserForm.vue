<template>
  <form @submit.prevent="submitForm">
    <label>First Name</label>
    <input type="text" required v-model="first_name">
    <br>
    <label>Last Name</label>
    <input type="text" required v-model="last_name">
    <br>
    <label>Email</label>
    <input type="email" required v-model="email">
    <br>
    <label>Password</label>
    <input type="password" required v-model="password">
    <br>
    <label>Role</label>
    <select v-model="role">
    </select>
    <br>
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

      try {
        axios.post(api, formData, {})
      }catch (e){
        console.log(e)
      }

    }
  }
}
</script>

<style scoped>

</style>