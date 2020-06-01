<template>
  <v-container fluid>
    <v-row>
      <v-col>
        <v-select
          v-if="profiles !== null"
          :items="Object.keys(profiles)"
          v-model="activeProfile"
          label="Active Profile"
        ></v-select>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="auto">
        <v-btn
          color="success darken-2"
          @click="openDialog({ name: 'add-profile' })"
        >Add Profile</v-btn>
      </v-col>
      <v-col>
        <v-btn
          v-if="profile !== null"
          color="error darken-2"
          @click="openDialog({ name: 'delete-profile' })"
        >Delete Profile</v-btn>
      </v-col>
    </v-row>

    <v-overlay
      absolute
      :value="showCreate"
      :opacity="0.9"
    >
      <v-container>
        <v-row>
          <v-col>
            <div>Create Profile</div>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-btn @click="showCreate = false">Yeet</v-btn>
          </v-col>
        </v-row>
      </v-container>
    </v-overlay>
  </v-container>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  data: () => ({
    showCreate: false,
    showDelete: false,
  }),
  computed: {
    ...mapState('profile', ['profile', 'profiles']),

    activeProfile: {
      get() {
        return this.profile ? this.profile.name : '';
      },
      set(val) {
        if (val !== this.profile) {
          this.setActiveProfile({ name: val });
        }
      },
    },
  },
  methods: {
    ...mapActions('profile', ['setActiveProfile']),
    ...mapActions('dialogs', ['openDialog']),
  },
};
</script>
