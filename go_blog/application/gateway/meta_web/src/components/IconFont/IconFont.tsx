import React from 'react';

type Props = {
  name: string;
};

/**
 * Icon Font
 * https://www.iconfont.cn/help/detail?helptype=code
 */
const IconFont: React.FC<Props> = ({ name }) => (
  <svg className="icon" aria-hidden="true">
    <use xlinkHref={`#${name}`} />
  </svg>
);

export default IconFont;
