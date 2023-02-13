import { Duration } from 'aws-cdk-lib';
import { SqsSubscription } from 'aws-cdk-lib/aws-sns-subscriptions';
import { type FunctionDefinition, Queue, type Stack, type Topic } from 'sst/constructs';

type TopicQueueLambdaStackContext = {
  lambdaFunc: FunctionDefinition;
  name: string;
  stack: Stack;
  topics: Topic[];
};

export const TopicQueueLambdaStack = ({
  stack,
  name,
  topics,
  lambdaFunc,
}: TopicQueueLambdaStackContext) => {
  const dlq = new Queue(stack, `${name}-dlq`, {
    cdk: {
      queue: {
        retentionPeriod: Duration.days(14),
      },
    },
  });
  const topicQueue = new Queue(stack, `${name}-queue`, {
    cdk: {
      queue: {
        deadLetterQueue: {
          queue: dlq.cdk.queue,
          maxReceiveCount: 1,
        },
      },
    },
    consumer: {
      function: lambdaFunc,
    },
  });

  topics.forEach((topic) => {
    topic.addSubscribers(stack, {
      [`${name}-queue`]: topicQueue,
    });
  });

  return {
    dlq,
    topicQueue,
  };
};
