export const transformData = (data: ServiceModule.Entity): ServiceModule.Entity => {
  const newData = data;

  if (newData.hosts) {
    newData.hosts = newData.hosts.filter((item) => {
      return !!item;
    });

    if (newData.hosts.length <= 0) {
      delete newData.hosts;
    }
  }

  return newData;
};
