import {
  Table,
  Stack,
  TableGlobalIndexProps,
  TableConsumerProps,
  Queue,
} from "sst/constructs";
import { StartingPosition } from "aws-cdk-lib/aws-lambda";
import { RemovalPolicy } from "aws-cdk-lib";
import { SqsDestination } from "aws-cdk-lib/aws-lambda-destinations";

type DDBStack = {
  stack: Stack;
  tableName: string;
  globalIndexes?: { pk: string; sk: string }[];
  consumers?: { name: string; consumer: TableConsumerProps }[];
};

export const DDBStack = ({
  stack,
  tableName,
  globalIndexes,
  consumers,
}: DDBStack) => {
  const table = new Table(stack, tableName, {
    fields: {
      pk: "string",
      sk: "string",
    },
    primaryIndex: { partitionKey: "pk", sortKey: "sk" },
    globalIndexes: globalIndexes?.reduce((acc, cv, index) => {
      acc[`gsi${index + 1}`] = { partitionKey: cv.pk, sortKey: cv.sk };
      return acc;
    }, {} as Record<string, TableGlobalIndexProps>),
    stream: !!consumers?.length,
    cdk: {
      table: {
        removalPolicy:
          stack.stage !== "prod" ? RemovalPolicy.DESTROY : RemovalPolicy.RETAIN,
      },
    },
  });

  // Dynamically add table lambda consumers
  consumers?.forEach(({ consumer, name }) => {
    const dlq = new Queue(stack, "DDB-DLQ");
    table.addConsumers(stack, {
      [name]: {
        ...consumer,
        cdk: {
          ...consumer.cdk,
          eventSource: {
            startingPosition: StartingPosition.LATEST,
            bisectBatchOnError: true,
            onFailure: new SqsDestination(dlq.cdk.queue),
            retryAttempts: 1,
          },
        },
      },
    });
  });

  stack.addOutputs({
    [`${tableName}Table`]: table.tableName,
  });

  return {
    table,
  };
};
