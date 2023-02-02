import { buildConfig } from "payload/config";
import path from "path";
import Admins from "./collections/Admins";
import Clients from "./collections/Clients";
import Devices from "./collections/Devices";
import Categories from "./collections/Categories";

export default buildConfig({
  serverURL: process.env.PAYLOAD_PUBLIC_BASE_DNS,
  rateLimit: {
    trustProxy: true,
  },
  admin: {
    user: Admins.slug,
  },
  collections: [Admins, Clients, Devices, Categories],
  typescript: {
    outputFile: path.resolve(__dirname, "payload-types.ts"),
  },
  graphQL: {
    schemaOutputFile: path.resolve(__dirname, "generated-schema.graphql"),
  },
});
