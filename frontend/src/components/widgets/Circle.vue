<template>
  <div
    id="outerWrapper"
    :style="{
      height: settings.height + 'px',
    }"
  >
    <div
      id="wrapper"
      :style="{
        'margin-top': (Math.floor(settings.height / 2) - 23) + 'px',
      }"
    >
      <div
        v-for="entry in gameClone"
        :key="entry.offset"
        :class="['entry', isAnimating ? `spin` : '']"
        :style="`
          width: ${sliceWidth + 530}px;
          z-index: ${count - Math.abs(entry.offset)};
          --var-duration: 0.5s;
          --var-start: ${entry.offset * 6}deg;
          --var-stop: ${(entry.offset * 6) + 6}deg;
          --var-innerWidth: ${sliceWidth}px;
        `"
      >
        <div class="inner">
          <v-row no-gutters align="center">
            <v-col cols="auto">
              <div
                :style="{
                  width: '55px',
                  height: '55px',
                  'background-size': 'cover',
                  'background-image': `url(${base}/Badge/${entry.BadgeName}${entry.offset < 1 ? '_lock' : ''}.png)`,
                }"
                class="elevation-20 mr-2"
              ></div>
            </v-col>
            <v-col>
              <div
                :class="['subtitle-1', 'text-no-wrap', 'innerText', entry.offset === 0 ? 'font-weight-bold' : '']"
              >{{ entry.Title }}</div>
              <div
                class="caption font-weight-light text-no-wrap innerText"
              >{{ entry.Description }}</div>
            </v-col>
          </v-row>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex';

export default {
  props: [
    'settings',
  ],

  data: () => ({
    base: 'http://i.retroachievements.org.s3.amazonaws.com',
    isAnimating: false,
    count: 6,
    offset: 0,
    gameClone: {},
    last: null,
  }),

  computed: {
    ...mapState('layout', ['layout']),
    ...mapState('game', ['game', 'gameMode']),

    sliceWidth() {
      if (!this.layout) return 0;
      return this.layout.width - 90;
    },

    runway() {
      if (!this.game) return [];

      let locked = this.game.order
        .filter(x => !this.game.achievements[x][this.gameMode])
        .slice(0, this.count + 1)
        .map(x => this.game.achievements[x]);
      // use offset to skip blank spaces if there's not enough locked
      locked = locked
        .reverse()
        .map((x, id) => ({ offset: id - locked.length + 1, ...x }))
      const unlocked = Object.values(this.game.achievements)
        .filter(x => x[this.gameMode])
        .sort((a,b) => b[this.gameMode].localeCompare(a[this.gameMode]))
        .slice(0, this.count)
        .map((x, id) => ({ offset: id + 1, ...x }));

      return [...locked, ...unlocked];
    },
  },

  methods: {
    clone() {
      if (this.runway.length < 1) return;

      console.log('in clone');
      if (this.last && this.runway[this.count + 1].ID !== this.last) {
        this.isAnimating = true;
        setTimeout(() => {
          this.gameClone = [...this.runway];
          this.isAnimating = false;
        }, 500);
      } else {
        console.log('in the else');
        this.gameClone = [...this.runway];
      }
      this.last = this.runway[this.count + 1].ID;
    },
  },

  watch: {
    'game.order'() {
      console.log("order change");
      this.clone();
    }
  },

  created() {
    this.clone();
  },
}
</script>

<style scoped>
  #outerWrapper {
    overflow: hidden;
  }
  #wrapper {
    position: relative;
    left: -445px;
  }
  .entry {
    width: 640px;
    position: absolute;
    transform-origin: center left;
    transform: rotate(var(--var-start));
  }
  .inner {
    float: right;
  }
  .innerText {
    overflow: hidden;
    text-overflow: ellipsis;
    text-shadow: 0 0 4px black;
    width: var(--var-innerWidth);
  }
  .spin {
    animation: spin var(--var-duration) linear;
  }
  @keyframes spin {
    0% { transform: rotate(var(--var-start)); }
    100% { transform: rotate(var(--var-stop)); }
  }
</style>
