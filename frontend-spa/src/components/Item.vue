<template>
  <div class="helllo">
    ddd
    <h1>{{item.Name}}</h1>
    <h2>{{item.Description}}</h2>
    <h2>{{item.Amount}}円</h2>
    <router-link to="/">HOMEへ</router-link>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'ItemCard',
  data () {
    return {
      item: {}
    }
  },
  created () {
    axios.get(`http://localhost:8888/api/v1/items/${this.$route.params.id}`).then(res => {
      console.log(res.data)
      this.item = res.data
    })
  }
}
</script>

<style scoped>

</style>
