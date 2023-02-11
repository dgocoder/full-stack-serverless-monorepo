import { Api } from "sst/constructs";
import { StackContext } from "sst/constructs";
import { DDBStack } from "../../macros/ddb/index";

export const UsersServiceStack = ({ stack }: StackContext) => {
  const { table } = DDBStack({
    stack,
    tableName: "users",
    consumers: [
      {
        name: "dbstream",
        consumer: {
          function: "services/users/cmd/lambdas/db-stream/dbstream.go",
        },
      },
    ],
  });

  const api = new Api(stack, "users-svc");
  api.addRoutes(stack, {
    "GET /{id}": {
      function: {
        handler: "services/users/cmd/lambdas/get-user/get-user.main.go",
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
