<template>
  <v-container class="py-0">
    <v-row>
      <v-progress-linear
        v-if="settings.showCount"
        v-model="count.percent"
        height="20"
        class="elevation-8"
        :striped="count.locked === 0"
        :color="count.locked === 0 ? 'yellow darken-3' : 'primary darken-1'"
      >
        <span
          class="font-weight-bold overline"
          style="text-shadow: 0 0 3px black"
        >achievements - {{ count.unlocked }} / {{ count.total }} ({{ count.percent }}%)</span>
      </v-progress-linear>
      <v-progress-linear
        v-if="settings.showPoints"
        v-model="points.percent"
        height="20"
        class="elevation-8 mt-1"
        :striped="points.locked === 0"
        :color="points.locked === 0 ? 'yellow darken-3' : 'indigo lighten-1'"
      >
        <span
          class="font-weight-bold overline"
          style="text-shadow: 0 0 3px black"
        >points - {{ points.unlocked }} / {{ points.total }} ({{ points.percent }}%)</span>
      </v-progress-linear>
      <v-progress-linear
        v-if="settings.showRatio"
        v-model="ratio.percent"
        height="20"
        class="elevation-8 mt-1"
        :striped="ratio.locked === 0"
        :color="ratio.locked === 0 ? 'yellow darken-3' : 'purple lighten-1'"
      >
        <span
          class="font-weight-bold overline"
          style="text-shadow: 0 0 3px black"
        >retro ratio - {{ ratio.unlocked }} / {{ ratio.total }} ({{ ratio.percent }}%)</span>
      </v-progress-linear>
    </v-row>
  </v-container>
</template>

<script>
import { mapState } from 'vuex';

export default {
  props: [
    'settings',
  ],

  computed: {
    ...mapState('game', ['game', 'gameMode']),

    count() {
      const result = { locked: 0, unlocked: 0, total: 0, percent: 0 };
      if (this.game === null) return result;

      Object.values(this.game.achievements).forEach((achievement) => {
        result.total += 1;
        if (achievement[this.gameMode]) {
          result.unlocked += 1;
        } else {
          result.locked += 1;
        }
      });
      result.percent = Math.floor((result.unlocked / result.total) * 100);
      return result;
    },

    points() {
      const result = { locked: 0, unlocked: 0, total: 0, percent: 0 };
      if (this.game === null) return result;

      Object.values(this.game.achievements).forEach((achievement) => {
        const points = Number(achievement.Points);
        result.total += points;
        if (achievement[this.gameMode]) {
          result.unlocked += points;
        } else {
          result.locked += points;
        }
      });
      result.percent = Math.floor((result.unlocked / result.total) * 100);
      return result;
    },

    ratio() {
      const result = { locked: 0, unlocked: 0, total: 0, percent: 0 };
      if (this.game === null) return result;

      Object.values(this.game.achievements).forEach((achievement) => {
        const points = Number(achievement.TrueRatio);
        result.total += points;
        if (achievement[this.gameMode]) {
          result.unlocked += points;
        } else {
          result.locked += points;
        }
      });
      result.percent = Math.floor((result.unlocked / result.total) * 100);
      return result;
    },
  },
};
</script>