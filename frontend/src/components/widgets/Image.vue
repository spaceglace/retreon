<template>
  <v-img
    :src="settings.url"
    :height="settings.height"
    contain
  ></v-img>
</template>

<script>
export default {
  props: [
    'settings',
  ],
};
</script>
