<template>
  <v-overlay
    absolute
    :opacity="0.9"
  >
    <v-container>
      <v-row>
        <v-col>
          <v-text-field
            label="New Layout Name"
            v-model="name"
            :rules="[rules.name]"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row justify="space-around">
        <v-col cols="auto">
          <v-btn
            small
            color="success darken-2"
            :loading="loading"
            :disabled="loading"
            @click="create()"
          >Create</v-btn>
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
    name: "",
    loading: false,
    rules: {
      // TODO: check for duplicate names
      name: v => v != "" || 'Name is required',
    },
  }),

  computed: {
    ...mapState('layout', ['layout']),
  },

  methods: {
    ...mapActions('dialogs', ['closeDialog']),
    ...mapActions('layout', ['addLayout', 'setActiveLayout']),

    async create() {
      this.loading = true;
      await this.addLayout({
        name: this.name,
        width: this.layout.width,
        height: this.layout.height,
        auto: this.layout.auto,
        interval: this.layout.interval,
      });
      await this.setActiveLayout({ name: this.name });
      this.loading = false;
      this.closeDialog();
    },
  },
};
</script>
