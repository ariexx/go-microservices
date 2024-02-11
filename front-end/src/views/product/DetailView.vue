<script>
import Navbar from '../../components/partials/NavbarDashboard.vue'
import FooterDashboard from '../../components/partials/FooterDashboard.vue'
import axios from 'axios'
import getImageBySlug from '../../helper/image'
import router from '@/router'
export default {
  name: 'DetailView',
  components: {
    Navbar,
    FooterDashboard
  },
  props: ['id', 'product-name'],
  mounted() {
    this.getProductDetailByProductId(this.id)
    this.getPaymentChannels()
    this.getProductById(this.id)
  },
  data() {
    return {
      product: {},
      productData: {},
      paymentChannels: [],
      player_id: '',
      email: '',
      product_id: '',
      payment_id: '',
      price: '',

      getImageBySlug: getImageBySlug
    }
  },
  methods: {
    async getProductDetailByProductId(id) {
      try {
        const res = await fetch(`http://localhost:8080/api/v1/products/${id}`)
        const data = await res.json()
        this.product = data.data
      } catch (err) {
        console.log(err.message)
      }
    },

    async getProductById(id) {
      try {
        const res = await fetch(`http://localhost:8080/api/v1/product/${id}`)
        const data = await res.json()
        this.productData = data.data
      } catch (err) {
        console.log(err.message)
      }
    },

    async getPaymentChannels() {
      try {
        const res = await fetch('http://localhost:8080/api/v1/payments')
        const data = await res.json()
        this.paymentChannels = data.data
      } catch (err) {
        console.log(err.message)
      }
    },

    createOrder() {
      try {
        this.getPriceByProductDetailId(this.product_id)
        axios
          .post('http://localhost:8080/api/v1/orders', {
            email: this.email,
            product_id: String(this.product_id),
            price: this.price,
            total: this.price,
            player_id: this.player_id,
            payment_id: this.payment_id
          })
          .then((response) => {
            alert('Order berhasil dibuat')
            router.push({ name: 'payment', params: { orderId: response.data.data.order.order_id } })
          })
          .catch((error) => {
            console.log(error)
            return alert(error)
          })
      } catch (err) {
        console.log(err.message)
      }
    },

    getPriceByProductDetailId(id) {
      this.product.forEach((element) => {
        if (element.id === id) {
          this.price = element.price
          this.total = element.price
        }
      })
    }
  }
}
</script>
<template>
  <Navbar name="Arief Commerce" />
  <!-- Product section-->
  <section class="py-5">
    <div class="container px-4 px-lg-5 my-5">
      <div class="row gx-4 gx-lg-5 align-items-center">
        <div class="col-md-5">
          <img
            class="card-img-top mb-5 mb-md-0"
            :src="getImageBySlug(productData.name)"
            :alt="productData.name"
          />
        </div>
        <div class="col-md-7">
          <form @submit.prevent="createOrder()">
            <div class="card mb-3">
              <div class="card-body">
                <div class="form-row">
                  <label for="basic-url">Player ID</label>
                  <div class="col-sm-6 col-md-6 col-lg-12">
                    <input
                      type="text"
                      class="form-control"
                      placeholder="Player ID Game"
                      v-model="player_id"
                      required
                    />
                  </div>
                </div>
              </div>
            </div>
            <div class="card mb-3">
              <div class="card-body">
                <label>Pilih Paket :</label>
                <div class="form-row">
                  <div class="col-sm-6 col-md-6 col-lg-12">
                    <select class="form-control" v-model="product_id" required>
                      <option
                        v-for="detailProduct in product"
                        :value="detailProduct.id"
                        :key="detailProduct.id"
                      >
                        {{ detailProduct.name }} -
                        {{
                          Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(
                            detailProduct.price
                          )
                        }}
                      </option>
                    </select>
                  </div>
                </div>
              </div>
            </div>
            <div class="card mb-3">
              <div class="card-body">
                <label>Pilih Channel Pembayaran :</label>
                <div class="form-row">
                  <div class="col-sm-6 col-md-6 col-lg-12">
                    <select class="form-control" v-model="payment_id">
                      <option
                        v-for="paymentChannel in paymentChannels"
                        :value="paymentChannel.id"
                        :key="paymentChannel.id"
                      >
                        {{ paymentChannel.name }}
                      </option>
                    </select>
                  </div>
                </div>
              </div>
            </div>
            <div class="card mb-3">
              <div class="card-body">
                <div class="form-row">
                  <label for="basic-url">Email</label>
                  <div class="col-sm-6 col-md-6 col-lg-12">
                    <input
                      type="email"
                      class="form-control"
                      placeholder="Email"
                      v-model="email"
                      required
                    />
                  </div>
                </div>
              </div>
            </div>
            <button class="btn btn-primary btn-lg">Beli</button>
          </form>
        </div>
      </div>
    </div>
  </section>
  <FooterDashboard />
</template>

<style></style>
