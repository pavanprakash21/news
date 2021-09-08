import { PathsEntity } from "../../types";

type params = {
  news: string
}

export const Navigation = (paramsObj: PathsEntity) => {
  // @ts-ignore
  const params: params = paramsObj.params.params;

  return (
    // @ts-ignore https://github.com/microsoft/TypeScript/issues/44418#issuecomment-910551704
    <a href={`news/${params.news}`} key={params.news}>
      {params.news}
      {/* @ts-ignore */}
    </a>
  );
};
