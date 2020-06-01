<template>
  <v-app>
    <v-content>
      <v-container fluid>
        <v-row>
          <v-col cols="auto">
            <v-card width="400" min-height="500">
              <v-tabs grow v-model="tab">
                <v-tab key="status">Status</v-tab>
                <v-tab key="game">Game</v-tab>
                <v-tab key="settings">Settings</v-tab>
                <v-tab key="profiles">Profile</v-tab>
              </v-tabs>
              <v-tabs-items v-model="tab">
                <v-tab-item key="status">
                  <status></status>
                </v-tab-item>
                <v-tab-item key="game">
                  <game></game>
                </v-tab-item>
                <v-tab-item key="settings">
                  <settings
                  ></settings>
                </v-tab-item>
                <v-tab-item key="profiles">
                  <profiles
                  ></profiles>
                </v-tab-item>
              </v-tabs-items>
              <component
                v-if="dialog !== null"
                :is="dialog"
              ></component>
            </v-card>
          </v-col>
          <v-col cols="auto">
            <v-sheet
              v-if="layout"
              style="position: fixed"
              :width="layout.width"
              :min-height="layout.height"
            >
              <template v-for="(widget, i) in layout.widgets">
                <component
                  v-if="widget.visible"
                  :key="i"
                  :is="widget.name"
                  :settings="widget.settings"
                >
                </component>
              </template>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
  import { mapState, mapActions } from 'vuex';

  // Tab components
  import Status from '@/components/tabs/Status.vue';
  import Game from '@/components/tabs/Game.vue';
  import Settings from '@/components/tabs/Settings.vue';
  import Profiles from '@/components/tabs/Profiles.vue';
  // Widget components
  import WidgetTitle from '@/components/widgets/Title.vue';
  import WidgetText from '@/components/widgets/Text.vue';
  import WidgetProgress from '@/components/widgets/ProgressBar.vue';
  import WidgetUnlocked from '@/components/widgets/LatestUnlocks.vue';
  import WidgetLocked from '@/components/widgets/LockedAchievements.vue';
  import WidgetFocused from '@/components/widgets/FocusedAchievement.vue';
  import WidgetAchievementOverlay from '@/components/widgets/AchievementOverlay.vue';
  import WidgetOverlay from '@/components/widgets/Overlay.vue';
  // Dialog components
  import AddProfile from '@/components/dialogs/AddProfile.vue';
  import DeleteProfile from '@/components/dialogs/DeleteProfile.vue';
  import DuplicateLayout from '@/components/dialogs/DuplicateLayout.vue';
  import ChangeLayout from '@/components/dialogs/ChangeLayout.vue';
  import WidgetSettings from '@/components/dialogs/WidgetSettings.vue';

  export default {
    data: () => ({
      tab: 'status',
    }),

    components: {
      // tabs
      Status,
      Game,
      Settings,
      Profiles,
      // widgets
      WidgetTitle,
      WidgetText,
      WidgetProgress,
      WidgetUnlocked,
      WidgetLocked,
      WidgetAchievementOverlay,
      WidgetOverlay,
      WidgetFocused,
      // dialogs
      AddProfile,
      DeleteProfile,
      DuplicateLayout,
      ChangeLayout,
      WidgetSettings,
    },

    computed: {
      ...mapState('dialogs', ['dialog']),
      ...mapState('profile', ['profile']),
      ...mapState('layout', ['layout']),
    },

    methods: {
      ...mapActions(['getGameInfo']),
      ...mapActions('profile', ['getActiveProfile', 'getProfiles']),
      ...mapActions('layout', ['getActiveLayout', 'getLayouts']),
    },

    async created() {
      await this.getLayouts();
      await this.getProfiles();
      await this.getActiveLayout();
      await this.getActiveProfile();
      // TODO: get game info should be somewhere, but not created
      //await this.getGameInfo(this.profile);
    },
  }
</script>
