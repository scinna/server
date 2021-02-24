import Vue from 'vue'
import Vuex from 'vuex'
import Account, {AccountStateProps} from "@/store/Account";
import Server, {ServerProps} from "@/store/Server";

Vue.use(Vuex)

export type GlobalState = {
  Account: AccountStateProps;
  Server: ServerProps;
}

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
