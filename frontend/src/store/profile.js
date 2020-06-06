import { getURL } from '../api/util';

export default {
  namespaced: true,

  state: {
    profiles: {},
    profile: {},
  },

  mutations: {
    SET_PROFILES(state, profiles) {
      state.profiles = profiles;
    },
    ADD_PROFILE(state, { name, key }) {
      state.profiles[name] = { name, key };
    },
    REMOVE_PROFILE(state, name) {
      delete state.profiles[name];
    },
    SET_ACTIVE_PROFILE(state, name) {
      state.profile = state.profiles[name];
    },
  },

  actions: {
    async getProfiles({ commit }) {
      try {
        const request = await fetch(getURL('/api/profile/list'));
        const profiles = await request.json();
        commit('SET_PROFILES', profiles.profiles);
        console.log('getProfiles:', request.statusText);
      } catch (e) {
        console.error('Failed to fetch profiles:', e);
      }
    },
    async addProfile({ commit }, { name, key }) {
      // TODO: make sure profile doesn't exist yet
      commit('ADD_PROFILE', { name, key });
      commit('SET_ACTIVE_PROFILE', name);
      try {
        const result = await fetch(getURL('/api/profile/add'), {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            name,
            key,
          }),
        });
        console.log('addProfile:', result.statusText);
      } catch (e) {
        console.error('Failed to add profile:', e);
      }
    },
    async removeProfile({ commit, state, dispatch }, { name }) {
      commit('REMOVE_PROFILE', name);
      try {
        const result = await fetch(getURL('/api/profile/remove'), {
          method: 'POST',
          body: JSON.stringify({
            name,
          }),
        });
        console.log('removeProfile:', result.statusText);
      } catch (e) {
        console.error('Failed to remove profile:', e);
      }

      // if this was the active profile, switch to whatever's first in the list
      if (name === state.profile) {
        let newProfile = '';
        if (Object.keys(state.profiles).length > 0) {
          newProfile = Object.keys(state.profiles)[0];
        }
        dispatch('setActiveProfile', { name: newProfile });
      }
    },
    async getActiveProfile({ commit }) {
      try {
        const result = await fetch(getURL('/api/profile/active'))
        const profile = await result.json();
        console.log('getActiveProfile:', profile.profile);
        commit('SET_ACTIVE_PROFILE', profile.profile);
      } catch (e) {
        console.error('Failed to get active profile:', e);
      }
    },
    async setActiveProfile({ commit }, { name }) {
      // TODO check if profile exists
      commit('SET_ACTIVE_PROFILE', name);
      try {
        const result = await fetch(getURL('/api/profile/active/set'), {
          method: 'POST',
          body: JSON.stringify({
            name,
          }),
        });
        console.log('setActiveProfile:', result.statusText);
      } catch (e) {
        console.error('Failed to set active profile:', e);
      }
    },
  },
};
