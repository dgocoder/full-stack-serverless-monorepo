import { Api, type StackContext, use } from 'sst/constructs';

import { DDBStack } from '../../macros/ddb/index';
import { SharedStack } from '../../shared';

export const UsersServiceStack = ({ stack }: StackContext) => {
  const { userCreated } = use(SharedStack);
  const { table } = DDBStack({
    stack,
    tableName: 'users',
    consumers: [
      {
        name: 'dbstream',
        consumer: {
          function: {
            handler: 'services/users/cmd/lambdas/db-stream/dbstream.go',
            environment: {
              USER_CREATED_TOPIC_ARN: userCreated.topicArn,
            },
            permissions: [userCreated],
          },
        },
      },
    ],
  });

  const api = new Api(stack, 'users-svc');
  api.addRoutes(stack, {
    'GET /{id}': {
      function: {
        handler: 'services/users/cmd/lambdas/get-user/get-user.main.go',
        environment: {
          USERS_TABLE_NAME: table.tableName,
        },
        permissions: [table],
      },
    },
  });

  stack.addOutputs({
    UsersApiEndpoint: api.url,
  });
};
