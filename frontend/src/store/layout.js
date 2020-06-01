import { getURL } from '../api/util';

export default {
  namespaced: true,

  state: {
    layouts: {},
    layout: null,
    saved: null,
    widget: null,
  },

  mutations: {
    SET_LAYOUTS(state, layouts) {
      state.layouts = Object.fromEntries(
        Object.entries(layouts).map(x => [x[0], JSON.parse(x[1])]),
      );
      //state.layouts = layouts;
    },
    ADD_LAYOUT(state, { name, width, height, auto, interval }) {
      state.layouts[name] = { name, width, height, auto, interval, widgets: [] };
    },
    REMOVE_LAYOUT(state, name) {
      delete state.layouts[name];
    },
    SET_ACTIVE_LAYOUT(state, name) {
      // chained parse/stringify to do a inefficient deep clone
      state.layout = JSON.parse(JSON.stringify(state.layouts[name]));
    },
    SAVED_LAYOUT(state, name) {
      // chained parse/stringify to do a inefficient deep clone
      state.layouts[name] = JSON.parse(JSON.stringify(state.layout));
      state.saved = JSON.parse(JSON.stringify(state.layouts[name]));
    },

    ADD_WIDGET(state, widget) {
      state.layout.widgets.push(widget);
    },
    REMOVE_WIDGET(state, index) {
      state.layout.widgets.splice(index, 1);
    },
    MOVE_WIDGET(state, { index, target }) {
      state.layout.widgets.splice(target, 0, ...state.layout.widgets.splice(index, 1));
    },
    SET_ACTIVE_WIDGET(state, widget) {
      state.widget = widget;
    },
    SET_WIDGET_VISIBILITY(state, { index, visible }) {
      state.layout.widgets[index].visible = visible;
    },
  },

  actions: {
    async getLayouts({ commit }) {
      try {
        const request = await fetch(getURL('/api/layout/list'));
        const layouts = await request.json();
        commit('SET_LAYOUTS', layouts.layouts);
        console.log('Successfully got layouts:', request.statusText);
      } catch (e) {
        console.error('Failed to get layouts:', e);
      }
    },
    async addLayout({ commit }, { name, width, height, auto, interval }) {
      commit('ADD_LAYOUT', { name, width, height, auto, interval });
      try {
        const payload = JSON.stringify({
          name,
          width,
          height,
          auto,
          interval,
          widgets: [],
        });
        const request = await fetch(getURL('/api/layout/add'), {
          method: 'POST',
          body: JSON.stringify({
            name,
            layout: payload,
          }),
        });
        console.log('Successfully added layout:', request.statusText);
      } catch (e) {
        console.error('Failed to add layout:', e);
      }
    },
    async removeLayout({ commit }, { name }) {
      commit('REMOVE_LAYOUT', name);
      try {
        const request = await fetch(getURL('/api/layout/remove'), {
          method: 'POST',
          body: JSON.stringify({
            name,
          }),
        });
        console.log('Successfully removed layout:', request.statusText);
      } catch (e) {
        console.error('Failed to remove layout:', e);
      }
    },
    async saveLayout({ commit, state }) {
      try {
        const payload = JSON.stringify(state.layout);
        const request = await fetch(getURL('/api/layout/update'), {
          method: 'POST',
          body: JSON.stringify({
            name: state.layout.name,
            layout: payload,
          }),
        });
        commit('SAVED_LAYOUT', state.layout.name);
        console.log('Successfully saved layout:', request.statusText);
      } catch (e) {
        console.error('Failed to save layout:', e);
      }
    },
    async getActiveLayout({ commit }) {
      try {
        const request = await fetch(getURL('/api/layout/active'));
        const layout = await request.json();
        console.log('got active layout:', layout.layout);
        commit('SET_ACTIVE_LAYOUT', layout.layout);
        commit('SAVED_LAYOUT', layout.layout);
      } catch (e) {
        console.error('Failed to get active layout:', e);
      }
    },
    async setActiveLayout({ commit }, { name }) {

      // TODO check if layout exists
      commit('SET_ACTIVE_LAYOUT', name);
      commit('SAVED_LAYOUT', name);
      try {
        const request = await fetch(getURL('/api/layout/active/set'), {
          method: 'POST',
          body: JSON.stringify({
            name,
          }),
        });
        console.log('Successfully set active layout:', request.statusText);
      } catch (e) {
        console.error('Failed to save active layout:', e);
      }
    },

    // widgets
    async addWidget({ commit }, widget) {
      commit('ADD_WIDGET', widget);
    },
    async moveWidget({ commit }, { index, target }) {
      commit('MOVE_WIDGET', { index, target });
    },
    async deleteWidget({ commit }, index) {
      commit('REMOVE_WIDGET', index);
    },
    async setActiveWidget({ commit }, widget) {
      commit('SET_ACTIVE_WIDGET', widget);
    },
    async setWidgetVisibility({ commit }, { index, visible }) {
      commit('SET_WIDGET_VISIBILITY', { index, visible });
    },
  },
};
