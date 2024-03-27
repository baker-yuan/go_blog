export default (initialState: { currentUser?: API.CurrentUser | undefined }) => {
  const { currentUser } = initialState || {};
  return {
    canAdmin: currentUser && currentUser.access === 'admin',
  };
};
