<template>
  <v-container>
    <div class="overline mb-1">
      Current Goal - {{ focused.Points }} ({{ focused.TrueRatio }}) points
    </div>
    <v-img
      :src="`http://i.retroachievements.org.s3.amazonaws.com/Badge/${focused.BadgeName}.png`"
      max-width="96"
      max-height="96"
      class="elevation-8 float-left mr-4"
    ></v-img>
    <div>
      <div class="subtitle-1 font-weight-bold text-center">
        {{ focused.Title }}
      </div>
      <v-divider></v-divider>
      <div style="overflow:hidden" class="my-3">
        <div
          :class="{
            'font-weight-light': true,
            'body-2': true,
            'text-no-wrap': settings.marquee,
            'text-center': true,
          }"
          :id="settings.marquee ? 'marquee' : ''"
        >{{ focused.Description }}</div>
      </div>
    </div>
    <div
      v-if="settings.softPercent"
      class="text-center caption"
    >
      {{ softcorePercent }}% of active Softcore players have this
    </div>
    <div
      v-if="settings.hardPercent"
      class="text-center caption"
    >
      {{ hardcorePercent }}% of active Hardcore players have this
    </div>
  </v-container>
</template>

<script>
import { mapState } from 'vuex';

export default {
  props: ['settings'],

  computed: {
    ...mapState('game', ['game', 'gameMode']),

    focused() {
      if (this.game === null) return {};
      const candidates = this.game.order
        .filter(x => !this.game.achievements[x][this.gameMode]);
      // TODO: make special 'u got em all' thing
      if (candidates.length === 0) return {};
      return this.game.achievements[candidates[0]];
    },
    softcorePercent() {
      if (this.game === null) return 0;
      return Math.floor(Number(this.focused.NumAwarded) / Number(this.game.softcore) * 10000) / 100;
    },
    hardcorePercent() {
      if (this.game === null) return 0;
      return Math.floor(Number(this.focused.NumAwardedHardcore) / Number(this.game.hardcore) * 10000) / 100;
    }
  },
};
</script>

<style scoped>
#marquee {
  animation: scroll 10s linear 0s infinite;
  display: inline-block;
}
@keyframes scroll {
    from {
        margin-left: 100%;
        transform: translate(0%,0)
    }

    to {
        margin-left: 0;
        transform: translate(-100%,0)
    }
}
</style>