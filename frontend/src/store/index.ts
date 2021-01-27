import Vue from 'vue'
import Vuex from 'vuex'
import Account from "@/store/Account";
import Server from "@/store/Server";

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    Account,
    Server,
  }
})
