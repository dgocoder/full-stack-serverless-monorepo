import { Stack, Queue, Topic, FunctionDefinition } from "sst/constructs";
import { Duration } from "aws-cdk-lib";

type TopicQueueLambdaStackContext = {
  stack: Stack;
  name: string;
  topics: Topic[];
  lambdaFunc: FunctionDefinition;
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
      [topicQueue.queueName]: topicQueue,
    });
  });

  return {
    dlq,
    topicQueue,
  };
};
