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
          :searching="searching"
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
      searching: false
    }
  },
  components: {
    SearchBox,
    InstagramMap
  },
  methods: {
    searchProfile(nickname) {
      this.$store.commit('clearPhotos')
      this.$store.commit('clearScanned')
      let url = location.host
      if(location.host == "localhost:5001") url = "localhost:5000"
      const ws = new WebSocket(`ws://${url}/websocket`)
      ws.onopen = () => {
        this.searching = true
        ws.send(nickname)
      }
      ws.onmessage = evt => {
        if(evt.data == "end") {
          ws.close()
          this.searching = false
        }
        else if(evt.data == "error") {
          ws.close()
          alert('User does not exists or profile is private')
          this.searching = false
        }
        else if(parseInt(evt.data)) this.$store.commit('addScanned', parseInt(evt.data))
        else this.$store.commit('addPhoto', JSON.parse(evt.data))
      }
    }
  }
}
</script>
