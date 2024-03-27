declare namespace ProtoModule {
  type EditMode = 'create' | 'update';

  type ProtoData = {
    id: ?string;
    content: string;
    desc: string;
  };

  type ProtoDrawerProps = {
    protoData: ProtoData;
    setProtoData: (protoData: ProtoData) => void;
    visible: boolean;
    setVisible: (visible: boolean) => void;
    editMode: EditMode;
    refreshTable: () => void;
  };

  type ResponseBody = {
    id: string;
    create_time: number;
    update_time: number;
    desc: string;
    content: string;
  };
}
