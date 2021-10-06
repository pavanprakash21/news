import fs from "fs";
import util from "util";

import { ExchangeData } from "../types";

const DATA_FOLDER = "../data";

export const readJsonFromFile = async (filePath: string) => {
  const readFile = util.promisify(fs.readFile);
  const content = await readFile(filePath);
  return JSON.parse(content.toString());
};

export const getFilesFromDataDir = (): Promise<string[]> => {
  return new Promise((resolve, reject) => {
    fs.readdir(DATA_FOLDER, (err, files) => {
      if (err) {
        return reject(err);
      }

      resolve(files.map((file) => file.slice(0, -5)));
    });
  });
};

export const generateRoutes = (arr: string[]) => {
  return arr.map((a: string) => {
    return {
      params: { news: a },
    };
  });
};

export const generateChartsData = async () => {
  const files = await getFilesFromDataDir();

  const exchangeDataArr = await Promise.all(
    files.slice(-10).map(exchangeResultData)
  );

  return exchangeDataArr;
};

const exchangeResultData = async (file: string) => {
  const fileContent = await readJsonFromFile(`../data/${file}.json`);
  const exchange_result = fileContent["exchange_result"] || {};
  const exchangeData: ExchangeData = {
    rates: exchange_result["rates"],
    date: smallDate(exchange_result["date"]),
  };
  return exchangeData;
};

const smallDate = (date: string) => {
  const dateParts = date.split("-");
  return `${dateParts[2]}/${dateParts[1]}`;
};
