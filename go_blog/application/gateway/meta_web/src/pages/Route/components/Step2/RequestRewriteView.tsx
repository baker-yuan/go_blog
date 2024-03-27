import React, { useEffect, useState } from 'react';

import UpstreamForm from '@/components/Upstream';
import { fetchUpstreamList } from '@/components/Upstream/service';

const RequestRewriteView: React.FC<RouteModule.Step2PassProps> = ({
  form,
  upstreamRef,
  disabled,
  hasServiceId = false,
}) => {
  const [list, setList] = useState<UpstreamComponent.ResponseData[]>([]);
  useEffect(() => {
    fetchUpstreamList().then(({ data }) => setList(data as UpstreamComponent.ResponseData[]));
  }, []);
  return (
    <UpstreamForm
      ref={upstreamRef}
      form={form}
      disabled={disabled}
      list={list}
      showSelector
      required={!hasServiceId}
      key={1}
    />
  );
};

export default RequestRewriteView;
