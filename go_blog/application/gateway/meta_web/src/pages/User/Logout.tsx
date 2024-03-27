import type React from 'react';

import { getUrlQuery } from '@/helpers';
import LoginMethodPassword from '@/pages/User/components/LoginMethodPassword';
import type { UserModule } from '@/pages/User/typing';

/**
 * Login Methods List
 */
const loginMethods: UserModule.LoginMethod[] = [LoginMethodPassword];

/**
 * User Logout Page
 * @constructor
 */
const Page: React.FC = () => {
  // run all logout method
  loginMethods.forEach((item) => {
    item.logout();
  });

  const redirect = getUrlQuery('redirect');
  window.location.href = `/user/login${redirect ? `?redirect=${redirect}` : ''}`;

  return null;
};

export default Page;
