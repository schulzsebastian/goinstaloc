<template>
  <div>
    <nav class="navbar navbar-expand-md navbar-dark bg-dark fixed-top">
      <a class="navbar-brand" href="/">Instaloc</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarsExampleDefault" aria-controls="navbarsExampleDefault" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarsExampleDefault">
        <ul class="navbar-nav mr-auto">
        </ul>
        <search-box
          :searching="searched"
          @search="searchProfile">
        </search-box>
      </div>
    </nav>
    <instagram-map>
    </instagram-map>
  </div>
</template>

<script>
import SearchBox from './SearchBox.vue'
import InstagramMap from './InstagramMap.vue'

export default {
  data() {
    return {
      searched: false
    }
  },
  components: {
    SearchBox,
    InstagramMap
  },
  methods: {
    searchProfile(nickname) {
      let url = location.host
      if(location.host == "localhost:5001") url = "localhost:5000"
      const ws = new WebSocket(`ws://${url}/websocket`)
      ws.onopen = () => {
        this.searched = true
        ws.send(nickname)
      }
      ws.onmessage = evt => {
        if(evt.data == "end") ws.close()
        else this.$store.commit('addPhoto', JSON.parse(evt.data))
      }
    }
  }
}
</script>
