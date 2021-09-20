import { GetStaticPaths, GetStaticProps } from "next";
import Head from "next/head";

import { Navigation } from "../components/Navigation";

import {
  readJsonFromFile,
  getFilesFromDataDir,
  generateRoutes,
} from "../utils";

import { NewsProps, PathsEntity } from "../types";
import { Article } from "../components/Article/Article";

const News = ({ news, paths }: NewsProps) => {
  return (
    <>
      <Head>
        <title>News</title>
        <meta name="description" content="News and other stuff" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <nav className="bg-gray-200 p-2 mt-0 fixed w-full z-10 top-0">
        <div className="flex flex-wrap items-center content-evenly justify-evenly">
          {news?.data?.news?.map((newsEntity) => {
            return (
              <a
                className="text-gray-500 text-sm sm:text-base line-clamp-3 px-3"
                key={newsEntity.topic}
                href={`#${newsEntity.topic}`}
              >
                {newsEntity.topic}
              </a>
            );
          })}
        </div>
      </nav>

      <main>
        {news &&
          news.data &&
          news.data.news &&
          news.data.news.map((newsEntity) => {
            return (
              <div key={newsEntity.topic}>
                <h3
                  id={newsEntity.topic}
                  className="text-center container mx-44 text-gray-500 text-sm sm:text-base line-clamp-3 px-3 mt-10"
                >
                  {newsEntity.topic}
                </h3>
                {newsEntity.articles &&
                  newsEntity.articles.map((article, index) => {
                    // @ts-ignore
                    return <Article article={article} key={index} />;
                  })}
              </div>
            );
          })}
      </main>
    </>
  );
};

export const getStaticPaths: GetStaticPaths = async () => {
  const files = (await getFilesFromDataDir()) as string[];

  return {
    paths: generateRoutes(files),
    fallback: false,
  };
};

export const getStaticProps: GetStaticProps = async ({ params }) => {
  const news = await readJsonFromFile(`../data/${params?.news}.json`);
  const files = (await getFilesFromDataDir()) as string[];
  const paths = generateRoutes(files);

  return {
    props: { news, paths },
    revalidate: 1,
  };
};

export default News;
