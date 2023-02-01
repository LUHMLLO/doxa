import {} from "solid-js";

const Header = (props) => {
  return (
    <>
      <header class="bg-complementary py--92 zindex--10">
        <container class="container block">
          <grid className="grid grid-cols--1 lg:grid-cols--2">
            <column className="flex column">
              <h2 class="clr-secondary text--capitalize m--0">
                {props.title || "title goes here"}
              </h2>
              <p class="clr-secondary m--0">
                {props.content || "description goes here"}
              </p>
            </column>
          </grid>
        </container>
      </header>
    </>
  );
};

export default Header;
