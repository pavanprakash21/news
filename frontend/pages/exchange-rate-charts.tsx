import { GetStaticProps } from "next";
import dynamic from 'next/dynamic'

import { generateChartsData } from "../utils";

type finalData = {
  name: string;
  btc: number;
  inr: number;
}

const ExchangeRateChart = dynamic(
  // @ts-ignore
  () => import('../components/ExchangeRateChart').then(comp => comp.ExchangeRateChart),
  { ssr: false }
)

const ExchangeRateCharts = (finalData: finalData[]) => {
  // @ts-ignore
  return <ExchangeRateChart finalData={finalData} style={{marginTop: '2rem'}}/>
};

export const getStaticProps: GetStaticProps = async ({ params }) => {
  const data = await generateChartsData();

  const finalData = data.map((item) => {
    return {
      name: item.date,
      btc: 1 / item.rates.BTC,
      inr: item.rates.INR,
    };
  });

  return {
    props: { finalData },
    revalidate: 1,
  };
};

export default ExchangeRateCharts;
