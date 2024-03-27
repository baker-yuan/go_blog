import { Button, Result } from 'antd';
import React from 'react';
import { history, useIntl } from 'umi';

type Props = {
  createNew: () => void;
};

const ResultView: React.FC<Props> = (props) => {
  const { formatMessage } = useIntl();
  return (
    <Result
      status="success"
      title={`${formatMessage({ id: 'component.global.submit' })} ${formatMessage({
        id: 'component.status.success',
      })}`}
      extra={[
        <Button type="primary" key="goto-list" onClick={() => history.replace('/routes/list')}>
          {formatMessage({ id: 'page.route.button.returnList' })}
        </Button>,
        <Button key="create-new" onClick={() => props.createNew()}>
          {`${formatMessage({ id: 'component.global.create' })} ${formatMessage({
            id: 'menu.routes',
          })}`}
        </Button>,
      ]}
    />
  );
};

export default ResultView;
