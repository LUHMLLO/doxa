<grid class="grid md:grid-cols--2 h-min--100vh">
  <column class="flex justify--center align--center dpt-secondary zindex--2">
    <section class="w--100">
      <container class="container block ">
        <slot />
      </container>
    </section>
  </column>

  <column
    class="none md:flex justify--center align--center zindex--1 bg-accent h--100"
    ><section>
      <container class="container block">
        <h3 class="text--center">Bienvenido a Doxatec</h3>
      </container>
    </section>
  </column>
</grid>
