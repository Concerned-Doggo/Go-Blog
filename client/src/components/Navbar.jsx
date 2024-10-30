import React from "react";

const Navbar = () => {
  return (
    <div className="flex justify-around p-4">
      <h1 className="text-2xl font-semibold">Blog website</h1>
      <div className="flex gap-4">
        <div>home</div>
        <div>about</div>
        <div>contact</div>
      </div>
    </div>
  );
};

export default Navbar;
