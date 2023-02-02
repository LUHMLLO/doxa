import { Show } from "solid-js";

async function CreateNewDevice() {
  const response = await fetch(
    "http://localhost:5000/api/devices",
    {
      method: "POST",
      body: JSON.stringify({
        name: "Post test fridge",
      }),
      headers: {
        "Access-Control-Allow-Origin": "*",
        "Content-Type": "application/json; charset=UTF-8",
        Accept: "application/json",
      },
    }
  );
  console.log("created new device");
}

const AddNewDevice = (props) => {
  return (
    <>
      <button
        class="clr-primary bg-dark radius-theme"
        onClick={CreateNewDevice}
      >
        add new device
      </button>
    </>
  );
};

export default AddNewDevice;
