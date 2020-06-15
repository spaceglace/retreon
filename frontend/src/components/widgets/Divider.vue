<template>
  <div
    :style="{
      'margin-top': `${settings.top}px`,
      'margin-bottom': `${settings.bottom}px`,
    }"
  >
    <v-divider
      v-if="settings.line"
    ></v-divider>
    <div v-else style="height: 1px"></div>
  </div>
</template>

<script>
export default {
  props: [
    'settings',
  ],
}
</script>
