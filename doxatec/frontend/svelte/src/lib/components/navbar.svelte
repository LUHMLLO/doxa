<script lang="ts">
  const Theme = {
    common:
      "flex row gap--16 align--center px--16 py--8 mx--auto bg-primary clr-secondary dpt-secondary ps--absolute zindex--12 left right",
    mobile: "justify--space-between bottom mb--16 w--92 radius-theme",
    desktop: "",
  };
</script>

<nav class={`${Theme.common} ${Theme.mobile} ${Theme.desktop}`}>
  <a href="/" class="clr-secondary">
    <icon class="icon size--64">
      <i class="iconoir-network-alt" />
    </icon>
  </a>
  <a href="/" class="clr-secondary">
    <icon class="icon size--64">
      <i class="iconoir-okrs" />
    </icon>
  </a>
  <a href="/" class="clr-secondary">
    <icon class="icon size--64">
      <i class="iconoir-flash" />
    </icon>
  </a>
  <a href="/" class="clr-secondary">
    <icon class="icon size--64">
      <i class="iconoir-user-circle" />
    </icon>
  </a>
  <a href="/" class="clr-secondary">
    <icon class="icon size--64">
      <i class="iconoir-more-horiz" />
    </icon>
  </a>
</nav>
