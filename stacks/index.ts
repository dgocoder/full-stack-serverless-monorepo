import { App } from "sst/constructs";
import { PaymentsServiceStack } from "./services/payments/payment-stack";
import { UsersServiceStack } from "./services/users/user-stack";
import { SharedStack } from "./shared";

export const Stacks = (app: App) => {
  app.setDefaultFunctionProps({
    runtime: "go1.x",
  });
  app.stack(SharedStack).stack(UsersServiceStack).stack(PaymentsServiceStack);
};
