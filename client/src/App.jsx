import React from "react";
import { Route, Routes } from "react-router-dom";
import { Home, CreateBlog, Blog } from "./pages";
import { Navbar, Footer } from "./components";

const App = () => {
  return (
    <>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/createBlog" element={<CreateBlog />} />
        <Route path="/blog/:blogId" element={<Blog />} />
      </Routes>
      <Footer />
    </>
  );
};

export default App;
