import { CollectionConfig } from "payload/types";

const Admins: CollectionConfig = {
  slug: "admins",
  auth: {
    useAPIKey: true,
  },
  admin: {
    useAsTitle: "email",
    group: "Admin",
  },
  access: {
    read: () => true,
  },

  fields: [
    {
      name: "name",
      type: "text",
      saveToJWT: true,
    },
  ],
};

export default Admins;
