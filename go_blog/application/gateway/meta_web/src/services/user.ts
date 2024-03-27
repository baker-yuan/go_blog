import logo from '@/assets/logo.svg';

export async function queryCurrent(): Promise<API.CurrentUser> {
  return Promise.resolve({
    name: 'APISIX User',
    avatar: logo,
    userid: '00000001',
    access: 'admin',
  });
}
