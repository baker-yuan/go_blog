
export const removeBtnStyle = {
  display: 'flex',
  alignItems: 'center',
};

export enum AlgorithmEnum {
  chash = 'chash',
  roundrobin = 'roundrobin',
  ewma = 'ewma',
  least_conn = 'least_conn',
}

export enum HashOnEnum {
  vars = 'vars',
  header = 'header',
  cookie = 'cookie',
  consumer = 'consumer',
  vars_combinations = 'vars_combinations',
}

export enum CommonHashKeyEnum {
  remote_addr = 'remote_addr',
  host = 'host',
  uri = 'uri',
  server_name = 'server_name',
  server_addr = 'server_addr',
  request_uri = 'request_uri',
  query_string = 'query_string',
  remote_port = 'remote_port',
  hostname = 'hostname',
  arg_id = 'arg_id',
}

export enum SchemeEnum {
  grpc = 'grpc',
  grpcs = 'grpcs',
  http = 'http',
  https = 'https',
}

export enum PassHostEnum {
  pass = 'pass',
  node = 'node',
  rewrite = 'rewrite',
}
