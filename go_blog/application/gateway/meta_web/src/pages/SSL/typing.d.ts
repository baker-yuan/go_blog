declare namespace SSLModule {
  type SSL = {
    sni: string[];
    cert: string;
    key: string;
  };

  type UploadPublicSuccessData = {
    cert: string;
    publicKeyList: UploadFile[];
  };

  type UploadPrivateSuccessData = {
    key: string;
    privateKeyList: UploadFile[];
  };

  type VerifyKeyPaireProps = {
    code: string;
    msg: string;
    data: {
      id: string;
      create_time: number;
      update_time: number;
      validity_start: number;
      validity_end: number;
      snis: string[];
      status: number;
    };
  };

  type ResponseBody = {
    id: string;
    cert: string;
    create_time: number;
    snis: string[];
    status: number;
    update_time: number;
    validity_start?: number;
    validity_end?: number;
  };
}
