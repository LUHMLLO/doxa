import { Show } from "solid-js";

let inputName = "";

async function CreateNewDevice(devicename) {
  if (devicename != "") {
    const response = await fetch(
      "https://doxamonitor-adminpanel.onrender.com/api/devices",
      {
        method: "POST",
        body: JSON.stringify({
          name: devicename,
          category: "63dc0dfd71febe988c1d4dce",
          owner: "63dc0e5671febe988c1d4df1",
        }),
        headers: {
          "Content-Type": "application/json; charset=UTF-8",
        },
      }
    );

    alert("created new device");
  } else {
    alert("Device name can't be blank");
  }
}

const addNewDevice = () => {
  return (
    <>
      <div class="p--16 bg-complementary radius-theme flex column gap--16">
        <input
          type="text"
          placeholder="Set device name"
          bind:value={inputName}
          class="bg-primary radius-theme"
        />
        <button
          class="clr-primary bg-dark radius-theme"
          onClick={() => CreateNewDevice(inputName)}
        >
          add new device
        </button>
      </div>
    </>
  );
};

export default addNewDevice;
