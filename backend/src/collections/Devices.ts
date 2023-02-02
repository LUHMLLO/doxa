import { CollectionConfig } from "payload/types";

const Devices: CollectionConfig = {
  slug: "devices",
  admin: {
    useAsTitle: "name",
    defaultColumns: ["id", "name", "category", "owner"],
    group: "Devices",
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
      name: "category",
      type: "relationship",
      relationTo: "categories",
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
