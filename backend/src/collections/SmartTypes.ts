import { CollectionConfig } from "payload/types";

const SmartTypes: CollectionConfig = {
  slug: "types",
  admin: {
    useAsTitle: "name",
    defaultColumns: ["id", "name",],
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
  ],
};

export default SmartTypes;
