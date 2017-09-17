<template>
  <form class="form-inline my-2 my-lg-0" @submit.prevent>
    <input class="form-control mr-sm-2" type="text" placeholder="Instagram nickname" v-model="nickname">
    <button class="btn btn-outline-success my-2 my-sm-0" type="submit" @click="search" :disabled="searching">Search</button>
  </form>
</template>

<script>
export default {
  props: {
    searching: Boolean
  },
  data() {
    return { nickname: '_schulzsebastian' }
  },
  methods: {
    search() { this.$emit('search', this.nickname) }
  }
}
</script>