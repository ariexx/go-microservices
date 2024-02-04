<script>
import getImageBySlug from '../helper/image'
export default {
  data() {
    return {
      getImageBySlug: getImageBySlug,
      products: []
    }
  },
  mounted() {
    this.getProducts()
  },
  methods: {
    async getProducts() {
      try {
        const res = await fetch('http://localhost:8080/api/v1/products')
        const data = await res.json()
        this.products = data.data
      } catch (err) {
        console.log(err.message)
      }
    }
  }
}
</script>
<template>
  <!-- Section-->
  <section class="py-5">
    <div class="container px-4 px-lg-5 mt-5">
      <div class="row gx-4 gx-lg-5 row-cols-2 row-cols-md-3 row-cols-xl-4 justify-content-center">
        <div class="col mb-5" v-for="product in products" :key="product.id">
          <div class="card h-100">
            <!-- Product image-->
            <img class="card-img-top" :src="getImageBySlug(product.name)" />
            <!-- Product details-->
            <div class="card-body p-4">
              <div class="text-center">
                <!-- Product name-->
                <h5 class="fw-bolder">{{ product.name }}</h5>
              </div>
            </div>
            <!-- Product actions-->
            <div class="card-footer p-4 pt-0 border-top-0 bg-transparent">
              <div class="text-center">
                <a class="btn btn-outline-dark mt-auto" :href="'/product/' + product.id"
                  >Lihat Selengkapnya</a
                >
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
