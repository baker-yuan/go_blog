declare namespace ConsumerModule {
  type Entity = {
    username: string;
    desc: string;
    plugins: Record<string, any>;
  };

  type ResEntity = Entity & {
    id: string;
    update_time: string;
  };
}
