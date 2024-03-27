
declare namespace PluginModule {
  type TransformedPlugin = {
    id: string;
    name: string;
    value: Record<string, any>;
  };

  type GlobalRule = {
    id: string;
    plugins: Record<string, any>;
  };
}
