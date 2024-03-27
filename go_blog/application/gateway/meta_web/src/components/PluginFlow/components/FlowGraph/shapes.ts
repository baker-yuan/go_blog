import { Graph } from '@antv/x6';

import defaultPluginImg from '../../../../../public/static/default-plugin.png';
import { DEFAULT_SHAPE_RECT_OPINIONS, FlowGraphShape } from '../../constants';

export const FlowChartRect = Graph.registerNode(FlowGraphShape.base, DEFAULT_SHAPE_RECT_OPINIONS);

export const FlowChartConditionRect = Graph.registerNode(FlowGraphShape.condition, {
  ...DEFAULT_SHAPE_RECT_OPINIONS,
  ports: {
    ...DEFAULT_SHAPE_RECT_OPINIONS.ports,
    items: [
      {
        group: 'top',
      },
      {
        group: 'right',
      },
      {
        group: 'bottom',
      },
    ],
  },
});

export const FlowChartStartRect = Graph.registerNode(FlowGraphShape.start, {
  ...DEFAULT_SHAPE_RECT_OPINIONS,
  ports: {
    ...DEFAULT_SHAPE_RECT_OPINIONS.ports,
    items: [
      {
        group: 'bottom',
      },
    ],
  },
});

export const FlowChartEndRect = Graph.registerNode(FlowGraphShape.end, {
  ...DEFAULT_SHAPE_RECT_OPINIONS,
  ports: {
    ...DEFAULT_SHAPE_RECT_OPINIONS.ports,
    items: [
      {
        group: 'top',
      },
    ],
  },
});

export const FlowChartPluginRect = Graph.registerNode(FlowGraphShape.plugin, {
  inherit: 'rect',
  width: 200,
  height: 60,
  attrs: {
    body: {
      stroke: '#5F95FF',
      strokeWidth: 1,
      fill: 'rgba(95,149,255,0.05)',
    },
    image: {
      'xlink:href': defaultPluginImg,
      width: 16,
      height: 16,
      x: 12,
      y: 12,
    },
    title: {
      text: 'Unknown Plugin',
      refX: 40,
      refY: 14,
      fill: 'rgba(0,0,0,0.85)',
      fontSize: 12,
      'text-anchor': 'start',
    },
    text: {
      text: '',
      refX: 40,
      refY: 38,
      fontSize: 12,
      fill: 'rgba(0,0,0,0.6)',
      'text-anchor': 'start',
    },
  },
  markup: [
    {
      tagName: 'rect',
      selector: 'body',
    },
    {
      tagName: 'image',
      selector: 'image',
    },
    {
      tagName: 'text',
      selector: 'title',
    },
    {
      tagName: 'text',
      selector: 'text',
    },
  ],
  ports: {
    groups: {
      top: {
        position: 'top',
        attrs: {
          circle: {
            r: 3,
            magnet: true,
            stroke: '#5F95FF',
            strokeWidth: 1,
            fill: '#fff',
            style: {
              visibility: 'hidden',
            },
          },
        },
      },
      bottom: {
        position: 'bottom',
        attrs: {
          circle: {
            r: 3,
            magnet: true,
            stroke: '#5F95FF',
            strokeWidth: 1,
            fill: '#fff',
            style: {
              visibility: 'hidden',
            },
          },
        },
      },
    },
    items: [
      {
        group: 'top',
      },
      {
        group: 'bottom',
      },
    ],
  },
});
