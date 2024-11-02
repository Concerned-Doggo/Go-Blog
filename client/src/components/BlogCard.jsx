import { Link } from "react-router-dom";

const BlogCard = ({ blog }) => {
  blog.body = blog.body.slice(0, 100);
  blog.body += "...";
  return (
    <div className="flex flex-col gap-4 m-4 bg-slate-300 p-4 min-w-[150px] max-w-[300px] rounded-xl">
      <Link
        to={`/blog/${blog._id}`}
        className="text-2xl font-semibold text-gray-700"
      >
        {blog.title}
      </Link>
      <div
        className="text-lg text-gray-400"
        dangerouslySetInnerHTML={{ __html: blog.body }}
      ></div>
    </div>
  );
};

export default BlogCard;
