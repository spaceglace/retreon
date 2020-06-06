<template>
  <v-container>
    <v-row no-gutters>
      <v-col>
        <v-select
          v-model="activeLayout"
          :items="Object.keys(layouts)"
          label="Active Layout"
        ></v-select>
      </v-col>
    </v-row>
    <v-row dense class="pb-9">
      <v-col cols="auto">
        <v-btn
          color="primary darken-1"
          small
          :disabled="!hasLayoutChanged"
          @click="saveLayout()"
        >Save</v-btn>
      </v-col>
      <v-spacer></v-spacer>
      <v-col cols="auto">
        <v-btn
          color="primary darken-2"
          small
          @click="openDialog({ name: 'duplicate-layout' })"
        >Duplicate</v-btn>
      </v-col>
      <v-col cols="auto">
        <v-btn
          color="error darken-2"
          small
          :disabled="activeLayout === 'default'"
        >Delete</v-btn>
      </v-col>
    </v-row>
    <v-row class="pb-3">
      <v-tabs grow v-model="tab">
        <v-tab key="widgets">Widgets</v-tab>
        <v-tab key="layout">Layout</v-tab>
      </v-tabs>
    </v-row>
    <v-tabs-items v-model="tab">
      <v-tab-item key="widgets">
        <v-form>
          <v-row dense align="baseline">
            <v-col>
              <v-select
                :items="widgetList"
                item-text="title"
                item-value="name"
                v-model="addSelection"
                label="Add Widget"
              ></v-select>
            </v-col>
            <v-col cols="auto">
              <v-btn
                :disabled="addSelection === null"
                color="primary darken-2"
                @click="newWidget()"
              >Add</v-btn>
            </v-col>
          </v-row>
        </v-form>
        <v-list dense>
          <v-list-item
            v-for="item in indexedWidgets"
            :key="item.id"
            style="min-height: 20px"
          >
            <v-list-item-content>
              <v-list-item-title v-text="item.title"></v-list-item-title>
            </v-list-item-content>
            <v-list-item-action class="my-0">
              <v-row no-gutters>
                <v-col>
                  <v-btn
                    icon
                    small
                    :disabled="item.id === 0"
                    @click="widgetUp(item)"
                  >
                    <v-icon>mdi-chevron-up</v-icon>
                  </v-btn>
                </v-col>
                <v-col>
                  <v-btn
                    icon
                    small
                    :disabled="item.id === layout.widgets.length - 1"
                    @click="widgetDown(item)"
                  >
                    <v-icon>mdi-chevron-down</v-icon>
                  </v-btn>
                </v-col>
                <v-col>
                  <v-btn
                    icon
                    x-small
                    @click="toggleVisibility(item)"
                  >
                    <v-icon>{{ item.visible ? 'mdi-eye' : 'mdi-eye-off' }}</v-icon>
                  </v-btn>
                </v-col>
                <v-col class="ml-1">
                  <v-btn
                    icon
                    x-small
                    @click="widgetDetails(item)"
                  >
                    <v-icon>mdi-cogs</v-icon>
                  </v-btn>
                </v-col>
              </v-row>
            </v-list-item-action>
          </v-list-item>
        </v-list>
      </v-tab-item>
      <v-tab-item key="layout">
        <v-form>
          <v-row dense align="end">
            <v-col>
              <v-slider
                label="Width"
                :min="minWidth"
                :max="maxWidth"
                :rules="[rules.widthMin, rules.widthMax]"
                v-model="layout.width"
              ></v-slider>
            </v-col>
            <v-col cols="3">
              <v-text-field
                v-model="layout.width"
                type="number"
                :rules="[rules.widthMin, rules.widthMax]"
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row dense align="end">
            <v-col>
              <v-slider
                label="Height"
                :min="minHeight"
                :max="maxHeight"
                v-model="layout.height"
              ></v-slider>
            </v-col>
            <v-col cols="3">
              <v-text-field
                v-model="layout.height"
                type="number"
                :rules="[rules.heightMin, rules.heightMax]"
              ></v-text-field>
            </v-col>
          </v-row>
        </v-form>
      </v-tab-item>
    </v-tabs-items>
  </v-container>
</template>

<script>
import { mapState, mapActions } from 'vuex';
import widgets from '@/components/widgets/index.js';

export default {
  data: () => ({
    tab: 'widgets',
    // layout
    minWidth: 200,
    maxWidth: 900,
    minHeight: 200,
    maxHeight: 1500,

    // widgets
    addSelection: null,
    headers: [
      { text: 'Widget', value: 'title' },
      { text: 'Actions', value: 'actions', align: 'right' },
    ],
  }),

  computed: {
    ...mapState('game', ['game']),
    ...mapState('layout', ['layout', 'saved', 'layouts']),

    rules: function() {
      return {
        widthMin: v => v >= this.minWidth || `Min ${this.minWidth}`,
        widthMax: v => v <= this.maxWidth || `Max ${this.maxWidth}`,
        heightMin: v => v >= this.minHeight || `Min ${this.minHeight}`,
        heightMax: v => v <= this.maxHeight || `Max ${this.maxHeight}`,
        intervalMin: v => v >= 5 || 'Minimum 5 seconds',
        intervalMax: v => v <= 60 || 'Maximum 60 seconds',
      };
    },
    widgetList() {
      return Object.values(widgets);
    },
    indexedWidgets() {
      return this.layout.widgets.map((item, index) => ({
        id: index,
        ...item,
      }));
    },

    activeLayout: {
      get() {
        return this.layout.name;
      },
      set(val) {
        console.log(val, this.hasLayoutChanged);

        if (this.hasLayoutChanged) {
          this.openDialog({ name: 'change-layout' });
        } else {
          this.setActiveLayout({ name: val });
        }
      },
    },

    hasLayoutChanged() {
      // Check low hanging fruit before using JSON
      if (this.layout.width !== this.saved.width) return true;
      if (this.layout.height !== this.saved.height) return true;
      if (this.layout.auto !== this.saved.auto) return true;
      if (this.layout.interval !== this.saved.interval) return true;
      if (this.layout.widgets.length !== this.saved.widgets.length) return true;
      // Widgets can be complicated so just serialized and check
      return JSON.stringify(this.layout) !== JSON.stringify(this.saved);
    }
  },

  methods: {
    ...mapActions('dialogs', ['openDialog']),
    ...mapActions('layout', [
      'setActiveLayout', 'saveLayout',
      'addWidget', 'moveWidget', 'setActiveWidget',
      'setWidgetVisibility',
    ]),

    // widgets
    newWidget() {
      const result = {
        name: widgets[this.addSelection].name,
        title: widgets[this.addSelection].title,
        settings: {},
        visible: true,
      };

      widgets[this.addSelection].settings.forEach((setting) => {
        result.settings[setting.name] = setting.default;
      });

      this.addWidget(result);
    },
    widgetUp(item) {
      if (item.id === 0) return;
      this.moveWidget({ index: item.id, target: item.id - 1 });
    },
    widgetDown(item) {
      if (item.id === this.layout.widgets.length - 1) return;
      this.moveWidget({ index: item.id, target: item.id + 1 });
    },
    widgetDetails(item) {
      this.setActiveWidget(item);
      this.openDialog({ name: 'widget-settings' });
    },

    toggleVisibility(item) {
      this.setWidgetVisibility({ index: item.id, visible: !item.visible });
    },
  },
}
</script>
