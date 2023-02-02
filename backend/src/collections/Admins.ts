import { CollectionConfig } from "payload/types";

const Admins: CollectionConfig = {
  slug: "admins",
  auth: true,
  admin: {
    useAsTitle: "name",
    defaultColumns: ["id", "name", "email"],
    group: "Users",
  },
  access: {
    read: () => true,
  },

  fields: [
    {
      name: "name",
      type: "text",
    },
  ],
};

export default Admins;
