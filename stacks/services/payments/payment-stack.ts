import { use } from "sst/constructs";
import { StackContext } from "sst/constructs";
import { TopicQueueLambdaStack } from "../../macros";
import { SharedStack } from "../../shared";

export const PaymentsServiceStack = ({ stack }: StackContext) => {
  const { userCreated } = use(SharedStack);
  TopicQueueLambdaStack({
    stack,
    name: "user-payment",
    topics: [userCreated],
    lambdaFunc: {
      handler:
        "services/payments/cmd/lambdas/user-payment/user-payment.main.go",
    },
  });
};
