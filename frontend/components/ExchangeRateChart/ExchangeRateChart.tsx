import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from "recharts";

type finalData = {
  name: string;
  btc: number;
  inr: number;
};

export const ExchangeRateChart = (finalData: finalData[]) => {
  // @ts-ignore
  const data = finalData.finalData.finalData
  const inrArr = data.map((rate: finalData) => rate.inr)
  const btcArr = data.map((rate: finalData) => rate.btc)
  const inrDomain = [Math.round(Math.min(...inrArr) - 1), Math.round(Math.max(...inrArr) + 1)]
  const btcDomain = [Math.round(Math.min(...btcArr) - 500), Math.round(Math.max(...btcArr) + 500)]
  return (
    <>
      {/* @ts-ignore */}
      <LineChart width={900} height={500} data={data} style={{float: 'left', width: '50%', marginTop: '25vh'}}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="name" padding={{ left: 30, right: 30 }} />
        <YAxis domain={btcDomain} />
        <Tooltip />
        <Legend />
        <Line type="monotone" dataKey="btc" stroke="#8884d8" />
      </LineChart>
      {/* @ts-ignore */}
      <LineChart width={900} height={500} data={data} style={{float: 'right', width: '50%', marginTop: '25vh'}}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="name" padding={{ left: 30, right: 30 }} />
        <YAxis domain={inrDomain} />
        <Tooltip />
        <Legend />
        <Line type="monotone" dataKey="inr" stroke="#82ca9d" />
      </LineChart>
    </>
  );
};
