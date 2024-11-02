import React from "react";
import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <div className="flex justify-around p-4">
      <h1 className="text-2xl font-semibold">Blog website</h1>
      <div className="flex gap-4">
        <Link to="/">home</Link>
        <Link to="/createBlog">Create Blog</Link>
        <Link to="/">contact</Link>
      </div>
    </div>
  );
};

export default Navbar;
