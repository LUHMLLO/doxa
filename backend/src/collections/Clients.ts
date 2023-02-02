import { CollectionConfig } from "payload/types";

const Clients: CollectionConfig = {
  slug: "clients",
  auth: true,
  admin: {
    useAsTitle: "name",
    defaultColumns: ["id", "name", "email", "phone"],
    group: "Users",
  },
  access: {
    create: () => true,
    read: () => true,
    update: () => true,
    delete: () => true,
  },
  fields: [
    {
      name: "name",
      type: "text",
      localized: true,
    },
    {
      name: "phone",
      type: "number",
      required: true,
    },
    {
      name: "secret",
      type: "text",
      required: true,
    },
  ],
};

export default Clients;
