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
  async mounted() {
    await this.getOrderByOrderId(this.orderId)
    await this.getPayment(this.order.payment_id)
    await this.getProductById(this.order.product_id)
    this.isQrisPayment()
    this.toRupiah(this.order.price)
    this.isDataLoaded = true
  },
  data() {
    return {
      order: {},
      payment: {},
      product: {},
      isDataLoaded: false,
      isDataNotFound: false,
      isQris: false,
      formattedPrice: ''
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
        this.order = data.data.order
      } catch (err) {
        this.isDataNotFound = true
        console.log(err.message)
      }
    },

    isQrisPayment() {
      if (this.payment.name.includes('QRIS') || this.payment.name.includes('qris')) {
        this.isQris = true
      }
    },

    async getProductById(id) {
      try {
        const res = await fetch(`http://localhost:8080/api/v1/product/${id}`)
        const data = await res.json()
        this.product = data.data
      } catch (err) {
        console.log(err.message)
      }
    },

    toRupiah(price) {
      this.formattedPrice = new Intl.NumberFormat('id-ID', {
        style: 'currency',
        currency: 'IDR'
      }).format(price)
    }
  },
}
</script>

<template>
  <NavbarDashboard name="Arief Commerce" />

  <div v-if="isDataNotFound" class="container mt-3">
    <div class="alert alert-danger" role="alert">
      Data not found
    </div>
  </div>

  <div v-else-if="!isDataLoaded">
    <div class="container mb-3 mt-3">
      <div class="row">
        <div class="col-md-8 w-100">
          <h3>Detail Transaksi</h3>
          <div class="card">
            <div class="card-body">
              <p aria-hidden="true">
                <span class="placeholder col-6"></span>
              </p>

              <a href="#" tabindex="-1" class="btn btn-primary disabled placeholder col-4" aria-hidden="true"></a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div v-else>
    <div class="container mb-3 mt-3">
      <div class="row">
        <div class="col-md-4">
          <h3>Detail Pembayaran</h3>
          <div class="card">
            <div class="card-body">
             <div v-if="isQris">
              <img src="https://via.placeholder.com/150" alt="QRIS" class="img-thumbnail w-100">
              <p class="mt-3 text-center">Scan QRIS ini untuk melakukan pembayaran</p>
             </div>

              <div v-else>
                <p>Transfer ke rekening berikut:</p>
                <p>{{payment.description}}</p>
              </div>

            </div>
          </div>
        </div>
        <div class="col-md-8">
          <h3>Detail Transaksi</h3>
          <div class="card">
            <div class="card-body">
              <table class="table table-bordered">
                <tr>
                  <td>Order ID</td>
                  <td>{{order.order_id}}</td>
                </tr>
                <tr>
                  <td>Player Id</td>
                  <td>{{order.player_id}}</td>
                </tr>
                <tr>
                  <td>Product Detail</td>
                  <td>{{product.product_detail[0].name}}</td>
                </tr>
                <tr>
                  <td>Price</td>
                  <td>{{formattedPrice}}</td>
                </tr>
                <tr>
                  <td>Payment Method</td>
                  <td>{{payment.name}}</td>
                </tr>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>


  <FooterDashboard />
</template>

<style scoped>

</style>