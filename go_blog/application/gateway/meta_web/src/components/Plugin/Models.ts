
/**
 * Model List of PluginType of Monaco editor
 */

import { editor, Uri } from 'monaco-editor';

import * as modelCode from './modelCode';

/**
 * Model type is authentication as fllows:
 */
export const authzcasbinModel = editor.createModel(
  modelCode.authzcasbin,
  'json',
  Uri.parse('file:authz-casbin'),
);

export const authzkeycloakModel = editor.createModel(
  modelCode.authzkeycloak,
  'json',
  Uri.parse('file:authz-keycloak'),
);

export const forwardauthModel = editor.createModel(
  modelCode.forwardauth,
  'json',
  Uri.parse('file:forward-auth'),
);

export const opaModel = editor.createModel(modelCode.opa, 'json', Uri.parse('file:opa'));

export const openidconnectModel = editor.createModel(
  modelCode.openidconnect,
  'json',
  Uri.parse('file:openid-connect'),
);

/**
 * Model type is security as fllows:
 */

export const csrfModel = editor.createModel(modelCode.csrf, 'json', Uri.parse('file:csrf'));

export const iprestrictionModel = editor.createModel(
  modelCode.iprestriction,
  'json',
  Uri.parse('file:ip-restriction'),
);

export const uarestrictionModel = editor.createModel(
  modelCode.uarestriction,
  'json',
  Uri.parse('file:ua-restriction'),
);

export const uriblockerModel = editor.createModel(
  modelCode.uriblocker,
  'json',
  Uri.parse('file:uri-blocker'),
);

/**
 * Model type is traffic as fllows:
 */

export const clientcontrolModel = editor.createModel(
  modelCode.clientcontrol,
  'json',
  Uri.parse('file:client-control'),
);

export const trafficsplitModel = editor.createModel(
  modelCode.trafficsplit,
  'json',
  Uri.parse('file:traffic-split'),
);

/**
 * Model type is serverless as fllows:
 */

export const awslambdaModel = editor.createModel(
  modelCode.awslambda,
  'json',
  Uri.parse('file:aws-lambda'),
);

export const azurefunctionsModel = editor.createModel(
  modelCode.azurefunctions,
  'json',
  Uri.parse('file:azure-functions'),
);

export const openwhiskModel = editor.createModel(
  modelCode.openwhisk,
  'json',
  Uri.parse('file:openwhisk'),
);

/**
 * Model type is observability as fllows:
 */

export const clickhouseloggerModel = editor.createModel(
  modelCode.clickhouselogger,
  'json',
  Uri.parse('file:clickhouse-logger'),
);

export const fileloggerModel = editor.createModel(
  modelCode.filelogger,
  'json',
  Uri.parse('file:file-logger'),
);

export const googlecloudloggingModel = editor.createModel(
  modelCode.googlecloudlogging,
  'json',
  Uri.parse('file:google-cloud-logging'),
);

export const httploggerModel = editor.createModel(
  modelCode.httplogger,
  'json',
  Uri.parse('file:http-logger'),
);

export const kafkaloggerModel = editor.createModel(
  modelCode.kafkalogger,
  'json',
  Uri.parse('file:kafka-logger'),
);

export const logglyModel = editor.createModel(modelCode.loggly, 'json', Uri.parse('file:loggly'));

export const rocketmqloggerModel = editor.createModel(
  modelCode.rocketmqlogger,
  'json',
  Uri.parse('file:rocketmq-logger'),
);

export const skywalkingModel = editor.createModel(
  modelCode.skywalking,
  'json',
  Uri.parse('file:sky-walking'),
);

export const slsloggerModel = editor.createModel(
  modelCode.slslogger,
  'json',
  Uri.parse('file:sls-logger'),
);

export const splunkhecloggingModel = editor.createModel(
  modelCode.splunkheclogging,
  'json',
  Uri.parse('file:splunk-hec-logging'),
);

export const syslogModel = editor.createModel(modelCode.syslog, 'json', Uri.parse('file:syslog'));

export const tcploggerModel = editor.createModel(
  modelCode.tcplogger,
  'json',
  Uri.parse('file:tcp-logger'),
);

export const zipkinModel = editor.createModel(modelCode.zipkin, 'json', Uri.parse('file:zipkin'));

/**
 * Model type is other as fllows:
 */

export const extpluginprereqModel = editor.createModel(
  modelCode.extpluginprereq,
  'json',
  Uri.parse('file:ext-plugin-pre-req'),
);

export const realipModel = editor.createModel(modelCode.realip, 'json', Uri.parse('file:real-ip'));
