<template>
  <div id="instagram-map" style="height:99.5%; width: 100%">
    <div class="progress" v-show="loading">
      <div id="pb" class="progress-bar bg-danger" role="progressbar" style="width: 0%; height: 0.5%;"></div>
    </div>
  </div>
</template>

<script>
import ol from 'openlayers'
import 'openlayers/css/ol.css'

export default {
  data() {
    return {
      loading: true
    }
  },
  methods: {
    addPoint(point) {
      let feature = new ol.Feature({
        geometry: new ol.geom.Point(ol.proj.transform([point.lng, point.lat], 'EPSG:4326', 'EPSG:3857'))
      })
      this.source.addFeature(feature)
    },
    fitToLayer() {
      this.map.getView().fit(this.source.getExtent(), this.map.getSize())
    },
    clearMap() {
      this.source.clear()
    }
  },
  computed: {
    photos() {
      // Slice to prevent the same old and new value
      return this.$store.getters.photos.slice()
    }
  },
  watch: {
    photos(after, before) {
      if(after.length > before.length) {
        let toAdd = after.filter(p => {
          return before.map(b => b.id).indexOf(p.id) < 0
        })[0]
        this.addPoint(toAdd)
        document.getElementById('pb').style.width = `${Math.round(this.photos.length/toAdd.total*100)}%`
        if(this.photos.length == toAdd.total) this.fitToLayer()
      }
      //TODO: prograssbar
    }
  },
  mounted() {
    this.map = new ol.Map({
      target: 'instagram-map',
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
    this.source = new ol.source.Vector()
    this.layer = new ol.layer.Vector({
        source: this.source,
    })
    this.map.addLayer(this.layer)
  }
}
</script>