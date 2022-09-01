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
      <option v-for="r in roles" :key="r.id" >
        {{r.name}}
      </option>
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
      role: '',
      roles: null
    }
  },
  beforeMount() {
    this.fetchRoles()
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

    },

    fetchRoles(){
      const api = 'http://localhost:8000/roles'

      axios.get(api).then((response) => {
        this.roles = response.data
      })
    }
  }
}
</script>

<style scoped>
input, select {
  width: 40%;
  padding: 12px 20px;
  margin: 8px 0;
  display: inline-block;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}

input[type=submit] {
  width: 10%;
  background-color: #4CAF50;
  color: white;
  padding: 14px 20px;
  margin: 8px 0;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

input[type=submit]:hover {
  background-color: #45a049;
}

div {
  border-radius: 5px;
  background-color: #f2f2f2;
  padding: 20px;
}
</style>
