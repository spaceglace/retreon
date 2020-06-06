export default {
  namespaced: true,

  state: {
    dialog: null,
  },

  mutations: {
    SET_DIALOG(state, name) {
      state.dialog = name;
    },
  },

  actions: {
    async openDialog({ commit }, { name }) {
      commit('SET_DIALOG', name);
    },
    async closeDialog({ commit }) {
      commit('SET_DIALOG', null);
    },
  },
};
