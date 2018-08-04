<template>
  <div class="helllo">
    <h1>{{item.Name}}</h1>
    <h2>{{item.Description}}</h2>
    <h2>{{item.Amount}}円</h2>

    <payjp-checkout
      api-key="pk_test_e96d9571031577b42b6df1a0"
      text="カードを情報を入力して購入"
      submit-text="購入確定"
      name-placeholder="田中 太郎"
      v-on:created="onTokenCreated"
      v-on:failed="onTokenFailed">
    </payjp-checkout>

    <p>{{message}}</p>

    <router-link to="/">HOMEへ</router-link>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'ItemCard',
  data () {
    return {
      item: {},
      message: ''
    }
  },
  created () {
    axios.get(`http://localhost:8888/api/v1/items/${this.$route.params.id}`).then(res => {
      console.log(res.data)
      this.item = res.data
    })
  },
  beforeDestroy () {
    window.PayjpCheckout = null
  },
  methods: {
    onTokenCreated: function (res) {
      console.log(res.id)
      const data = { Token: res.id }
      axios.post(`http://localhost:8888/api/v1/charge/items/${this.$route.params.id}`, data)
      this.message = '商品の購入が完了しました'
    },

    onTokenFailed: function (e) {
      console.error(e)
    }
  }
}
</script>

<style scoped>

</style>
