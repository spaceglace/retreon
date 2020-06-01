export default {
  'widget-text': {
    name: 'widget-text',
    title: 'Static Text',
    settings: [
      {
        name: 'text',
        title: 'Text',
        description: 'Static Text',
        type: 'string',
        default: '',
      },
    ],
  },
  'widget-title': {
    name: 'widget-title',
    title: 'Title',
    settings: [
      {
        name: 'height',
        title: 'Height',
        type: 'slider',
        default: 200,
        min: 100,
        max: 800,
      },
      {
        name: 'beautify',
        title: 'Format Title',
        type: 'boolean',
        default: true,
      },
    ],
  },
  'widget-progress': {
    name: 'widget-progress',
    title: 'Progress Bars',
    settings: [
      {
        name: 'showCount',
        title: 'Show Count',
        type: 'boolean',
        default: true,
      },
      {
        name: 'showPoints',
        title: 'Show Points',
        type: 'boolean',
        default: true,
      },
      {
        name: 'showRatio',
        title: 'Show Retro Ratio',
        type: 'boolean',
        default: true,
      },
    ],
  },
  'widget-unlocked': {
    name: 'widget-unlocked',
    title: 'Recent Unlocks',
    settings: [
      {
        name: 'count',
        title: 'Count',
        type: 'slider',
        default: 8,
        min: 0,
        max: 50,
      },
      {
        name: 'iconHeight',
        title: 'Icon Height',
        type: 'slider',
        default: 64,
        min: 12,
        max: 128,
      },
    ],
  },
  'widget-locked': {
    name: 'widget-locked',
    title: 'Next Locked Achievements',
    settings: [
      {
        name: 'count',
        title: 'Count',
        type: 'slider',
        default: 8,
        min: 0,
        max: 50,
      },
      {
        name: 'iconHeight',
        title: 'Icon Height',
        type: 'slider',
        default: 64,
        min: 12,
        max: 128,
      },
    ],
  },
  'widget-focused': {
    name: 'widget-focused',
    title: 'Focused Achievement',
    settings: [
      {
        name: 'marquee',
        title: 'Scroll Description',
        type: 'boolean',
        default: true,
      },
      {
        name: 'softPercent',
        title: 'Show Softcore Percent',
        type: 'boolean',
        default: true,
      },
      {
        name: 'hardPercent',
        title: 'Show Hardcore Percent',
        type: 'boolean',
        default: true,
      },
    ],
  },
  'widget-achievement-overlay': {
    name: 'widget-achievement-overlay',
    title: 'Achievement Overlay',
    settings: [
      {
        name: 'testAchieve',
        title: 'Show Test Overlay',
        type: 'boolean',
        default: false,
      },
      {
        name: 'shakeType',
        title: 'Animation',
        type: 'selection',
        default: 'shake-chunk',
        options: [
          { text: 'None', value: '' },
          { text: 'Basic', value: 'shake' },
          { text: 'Slow', value: 'shake-slow' },
          { text: 'Little', value: 'shake-little' },
          { text: 'Rotate', value: 'shake-rotate' },
          { text: 'Opacity', value: 'shake-opacity' },
          { text: 'Chunk', value: 'shake-chunk' },
        ],
      },
    ],
  },
  'widget-overlay': {
    name: 'widget-overlay',
    title: 'Static Overlay',
    settings: [
      {
        name: 'text',
        title: 'Text',
        type: 'textarea',
        default: '',
      },
    ],
  },
};
