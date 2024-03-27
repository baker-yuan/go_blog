declare namespace PluginComponent {
  type Data = Record<string, any>;

  type Schema = '' | 'route' | 'consumer';

  type Meta = {
    name: string;
    priority: number;
    schema: Record<string, any>;
    type: string;
    originType: string;
    version: number;
    consumer_schema?: Record<string, any>;
    hidden?: boolean;
  };

  type ReferPage = '' | 'route' | 'consumer' | 'service' | 'plugin';

  type PluginDetailValues = {
    formData: any;
    monacoData: any;
    shouldDelete?: boolean;
  };

  type MonacoLanguage = 'JSON' | 'YAML' | 'Form';
}
