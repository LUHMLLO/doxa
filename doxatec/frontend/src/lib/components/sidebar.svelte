<script lang="ts">
  import { Link } from "svelte-navigator";
</script>

<nav>
  <row>
    <Link to="/">
      <icon class="icon">
        <icon class="iconoir-report-columns" />
      </icon>
      <small>Dashboard</small>
    </Link>
    <Link to="/devices">
      <icon class="icon">
        <icon class="iconoir-fridge" />
      </icon>
      <small>Neveras</small>
    </Link>
  </row>
  <nav>
    <Link to="/">
      <icon class="icon">
        <icon class="iconoir-sidebar-expand" />
      </icon>
      <small> Dashboard</small>
    </Link>
    <Link to="/devices">
      <icon class="icon">
        <icon class="iconoir-sidebar-expand" />
      </icon>
      <small>Neveras</small>
    </Link>
  </nav>
</nav>
