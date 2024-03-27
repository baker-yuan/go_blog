
/**
 * Schema: https://github.com/apache/apisix/blob/master/apisix/schema_def.lua#L335
 */
declare namespace UpstreamComponent {
  type ActiveCheck = {};

  type PassiveCheck = {};

  type TLS = {
    client_cert: string;
    client_key: string;
  };

  type Node = {
    host: string;
    port: number;
    weight: number;
    priority?: number;
  };

  type SubmitNode = Record<string, number>;

  type Timeout = {
    connect: number;
    send: number;
    read: number;
  };

  type ResponseData = {
    nodes?: SubmitNode | Node[];
    retries?: number;
    timeout?: Timeout;
    tls?: TLS;
    type?: string;
    checks?: {
      active?: ActiveCheck;
      passive?: PassiveCheck;
    };
    hash_on?: string;
    key?: string;
    scheme?: string;
    discovery_type?: string;
    pass_host?: string;
    upstream_host?: string;
    name?: string;
    desc?: string;
    service_name?: string;
    id?: string;
    upstream_id?: string;
    upstream_type?: string;

    // NOTE: custom field
    custom?: Record<string, any>;
  };
}
