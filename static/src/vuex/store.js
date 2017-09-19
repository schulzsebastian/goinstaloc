import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const state = {
  photos: [],
  scanned: []
}

const getters = {
  photos: state => state.photos,
  scanned: state => state.scanned
}

const mutations = {
  addPhoto (state, photo) {
    state.photos.push(photo)
  },
  clearPhotos (state) {
    state.photos = []
  },
  addScanned (state, scanned) {
    state.scanned.push(scanned)
  },
  clearScanned (state) {
    state.scanned = []
  }
}

export default new Vuex.Store({
  state,
  getters,
  mutations
})