<template>
  <v-container>
    <v-alert
      v-if="profile.key === ''"
      dense
      outlined
      type="error"
    >No profile set, cannot query API</v-alert>
    <v-row justify="space-between" class="mt-6">
      <v-col cols="auto">
        Last refresh was {{ updateDuration }} ago
      </v-col>
      <v-col cols="auto">
        <v-btn
          small
          @click="update()"
          :loading="loading"
          class="mr-2"
          color="primary"
        >Refresh Now</v-btn>
      </v-col>
    </v-row>
    <v-row no-gutters class="mt-3">
      <v-col>
        Auto refresh interval in seconds
      </v-col>
    </v-row>
    <v-row no-gutters>
      <v-col>
        <v-slider
          :tick-labels="['Off', '5', '10', '15', '20', '25', '30']"
          :max="6"
          :step="1"
          ticks="always"
          tick-size="4"
          v-model="refreshRate"
          @change="test"
        ></v-slider>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  data: () => ({
    loading: false,
  }),

  computed: {
    ...mapState('profile', ['profile']),
    ...mapState('game', ['handle', 'timer', 'refresh']),

    updateDuration() {
      let seconds = this.timer;
      const hours = Math.floor(seconds / 3600);
      seconds -= hours * 3600;
      const minutes = Math.floor(seconds / 60);
      seconds -= minutes * 60;

      const hourChunk = hours > 0 ? `${hours}h ` : '';
      const minuteChunk = minutes > 0 || hours > 0 ? `${minutes}m ` : '';
      return `${hourChunk}${minuteChunk}${seconds}s`;
    },

    refreshRate: {
      get() {
        return this.refresh;
      },
      set(val) {},
    },
  },

  methods: {
    ...mapActions('game', [
      'getGameInfo', 'getMetadata',
      'setRefreshTicker', 'clearRefreshTicker',
      'incrementTimer', 'resetTimer',
      'setRefresh',
    ]),
    tick() {
      this.incrementTimer();
      if (this.refresh === 0) return;
      if (this.timer >= this.refresh * 5) {
        this.update();
      }
    },
    async update() {
      this.loading = true;
      await this.getGameInfo(this.profile);
      this.loading = false;
      this.resetTimer();
    },
    test(val) {
      // have this as a separate event instead of a computed setter
      // so that it doesn't proc on the initial value being set
      this.setRefresh(val);
    }
  },

  mounted() {
    this.getMetadata();
    if (this.handle !== null) this.clearRefreshTicker();
    this.setRefreshTicker(setInterval(this.tick, 1000));
  },
};
</script>
