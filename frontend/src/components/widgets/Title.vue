<template>
  <v-img
    :src="imageSource"
    :height="settings.height"
  >
    <v-card-title
      style="white-space: pre-wrap; word-break: break-word; background-color:#00000070; height:100%; text-align:center; text-shadow:1px 1px 2px black, 0 0 25px black, 0 0 5px black"
      class="display-1 font-weight-light justify-center"
    >
      <div>
        <div
          v-if="prefix !== ''"
          class="body-1"
        >{{ prefix }}</div>
        <span>{{ title }}</span>
      </div>
    </v-card-title>
  </v-img>
</template>

<script>
import { mapState } from 'vuex';

export default {
  props: [
    'settings',
  ],

  data: () => ({
    base: 'http://i.retroachievements.org.s3.amazonaws.com',
  }),

  computed: {
    ...mapState('game', ['game']),

    imageSource() {
      if (this.game === null) return '';
      return `${this.base}${this.game[this.settings.icon]}`;
    },

    prefix() {
      if (this.game === null) return '';
      if (!this.settings.beautify) return '';
      const chunks = this.game.title.match(/(~[^ ]+~ )?(.+?)(, the)?$/i);
      if (chunks[1]) return chunks[1].trim();
      return '';
    },

    title() {
      if (this.game === null) return '';
      let result = this.game.title;
      if (this.settings.beautify) {
        const fullTitle = result.match(/(~[^ ]+~ )?(.+?)(, the)?$/i);
        result = fullTitle[2];
        if (fullTitle[3]) result = `The ${result}`;
      }

      // turn dashes into line breaks
      result = result.replace(' - ', '\n');
      return result;
    },
  },
};
</script>
