import {} from "solid-js";

const Footer = (props) => {
  return (
    <>
      <footer class="clr-tertiary bg-complementary flex align--center py--16 noshrink">
        <container class="container block radius-theme">
          <grid class="grid grid-cols--1 md:grid-cols--3 gap--32">
            <row class="flex row justify--start align--center gap--8">
              <icon class="icon">
                <i class="iconoir-3d-select-face"></i>
              </icon>
              <h6 class="m--0">Doxavet</h6>
            </row>

            <row class="flex row justify--start md:justify--center align--center gap--8">
              <small class="text--12"> {new Date().getFullYear()} </small>
              <icon class="icon">
                <i class="iconoir-copyright"></i>
              </icon>
              <small class="text--12"> All rights reserved </small>
            </row>

            <row class="flex row justify--start md:justify--end align--center gap--8">
              <a href={void 0} class="icon">
                <i class="iconoir-facebook-tag"></i>
              </a>
              <a href={void 0} class="icon">
                <i class="iconoir-instagram"></i>
              </a>
              <a href={void 0} class="icon">
                <i class="iconoir-tiktok"></i>
              </a>
            </row>
          </grid>
        </container>
      </footer>
    </>
  );
};

export default Footer;
