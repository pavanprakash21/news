import fs from "fs";
import util from 'util'

const DATA_FOLDER = '../data'

export const readJsonFromFile = async (filePath: string) => {
  const readFile = util.promisify(fs.readFile)
  const content = await readFile(filePath)
  return JSON.parse(content.toString())
}

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
