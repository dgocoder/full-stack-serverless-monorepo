import { type Stack, Topic } from 'sst/constructs';

type TopicsContext = {
  stack: Stack;
};

export const TopicsStack = ({ stack }: TopicsContext) => {
  const userCreated = new Topic(stack, 'user-created');

  return {
    userCreated,
  };
};
