import pb from "./pocketbase";

let devices = [];

const devicesList = await pb.collection("devices").getList(1, 20, {
  sort: "created",
  expand: "user",
});

devices = devicesList.items;

export { devicesList };
