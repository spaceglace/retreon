<template>
  <v-overlay
    absolute
    opacity="0.9"
  >
    <v-row>
      <v-col
        class="px-6 text-center display-1 font-weight-light"
        style="white-space: pre-wrap; word-break: break-word"
      >{{ text }}</v-col>
    </v-row>
  </v-overlay>
</template>

<script>
export default {
  props: [
    'settings',
  ],

  computed: {
    text() {
      return this.settings.text;
    },
  },
};
</script>
