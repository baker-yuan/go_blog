declare namespace ServerInfoModule {
  type Node = {
    id: string;
    last_report_time: integer;
    up_time: integer;
    boot_time: integer;
    etcd_version: string;
    hostname: string;
    version: string;
  };

  type DashboardInfo = {
    commit_hash: string;
    version: string;
  };
}
