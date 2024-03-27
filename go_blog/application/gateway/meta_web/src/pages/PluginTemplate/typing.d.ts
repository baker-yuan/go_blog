
declare namespace PluginTemplateModule {
  type Entity = {
    desc: string;
    labels: Record<string, string>;
    plugins: Record<string, any>;
  };

  type ResEntity = Entity & {
    id: string;
    update_time: string;
  };
}
