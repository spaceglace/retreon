import Vue from 'vue';
import Vuex from 'vuex';

import game from './game';
import dialogs from './dialogs';
import profile from './profile';
import layout from './layout';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    game,
    dialogs,
    profile,
    layout,
  },
});
