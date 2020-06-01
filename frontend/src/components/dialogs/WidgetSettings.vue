<template>
  <v-overlay
    absolute
    :opacity="0.8"
  >
    <v-card
      raised
      color="grey darken-3"
      width="300"
    >
      <v-card-title
        class="subtitle-1 font-weight-light"
      >{{ widget.title }} Settings</v-card-title>
      <v-card-text>
        <v-row
          v-for="(w, k) in widgetSettings"
          :key="k"
          no-gutters
        >
          <v-col>
            <v-text-field
              v-if="w.type === 'string'"
              :label="w.title"
              v-model="widget.settings[w.name]"
            ></v-text-field>
            <v-switch
              v-else-if="w.type === 'boolean'"
              :label="w.title"
              v-model="widget.settings[w.name]"
              hide-details
            ></v-switch>
            <v-slider
              v-else-if="w.type === 'slider'"
              :min="w.min"
              :max="w.max"
              :label="w.title"
              v-model="widget.settings[w.name]"
              hide-details
              thumb-label
            ></v-slider>
            <v-select
              v-else-if="w.type === 'selection'"
              class="mt-4"
              :items="w.options"
              :label="w.title"
              v-model="widget.settings[w.name]"
            ></v-select>
            <v-textarea
              v-else-if="w.type === 'textarea'"
              :label="w.title"
              v-model="widget.settings[w.name]"
            ></v-textarea>
          </v-col>
        </v-row>
      </v-card-text>
      <v-card-actions>
        <v-btn
          @click="remove()"
          color="error darken-2"
          small
        >Delete</v-btn>
        <v-spacer></v-spacer>
        <v-btn
          @click="closeDialog()"
          color="primary darken-1"
          small
        >Close</v-btn>
      </v-card-actions>
    </v-card>
  </v-overlay>
</template>

<script>
import { mapState, mapActions } from 'vuex';
import widgets from '@/components/widgets/index.js';


export default {
  computed: {
    ...mapState('layout', ['widget']),
    widgetSettings() {
      return widgets[this.widget.name].settings;
    },
  },

  methods: {
    ...mapActions('dialogs', ['closeDialog']),
    ...mapActions('layout', ['deleteWidget']),

    remove() {
      this.deleteWidget(this.widget.id);
      this.closeDialog();
    },
  },

}
</script>