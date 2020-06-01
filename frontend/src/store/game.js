import { getURL } from '../api/util';

export default {
  namespaced: true,

  state: {
    game: null,
    gameMode: 'DateEarned',
    handle: null,
    timer: 0,
    order: [],
    refresh: 0,
  },

  mutations: {
    SET_GAME_INFO(state, game) {
      state.game = game;
    },
    SET_HANDLE(state, handle) {
      state.handle = handle;
    },
    RESET_HANDLE(state) {
      clearInterval(state.handle);
      state.handle = null;
    },
    INCREMENT_TIMER(state) {
      state.timer += 1;
    },
    RESET_TIMER(state) {
      state.timer = 0;
    },
    SET_ORDER(state, order) {
      state.order = order;
    },
    SET_MODE(state, hardcore) {
      state.gameMode = hardcore;
    },
    SET_REFRESH(state, refresh) {
      state.refresh = refresh;
    },
  },

  actions: {
    async getGameInfo({ commit }, { name, key }) {
      try {
        const request = await fetch(getURL('/api/game'), {
          method: 'POST',
          body: JSON.stringify({
            name,
            key,
          }),
        });
        const result = await request.json();
        commit('SET_GAME_INFO', result);
        console.log('Got last game:', request.statusText);
      } catch (e) {
        console.error('Failed to get game info:', e);
      }
    },

    async setRefreshTicker({ commit }, handle) {
      commit('SET_HANDLE', handle);
    },
    async clearRefreshTicker({ commit }) {
      commit('RESET_HANDLE');
    },

    async resetTimer({ commit }) {
      commit('RESET_TIMER');
    },
    async incrementTimer({ commit }) {
      commit('INCREMENT_TIMER');
    },

    async setGameOrder({ commit }, { game, order }) {
      commit('SET_ORDER', order);
      try {
        const request = await fetch(getURL('/api/game/order'), {
          method: 'POST',
          body: JSON.stringify({
            game,
            order,
          }),
        });
        console.log('Successfuly set game order:', request.statusText);
      } catch (e) {
        console.error('Failed to set game order:', e);
      }
    },
    async setGameMode({ commit }, mode) {
      commit('SET_MODE', mode);
      try {
        const request = await fetch(getURL('/api/metadata/mode'), {
          method: 'POST',
          body: JSON.stringify({
            mode,
          }),
        });
        console.log('Successfully set game mode:', request.statusText);
      } catch (e) {
        console.error('Failed to set game mode:', e);
      }
    },
    async setRefresh({ commit }, refresh) {
      commit('SET_REFRESH', refresh);
      try {
        const request = await fetch(getURL('/api/metadata/refresh'), {
          method: 'POST',
          body: JSON.stringify({
            refresh,
          }),
        });
        console.log('Status of setRefresh request:', request.statusText);
      } catch (e) {
        console.error('Failed to set refresh:', e);
      }
    },
    async getMetadata({ commit }) {
      try {
        const request = await fetch(getURL('/api/metadata'));
        const result = await request.json();
        commit('SET_MODE', result.mode);
        commit('SET_REFRESH', result.refresh);
        console.log('Status of getMetadata:', request.statusText);
      } catch (e) {
        console.error('Failed to get metadata:', e);
      }
    },
  },
};
