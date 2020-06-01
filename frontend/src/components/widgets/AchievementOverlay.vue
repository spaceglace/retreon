<template>
  <v-overlay
    v-if="current"
    absolute
    opacity="0.8"
  >
    <div
      :class="[settings.shakeType, 'shake-constant']"
    >
      <v-container>
        <v-row justify="center">
          <v-col cols="auto">
            <div
              :style="{
                height: '150px',
                width: '150px',
                'background-size': 'cover',
                'background-image': settings.testAchieve
                  ? `url(${require('@/assets/images/glaceon.jpg')})`
                  : `url(http://i.retroachievements.org.s3.amazonaws.com/Badge/${current.BadgeName}.png)`,
                'box-shadow': '0px 0px 40px 4px #FFF',
              }"
            ></div>
          </v-col>
        </v-row>
        <v-row justify="center">
          <div class="subtitle-1 px-3 font-weight-bold text-center mt-5">{{ current.Title }}</div>
        </v-row>
        <v-row justify="center">
          <div class="caption font-weight-light text-center px-2 mt-4">{{ current.Description }}</div>
        </v-row>
        <v-row justify="center">
          <div
            class="headline font-weight-bold mt-5 text-center"
          >+{{ current.Points }} (+{{ current.TrueRatio }}) points</div>
        </v-row>
      </v-container>
    </div>
  </v-overlay>
</template>

<script>
import { mapState } from 'vuex';

export default {
  props: [
    'settings',
  ],

  data: () => ({
    queue: [],
    current: null,
  }),
  
  computed: {
    ...mapState('game', ['game', 'gameMode']),
  },

  methods: {
    showNext() {
      if (this.queue.length === 0) {
        this.current = null;
        return;
      }

      this.current = this.game.achievements[this.queue.shift()];
      console.log(this.current);
      setTimeout(this.showNext, 5000);
    },
  },

  watch: {
    game: function(newValue, oldValue) {
      if (!oldValue || oldValue.id !== newValue.id) return;
      const unlocks = Object.values(oldValue.achievements)
        .filter((achievement) => {
          return !achievement[this.gameMode]
            && newValue.achievements[achievement.ID][this.gameMode];
        })
        .map(x => x.ID);
      const priorEmpty = this.queue.length === 0;
      this.queue.push(...unlocks);
      if (priorEmpty) this.showNext();
    },

    'settings.testAchieve': function(newValue) {
      if (newValue) {
        this.current = {
          Title: "i am glaceon",
          Description: "This is a test achievement for testing settings",
          Points: 0,
          TrueRatio: 0,
          isTest: true,
        };
      } else {
        this.current = null;
        this.queue = [];
      }
    },
  },
};
</script>

<style>
@import '../../style/shake.css';
</style>
