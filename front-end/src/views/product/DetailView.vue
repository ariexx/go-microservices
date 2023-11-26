<script>
import Navbar from '../../components/partials/NavbarDashboard.vue'
export default {
  name: 'DetailView',
  components: {
    Navbar
  },
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
<template>
  <Navbar name="Arief Commerce" />
  <!-- Product section-->
  <section class="py-5">
    <div class="container px-4 px-lg-5 my-5">
      <div class="row gx-4 gx-lg-5 align-items-center">
        <div class="col-md-5">
          <img
            class="card-img-top mb-5 mb-md-0"
            src="https://dummyimage.com/600x700/dee2e6/6c757d.jpg"
            alt="..."
          />
        </div>
        <div class="col-md-7">
          <div class="card mb-3">
            <div class="card-body">
              <div class="form-row">
                <label for="basic-url">Player ID</label>
                <div class="col-sm-6 col-md-6 col-lg-12">
                  <input type="text" class="form-control" placeholder="Player ID Game" />
                </div>
              </div>
            </div>
          </div>
          <div class="card mb-3">
            <div class="card-body">
              <label>Pilih Paket :</label>
              <div class="form-row">
                <div class="col-sm-6 col-md-6 col-lg-12">
                  <select class="form-control">
                    <option value="1">Paket 1</option>
                    <option value="2">Paket 2</option>
                    <option value="3">Paket 3</option>
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
                  <select class="form-control">
                    <option value="1">Mandiri</option>
                    <option value="2">BCA</option>
                    <option value="3">QRIS</option>
                  </select>
                </div>
              </div>
            </div>
          </div>
          <button class="btn btn-primary btn-lg" type="submit">Beli</button>
        </div>
      </div>
    </div>
  </section>
</template>

<style></style>
