import DOMPurify from "dompurify";
import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

const Blog = () => {
  const params = useParams();
  const id = params.blogId;
  const [blog, setBlog] = useState(false);
  useEffect(() => {
    try {
      fetch(`/api/blog/${id}`)
        .then((res) => res.json())
        .then((data) => {
          // important otherwise XSS attack can be done
          const sanitizedBody = DOMPurify.sanitize(data.body);
          setBlog({ ...data, body: sanitizedBody });
        });
    } catch (error) {
      console.log(error);
    }
  }, []);
  return (
    <>
      <div className="mx-auto bg-gray-200 w-3/4 p-4 rounded-xl flex flex-col gap-4 my-10">
        <h1 className="text-center text-3xl font-bold text-gray-700 ">
          {blog && blog.title}
        </h1>
        <h1
          className="text-gray-700 bg-gray-100 p-4 rounded-xl text-lg"
          dangerouslySetInnerHTML={{ __html: blog.body }}
        ></h1>
      </div>
    </>
  );
};

export default Blog;
