<script>
import NavbarDashboard from '@/components/partials/NavbarDashboard.vue'
import FooterDashboard from '@/components/partials/FooterDashboard.vue'

export default {
  name: 'PaymentView',
  components: {
    NavbarDashboard,
    FooterDashboard
  },
  props: ['order-id'],
  mounted() {
    this.getOrderByOrderId(this.$route.params.id)
    // this.getPayment(this.$route.params.id)
  },
  data() {
    return {

    }
  },
  methods: {
    async getPayment(id) {
      try {
        const res = await fetch(`http://localhost:8080/api/v1/payment/${id}`)
        const data = await res.json()
        this.payment = data.data
      } catch (err) {
        console.log(err.message)
      }
    },

    //get order by id
    async getOrderByOrderId(id) {
      try {
        const res = await fetch(`http://localhost:8080/api/v1/order/${id}`)
        const data = await res.json()
        console.log(data.data)
        this.order = data.data
      } catch (err) {
        console.log(err.message)
      }
    },
  },
}
</script>

<template>
  <NavbarDashboard name="Arief Commerce" />

  <div class="container">
    <div class="row">
      <div class="col-md-8">
        <h1>Payment</h1>
        <div class="card">
          <div class="card-body">
<!--      Order Status and Payment      -->
          </div>
        </div>
      </div>
    </div>
  </div>

  <FooterDashboard />
</template>

<style scoped>

</style>