<template>
  <v-overlay
    absolute
    :opacity="0.9"
  >
    <v-container fluid>
      <v-form>
        <v-row no-gutters>
          <v-col>
            <v-text-field
              label="Username"
              v-model="username"
              clearable
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row no-gutters>
          <v-col>
            <v-text-field
              label="API Key"
              v-model="key"
              type="password"
              clearable
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row justify="space-between">
          <v-col cols="auto">
            <v-btn
              small
              color="error darken-2"
              @click="cancel()"
            >Cancel</v-btn>
          </v-col>
          <v-col cols="auto">
            <v-btn
              small
              color="success darken-2"
              @click="submit()"
            >Submit</v-btn>
          </v-col>
        </v-row>
      </v-form>
    </v-container>
  </v-overlay>
</template>

<script>
import { mapActions } from 'vuex';
export default {
  data: () => ({
    username: "",
    key: "",
  }),
  methods: {
    ...mapActions('dialogs', ['closeDialog']),
    ...mapActions('profile', ['addProfile']),
    submit() {
      this.addProfile({
        name: this.username,
        key: this.key,
      });
      // TODO: show success/failure, close dialog?
    },
    cancel() {
      this.closeDialog();
    },
  },
};
</script>
