<script lang="ts">
  import WindowWrapper from "./lib/components/window.svelte";

  import SignUp from "./pages/signup.svelte";
  import LogIn from "./pages/login.svelte";
  import Account from "./pages/account.svelte";
</script>

<main class="bg-grayscale--5 flex column gap--16 p--16">
  <WindowWrapper>
    <SignUp />
  </WindowWrapper>
  <WindowWrapper>
    <LogIn />
  </WindowWrapper>
  <WindowWrapper>
    <Account />
  </WindowWrapper>
</main>
