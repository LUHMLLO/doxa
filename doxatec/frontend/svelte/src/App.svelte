<script lang="ts">
  import { Router, Route } from "svelte-navigator";
  import BaseLayout from "./lib/layouts/base.svelte";

  import SignUp from "./pages/signup.svelte";
  import LogIn from "./pages/login.svelte";
  import Account from "./pages/account.svelte";
  import Dashboard from "./pages/dashboard.svelte";
  import DeviceDetails from "./pages/deviceDetails.svelte";
  import Devices from "./pages/devices.svelte";
</script>

<Router>
  <Route path="signup">
    <SignUp />
  </Route>

  <Route path="login">
    <LogIn />
  </Route>

  <Route path="account">
    <Account layout={BaseLayout}/>
  </Route>

  <Route path="devices">
    <Devices layout={BaseLayout}/>
  </Route>

  <Route>
    <Dashboard layout={BaseLayout}/>
  </Route>

  <Route path="device/:id" let:params>
    <DeviceDetails layout={BaseLayout} id={params.id} />
  </Route>
</Router>
