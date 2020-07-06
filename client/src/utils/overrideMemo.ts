import React, { PropsWithChildren, SFC } from 'react';

/**
 * 重写React.memo方法使其能在devtool中显示名字
 */

const originMemo = React.memo;

React.memo = function overriddenMemo<P extends object>(
  Component: SFC<P>,
  propsAreEqual?: (
    prevProps: Readonly<PropsWithChildren<P>>,
    nextProps: Readonly<PropsWithChildren<P>>
  ) => boolean
) {
  const memoized = originMemo(Component, propsAreEqual);

  Object.defineProperty(memoized, 'displayName', {
    set(v) {
      Component.displayName = v;
    },
  });

  return memoized;
} as any;
