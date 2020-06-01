<template>
  <v-row>
    <v-col
      class="px-5 text-center subtitle-1"
    >{{ settings.text }}</v-col>
  </v-row>
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
