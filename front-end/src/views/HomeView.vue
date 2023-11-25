<template>
  <Navbar name="Arief Commerce" />
  <HeaderVue display="Tempat Topup Game Termurah" fw="Arief Store" />
  <SectionDashboard />
  <FooterDashboard />
  <router-view />
</template>
<script>
import Navbar from '../components/partials/NavbarDashboard.vue'
import HeaderVue from '../components/partials/HeaderDashboard.vue'
import SectionDashboard from '../components/SectionDashboard.vue'
import FooterDashboard from '../components/partials/FooterDashboard.vue'
export default {
  name: 'HomeView',
  components: {
    Navbar,
    HeaderVue,
    SectionDashboard,
    FooterDashboard
  }
}
</script>
