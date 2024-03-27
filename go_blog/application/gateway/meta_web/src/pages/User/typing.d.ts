
import type React from 'react';

declare namespace UserModule {
  type LoginMethod = {
    id: string;
    name: string;
    render: () => React.ReactElement;
    getData: () => LoginData;
    checkData: () => Promise<boolean>;
    submit: (data) => Promise<LoginResponse>;
    logout: () => void;
  };

  type LoginData = {
    [string]: string;
  };

  type LoginResponse = {
    status: boolean;
    message: string;
    data: {
      [string]: any;
    };
  };
}
