import { ArticlesEntity } from "../../types";

export const Article = (articleObj: ArticlesEntity) => {
  // @ts-ignore
  const article = articleObj.article;

  return (
    <div className="bg-white p-2 my-2 mx-auto w-90 max-w-7xl sm:w-full sm:p-4 h-auto sm:h-64 rounded-2xl shadow-lg flex flex-col sm:flex-row gap-5 select-none">
      <div className="flex sm:flex-1 flex-col gap-2 p-1">
        <h1 className="text-lg sm:text-xl font-semibold  text-gray-600">
          {article.title}
        </h1>
        <p className="text-gray-500 text-sm sm:text-base line-clamp-3">
          {article.description}
        </p>
        <p className="text-gray-500 text-sm sm:text-base line-clamp-3">
          {article.content}
        </p>
        <div className="flex gap-4 mt-auto">
          <p className="flex items-center gap-1 sm:text-sm px-3 py-1 rounded-full hover:bg-gray-50 transition-colors focus:bg-gray-100 focus:outline-none focus-visible:border-gray-500">
            {new Intl.DateTimeFormat("en-GB", {
              dateStyle: "short",
              timeStyle: "short",
            }).format(new Date(article.publishedAt))}
          </p>
          <a
            href={article.url}
            className="ml-auto flex items-center gap-1 sm:text-lg border border-gray-300 px-3 py-1 rounded-full hover:bg-gray-50 transition-colors focus:bg-gray-100 focus:outline-none focus-visible:border-gray-500"
          >
            <span>Read more</span>
          </a>
        </div>
      </div>
    </div>
  );
};
