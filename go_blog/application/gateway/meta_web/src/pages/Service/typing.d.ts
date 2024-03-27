declare namespace ServiceModule {
  type Entity = {
    name: string;
    desc: string;
    hosts?: string[];
    upstream: any;
    upstream_id: string;
    labels: string;
    enable_websocket: boolean;
    plugins: Record<string, any>;
  };

  type ResponseBody = {
    id: string;
    plugins: Record<string, any>;
    upstream_id: string;
    upstream: Record<string, any>;
    name: string;
    desc: string;
    enable_websocket: boolean;
  };

  type Step1PassProps = {
    form: FormInstance;
    upstreamForm: FormInstance;
    disabled?: boolean;
    upstreamRef: any;
  };
}
