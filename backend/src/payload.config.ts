import { buildConfig } from "payload/config";
import path from "path";
import Admins from "./collections/Admins";
import Clients from "./collections/Clients";
import Devices from "./collections/Devices";
import SmartTypes from "./collections/SmartTypes";

export default buildConfig({
  serverURL: "http://localhost:5000",
  admin: {
    user: Admins.slug,
  },
  collections: [Admins, Clients, Devices, SmartTypes],
  typescript: {
    outputFile: path.resolve(__dirname, "payload-types.ts"),
  },
  graphQL: {
    schemaOutputFile: path.resolve(__dirname, "generated-schema.graphql"),
  },
});
