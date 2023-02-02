import { CollectionConfig } from "payload/types";

const Devices: CollectionConfig = {
  slug: "devices",
  admin: {
    useAsTitle: "name",
    defaultColumns: ["id", "name", "type", "owner"],
    group: "Products",
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
