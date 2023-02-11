import { App } from "sst/constructs";
import { UsersServiceStack } from "./services/users/user-stack";

export const Stacks = (app: App) => {
  app.setDefaultFunctionProps({
    runtime: "go1.x",
  });
  app.stack(UsersServiceStack);
};
