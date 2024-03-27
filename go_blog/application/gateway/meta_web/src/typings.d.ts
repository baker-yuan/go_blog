declare module 'slash2';
declare module '*.css';
declare module '*.less';
declare module '*.scss';
declare module '*.sass';
declare module '*.svg';
declare module '*.png';
declare module '*.jpg';
declare module '*.jpeg';
declare module '*.gif';
declare module '*.bmp';
declare module '*.tiff';
declare module 'omit.js';

// google analytics interface
type GAFieldsObject = {
  eventCategory: string;
  eventAction: string;
  eventLabel?: string;
  eventValue?: number;
  nonInteraction?: boolean;
};
type Window = {
  ga: (
    command: 'send',
    hitType: 'event' | 'pageview',
    fieldsObject: GAFieldsObject | string,
  ) => void;
  reloadAuthorized: () => void;
  codemirror: Record<string, any>;
};

declare let ga: Function;

// preview.pro.ant.design only do not use in your production ;
declare let ANT_DESIGN_PRO_ONLY_DO_NOT_USE_IN_YOUR_PRODUCTION: 'site' | undefined;

declare const REACT_APP_ENV: 'test' | 'dev' | 'pre' | false;

type PageMode = 'CREATE' | 'EDIT' | 'VIEW';

type Res<T> = {
  code: number;
  message: string;
  request_id: string;
  data: T;
};

type ResListData<T> = {
  rows: T[];
  total_size: number;
};

type HttpMethod =
  | 'GET'
  | 'POST'
  | 'DELETE'
  | 'PUT'
  | 'OPTIONS'
  | 'HEAD'
  | 'PATCH'
  | 'CONNECT'
  | 'TRACE'
  | 'PURGE';

type ResponseLabelList = Record<string, string>[];

type LabelList = Record<string, string[]>;
