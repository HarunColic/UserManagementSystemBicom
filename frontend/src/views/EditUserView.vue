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

</style>