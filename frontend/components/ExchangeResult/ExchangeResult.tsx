import {ExchangeResult as ExchangeResultProp, Rates} from '../../types'

export const ExchangeResult = (result: ExchangeResultProp | undefined) => {
  // @ts-ignore
  result = result.result;
  if (result && !result.success) {
    return(
      <h3>Exchange Rates couldn&apos;t be determined for today</h3>
    )
  }

  return(
    <section className="container mt-12 ml-24">
      <h3>1 BTC = {1 / (result?.rates['BTC'] || 1)} Euros</h3>
      <h3>1 EUR = {result && result.rates['INR']} Rupees</h3>
    </section>
  )
}
