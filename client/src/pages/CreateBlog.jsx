import { useState } from "react";
import ReactQuill from "react-quill";
import "react-quill/dist/quill.snow.css";

const CreateBlog = () => {
  const [body, setBody] = useState("");
  const [title, setTitle] = useState("");
  console.log(title);
  console.log(body);

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      await fetch("/api/createBlog", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ title, body }),
      })
        .then((res) => res.json())
        .then((data) => console.log(data));
    } catch (error) {
      console.log(error);
    }
  };
  return (
    <div className="w-3/4 mx-auto flex flex-col gap-4 justify-center">
      <form onSubmit={handleSubmit}>
        <p>Title: </p>
        <input
          className="border-2 border-gray-400  p-2"
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <p>Content:</p>
        <div className="h-[200px]">
          <ReactQuill
            theme="snow"
            value={body}
            onChange={setBody}
            className="h-36 "
          />
        </div>
        <button
          type="submit"
          className="bg-sky-300 px-2 py-1 rounded-md border-2 border-gray-500 self-center"
        >
          submit
        </button>
      </form>
    </div>
  );
};

export default CreateBlog;
