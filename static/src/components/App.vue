<template>
  <div>
    <nav class="navbar navbar-expand-md navbar-dark bg-dark fixed-top">
      <a class="navbar-brand" href="#">Instaloc</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarsExampleDefault" aria-controls="navbarsExampleDefault" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarsExampleDefault">
        <ul class="navbar-nav mr-auto"></ul>
        <form class="form-inline my-2 my-lg-0" @submit.prevent>
          <input class="form-control mr-sm-2" type="text" placeholder="Instagram nickname" aria-label="Search" v-model="nick">
          <button class="btn btn-outline-success my-2 my-sm-0" type="submit" @click="runWebsocket" :disabled="!button">Search</button>
        </form>
      </div>
    </nav>
    <div id="map" style="height:100%; width: 100%"></div>
  </div>
</template>

<script>
import ol from 'openlayers'
export default {
  data() {
    return {
      nick: '_schulzsebastian',
      button: true
    }
  },
  methods: {
    getData() {
      fetch('/data')
      .then(resp => resp.json())
      .then(data => {
        alert(data.data)
      })
    },
    runWebsocket() {
      const ws = new WebSocket("ws://localhost:5000/websocket")
      ws.onopen = () => {
        this.button = false
        if(this.source) this.source.clear()
        this.source = new ol.source.Vector()
        this.layer = new ol.layer.Vector({
          source: this.source,
        })
        this.map.addLayer(this.layer)
        ws.send(this.nick)
      }
      ws.onmessage = evt => {
        if(evt.data == "end") ws.close()
        else this.addPoint(JSON.parse(evt.data))
      }
      ws.onclose = () => {
        this.button = true
        this.map.getView().fit(this.source.getExtent(), this.map.getSize());
      }
    },
    addPoint(point) {
      let feature = new ol.Feature({
        geometry: new ol.geom.Point(ol.proj.transform([point.lng, point.lat], 'EPSG:4326', 'EPSG:3857'))
      })
      this.source.addFeature(feature)
    }
  },
  mounted() {
    this.map = new ol.Map({
      target: 'map',
      layers: [
        new ol.layer.Tile({
          source: new ol.source.XYZ({
            urls: [
              'http://mt0.google.com/vt/lyrs=m&x={x}&y={y}&z={z}',
              'http://mt1.google.com/vt/lyrs=m&x={x}&y={y}&z={z}',
              'http://mt2.google.com/vt/lyrs=m&x={x}&y={y}&z={z}',
              'http://mt3.google.com/vt/lyrs=m&x={x}&y={y}&z={z}',
            ]
          })
        })
      ],
      view: new ol.View({
        center: ol.proj.fromLonLat([16.9, 52.4]),
        zoom: 3
      })
    })
  }
}
</script>
