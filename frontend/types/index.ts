export type NewsProps = {
  news: News;
  paths: Paths;
};

export type PathsEntity = {
  params: Params;
};

type News = {
  data: Data;
};

type Data = {
  news?: NewsEntity[] | null;
};

type NewsEntity = {
  status: string;
  topic: string;
  articles?: ArticlesEntity[] | null;
};

export type ArticlesEntity = {
  source: Source;
  author: string;
  title: string;
  description: string;
  url: string;
  urlToImage: string;
  publishedAt: string;
  content: string;
};

type Source = {
  id: string;
  name: string;
};

type Paths = {
  paths: Array<PathsEntity>;
};


type Params = {
  news: string;
};
