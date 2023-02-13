import { SSTConfig } from "sst";
import { Stacks } from "./stacks";

export default {
  config(_input) {
    return {
      name: "monorepo",
      region: "us-east-2",
    };
  },
  stacks: (app) => Stacks(app),
} satisfies SSTConfig;
