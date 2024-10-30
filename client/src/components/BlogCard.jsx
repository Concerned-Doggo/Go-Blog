const BlogCard = ({ blog }) => {
  return (
    <div className="flex flex-col gap-4 m-4 bg-slate-300 p-4 min-w-[150px] max-w-[300px] rounded-xl">
      <div className="text-2xl font-semibold text-gray-700">{blog.title}</div>
      <div className="text-lg text-gray-400">
        {blog.body.slice(0, 50) + "..."}
      </div>
    </div>
  );
};

export default BlogCard;
