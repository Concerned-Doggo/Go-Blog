import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import BlogCard from "../components/BlogCard";
import DOMPurify from "dompurify";

function Home() {
  const [apiData, setApiData] = useState(false);
  useEffect(() => {
    const fetchData = async () => {
      try {
        console.log("calling api");
        await fetch("/api/blogs")
          .then((res) => res.json())
          .then((data) => {
            const sanitizedBody = DOMPurify.sanitize(data.body);
            data.body = sanitizedBody;
            setApiData(data);
          });
      } catch (error) {
        console.log(error);
      }
    };

    fetchData();
  }, []);

  return (
    <>
      <h1 className="text-3xl text-center font-bold text-gray-700">
        Recent Blogs!
      </h1>
      <div className="flex gap-4 justify-center flex-wrap">
        {apiData &&
          apiData.map((blog) => <BlogCard blog={blog} key={blog._id} />)}
      </div>
    </>
  );
}

export default Home;
