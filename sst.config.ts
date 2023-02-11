import { SSTConfig } from "sst";
import { Stacks } from "./stacks";

export default {
  config(_input) {
    return {
      name: "my-sst-app",
      region: "us-east-1",
    };
  },
  stacks: (app) => Stacks(app),
} satisfies SSTConfig;
