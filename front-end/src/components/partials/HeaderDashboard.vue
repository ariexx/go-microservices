<script setup>
defineProps({
  display: {
    type: String,
    required: true
  },
  fw: {
    type: String,
    required: true
  }
})
</script>
<template>
  <header class="bg-dark py-5">
    <div class="container px-4 px-lg-5 my-5">
      <div class="text-center text-white">
        <h2 class="display-4 fw-bolder">
          {{ display ?? 'All In One Place To Topup Game Currency' }}
        </h2>
        <p class="lead fw-normal text-white-50 mb-0">{{ fw ?? 'Arief Commerce' }}</p>
      </div>
    </div>
  </header>
</template>
