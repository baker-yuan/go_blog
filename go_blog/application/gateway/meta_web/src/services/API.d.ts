declare namespace API {
  export type CurrentUser = {
    avatar?: string;
    name?: string;
    userid?: string;
    access?: 'user' | 'guest' | 'admin';
  };

  export type LoginStateType = {
    status?: 'ok' | 'error';
    type?: string;
  };

  export type NoticeIconData = {
    avatar?: string | React.ReactNode;
    title?: React.ReactNode;
    description?: React.ReactNode;
    datetime?: React.ReactNode;
    extra?: React.ReactNode;
    style?: React.CSSProperties;
    key?: string | number;
    read?: boolean;
  };
}
