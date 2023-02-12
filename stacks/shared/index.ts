import type { Stack } from 'sst/constructs';

import { TopicsStack } from './topics';

export * from './topics';

type SharedContext = {
  stack: Stack;
};

export const SharedStack = ({ stack }: SharedContext) => {
  return {
    ...TopicsStack({ stack }),
  };
};
