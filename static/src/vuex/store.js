import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const state = {
  photos: []
}

const getters = {
  photos: state => state.photos
}

const mutations = {
  addPhoto (state, photo) {
    state.photos.push(photo)
  },
  clearPhotos (state) {
    state.photos = []
  }
}

export default new Vuex.Store({
  state,
  getters,
  mutations
})