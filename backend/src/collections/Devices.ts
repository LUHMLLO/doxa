import { CollectionConfig } from "payload/types";

const Devices: CollectionConfig = {
  slug: "devices",
  admin: {
    useAsTitle: "name",
    defaultColumns: ["id", "name", "type", "owner"],
    group: "Customers",
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
      name: "type",
      type: "relationship",
      relationTo: "types",
      hasMany: false,
    },
    {
      name: "owner",
      type: "relationship",
      relationTo: "clients",
      hasMany: false,
    },
  ],
};

export default Devices;
