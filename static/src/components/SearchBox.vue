<template>
  <form class="form-inline my-2 my-lg-0" @submit.prevent>
    <input class="form-control mr-sm-2" type="text" placeholder="Instagram nickname" v-model="nickname">
    <button class="btn btn-outline-success" type="button" @click="search" :disabled="searching" v-text="buttonText"></button>
  </form>
</template>

<script>
export default {
  props: {
    searching: Boolean
  },
  data() {
    return {
      nickname: '_schulzsebastian',
      buttonText: 'Search'  
    }
  },
  methods: {
    search() { this.$emit('search', this.nickname) }
  },
  computed: {
    scanned() {
      return this.$store.getters.scanned
    }
  },
  watch: {
    searching(after, before) {
      if(before == true) this.buttonText = 'Search'
    },
    scanned(after, before) {
      if(after[0]) this.buttonText = `Scanned ${this.scanned.length}/${after[0]} photos...`
    }
  }
}
</script>