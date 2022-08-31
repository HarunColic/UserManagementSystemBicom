<template>
  <div>
    <div v-for="user in users" :key="user.id" >
      <p>{{user.first_name}}</p><br>
      <p>{{user.last_name}}</p><br>
      <p>{{user.email}}</p><br>
      <router-link :to="{name: 'editUser', params: {id: user.id}}">Edit User</router-link> |

    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: "ListUsers.vue",
  data(){
    return {
      users: null
    }
  },
  created() {
    this.fetchData()
  },
  methods:{
    fetchData(){
      const api = 'http://localhost:8000/users';
      axios.get(api).then((response) => {
        //console.log(response.data)
        this.users = response.data
      })
    }
  }
}
</script>

<style scoped>

</style>