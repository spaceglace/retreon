<template>
  <v-overlay
    absolute
    :opacity="0.9"
  >
    <v-container fluid>
      <v-row no-gutters>
        <v-col>
          <div>Are you sure you want to delete {{ profile.name }}?</div>
        </v-col>
      </v-row>
      <v-row dense justify="space-between">
        <v-col cols="auto">
          <v-btn
            small
            color="primary darken-2"
            :loading="loading"
            :disabled="loading"
            @click="confirm()"
          >I'm sure</v-btn>
        </v-col>
        <v-col cols="auto">
          <v-btn
            small
            color="error darken-2"
            :disabled="loading"
            @click="closeDialog()"
          >Never mind</v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-overlay>
</template>

<script>
import { mapState, mapActions } from 'vuex';
export default {
  data: () => ({
    loading: false,
  }),

  computed: {
    ...mapState('profile', ['profile']),
  },

  methods: {
    ...mapActions('dialogs', ['closeDialog']),
    ...mapActions('profile', ['removeProfile']),

    async confirm() {
      console.log('removing', this.profile);

      this.loading = true;
      await this.removeProfile({ name: this.profile });
      this.loading = false;
      this.closeDialog();
    }
  },
};
</script>
