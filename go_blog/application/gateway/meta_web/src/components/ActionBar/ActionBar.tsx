import { Button, Col, Row } from 'antd';
import type { CSSProperties } from 'react';
import React from 'react';
import { useIntl } from 'umi';

type Props = {
  step: number;
  lastStep: number;
  onChange: (nextStep: number) => void;
  withResultView?: boolean;
  loading?: boolean;
};

const style: CSSProperties = {
  position: 'fixed',
  bottom: 0,
  right: 10,
  margin: '-24px -24px 0',
  backgroundColor: '#fff',
  padding: '6px 36px',
  borderTop: '1px solid #ebecec',
  width: '100%',
};

const ActionBar: React.FC<Props> = ({ step, lastStep, onChange, withResultView, loading }) => {
  const { formatMessage } = useIntl();

  if (step > lastStep && !withResultView) {
    onChange(lastStep);
    return null;
  }

  // resultView should not show actionbar
  if (step > lastStep && withResultView) {
    return null;
  }

  return (
    <div style={style}>
      <Row gutter={10} justify="end">
        <Col>
          <Button type="primary" onClick={() => onChange(step - 1)} disabled={step === 1}>
            {formatMessage({ id: 'component.actionbar.button.preStep' })}
          </Button>
        </Col>
        <Col>
          <Button type="primary" onClick={() => onChange(step + 1)} loading={loading}>
            {step < lastStep
              ? formatMessage({ id: 'component.actionbar.button.nextStep' })
              : formatMessage({ id: 'component.global.submit' })}
          </Button>
        </Col>
      </Row>
    </div>
  );
};

export default ActionBar;
