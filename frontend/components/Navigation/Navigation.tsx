import { PathsEntity } from "../../types";

export const Navigation = (paramsObj: PathsEntity) => {
  // @ts-ignore
  const params = paramsObj.params.params;

  return (
    <a href={`news/${params.news}`} key={params.news}>
      {params.news}
    </a>
  );
};
