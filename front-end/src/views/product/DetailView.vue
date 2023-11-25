<template><h1>Hello</h1></template>
<script>
export default {
  name: 'DetailView',
  props: ['id'],
  mounted() {
    console.log('id', this.id)
  },
  data() {
    return {
      product: {}
    }
  },
  methods: {
    async getProductById(id) {
      try {
        const res = await fetch(`http://localhost:8080/api/v1/products/${id}`)
        const data = await res.json()
        this.product = data.data
      } catch (err) {
        console.log(err.message)
      }
    }
  }
}
</script>

<style></style>
