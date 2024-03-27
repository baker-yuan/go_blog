import { PlusOutlined } from '@ant-design/icons';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import type { ActionType, ProColumns } from '@ant-design/pro-table';
import ProTable from '@ant-design/pro-table';
import { Button, notification, Popconfirm, Tag } from 'antd';
import React, { useRef, useState } from 'react';
import { history, useIntl } from 'umi';

import { timestampToLocaleString } from '@/helpers';
import usePagination from '@/hooks/usePagination';
import { fetchList, remove as removeSSL } from '@/pages/SSL/service';

const Page: React.FC = () => {
  const tableRef = useRef<ActionType>();
  const { formatMessage } = useIntl();
  const { paginationConfig, savePageList, checkPageList } = usePagination();
  const [deleteLoading, setDeleteLoading] = useState('');

  const columns: ProColumns<SSLModule.ResponseBody>[] = [
    {
      title: 'SNI',
      dataIndex: 'sni',
      render: (_, record) => {
        return (record.snis || []).map((sni) => (
          <Tag color="geekblue" key={sni}>
            {sni}
          </Tag>
        ));
      },
    },
    {
      title: formatMessage({ id: 'page.ssl.list.expirationTime' }),
      dataIndex: 'validity_end',
      hideInSearch: true,
      render: (text) => timestampToLocaleString(text as number),
    },
    {
      title: formatMessage({ id: 'component.global.updateTime' }),
      dataIndex: 'update_time',
      hideInSearch: true,
      render: (text) => timestampToLocaleString(text as number),
    },
    {
      title: formatMessage({ id: 'component.global.operation' }),
      valueType: 'option',
      render: (_, record) => (
        <>
          <Button
            type="primary"
            onClick={() => {
              history.push(`/ssl/${record.id}/edit`);
            }}
            style={{ marginRight: 10 }}
          >
            {formatMessage({ id: 'component.global.edit' })}
          </Button>
          <Popconfirm
            title={formatMessage({ id: 'component.ssl.removeSSLItemModalContent' })}
            onConfirm={() => {
              setDeleteLoading(record.id);
              removeSSL(record.id)
                .then(() => {
                  notification.success({
                    message: formatMessage({ id: 'component.ssl.removeSSLSuccess' }),
                  });
                  requestAnimationFrame(() => checkPageList(tableRef));
                })
                .finally(() => {
                  setDeleteLoading('');
                });
            }}
            cancelText={formatMessage({ id: 'component.global.cancel' })}
            okText={formatMessage({ id: 'component.global.confirm' })}
          >
            <Button type="primary" danger loading={record.id === deleteLoading}>
              {formatMessage({ id: 'component.global.delete' })}
            </Button>
          </Popconfirm>
        </>
      ),
    },
    {
      title: formatMessage({ id: 'page.ssl.list.periodOfValidity' }),
      dataIndex: 'expire_range',
      hideInTable: true,
      hideInSearch: true,
    },
  ];

  return (
    <PageHeaderWrapper
      title={formatMessage({ id: 'page.ssl.list' })}
      content={formatMessage({ id: 'component.ssl.description' })}
    >
      <ProTable<SSLModule.ResponseBody>
        rowKey="id"
        columns={columns}
        actionRef={tableRef}
        request={fetchList}
        pagination={{
          onChange: (page, pageSize?) => savePageList(page, pageSize),
          pageSize: paginationConfig.pageSize,
          current: paginationConfig.current,
        }}
        search={{
          searchText: formatMessage({ id: 'component.global.search' }),
          resetText: formatMessage({ id: 'component.global.reset' }),
        }}
        toolBarRender={() => [
          <Button type="primary" onClick={() => history.push(`/ssl/create`)}>
            <PlusOutlined />
            {formatMessage({ id: 'component.global.create' })}
          </Button>,
        ]}
      />
    </PageHeaderWrapper>
  );
};

export default Page;
