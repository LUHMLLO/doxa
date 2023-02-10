import "./app.css";
import App from "./App.svelte";
import "./lib/hooks/authstore";

const app = new App({
  target: document.getElementById("app"),
});

export default app;
