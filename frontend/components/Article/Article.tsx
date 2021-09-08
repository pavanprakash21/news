import { ArticlesEntity } from "../../types";

export const Article = (articleObj: ArticlesEntity) => {
  // @ts-ignore
  const article = articleObj.article;
  return (
    <div
      style={{ display: "flex", flexDirection: "column", marginBottom: "1em" }}
    >
      <details>
        <summary>
          <a href={article.url}>
            <h4>{article.title}</h4>
          </a>
        </summary>
        <h5>{article.description}</h5>
        <p>{article.content}</p>
      </details>
      <h6>
        {new Intl.DateTimeFormat("en-GB", {
          dateStyle: "short",
          timeStyle: "short",
        }).format(new Date(article.publishedAt))}
      </h6>
    </div>
  );
};
