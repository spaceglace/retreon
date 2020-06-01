<template>
  <v-overlay
    absolute
    :opacity="0.9"
  >
    <v-container>
      <v-row>
        <v-col>
          <div>Change Layout</div>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-btn @click="close()">Yeet</v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-overlay>
</template>

<script>
import { mapActions } from 'vuex';
export default {
  methods: {
    ...mapActions('dialogs', ['closeDialog']),
    close() {
      this.closeDialog();
    },
  },
};
</script>
