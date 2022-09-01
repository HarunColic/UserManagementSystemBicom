<template>
  <div>
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
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "EditUserView",
  data(){
    return{
      id: this.$route.params.id,
      first_name: '',
      last_name: '',
      email: '',
      password: '',
      role: ''
    }
  },
  components: {
  },
  methods:{
    submitForm() {
      const api = `http://localhost:8000/users/${this.id}`

      let formData = new FormData();

      formData.append('first_name', this.first_name);
      formData.append('last_name', this.last_name);
      formData.append('email', this.email);
      formData.append('password', this.password);
      formData.append('role', this.role);

      axios.put(api, formData, {})
    },
    fetchUser(){
      const api = `http://localhost:8000/users${this.id}`

      axios.get(api).then((response) => {
        this.first_name = response.data.first_name
        this.last_name = response.data.last_name
        this.email = response.data.email
        this.password = response.data.password
        this.role = response.data.role
      })
    }
  },
  beforeMount() {
    this.fetchUser()
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