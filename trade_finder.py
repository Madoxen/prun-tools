#!/bin/usr/python3
# Trade finder will try to find best routes based on price, volume 

import requests
import pandas as pd
# data sample of /exchange/all endpoint
# [
# {
#    "MaterialTicker": "WOR",
#    "ExchangeCode": "CI2",
#    "MMBuy": null,
#    "MMSell": null,
#    "PriceAverage": 85078.04,
#    "AskCount": null,
#    "Ask": null,
#    "Supply": 0,
#    "BidCount": null,
#    "Bid": null,
#    "Demand": 0
#  }
#  ... ]

# We will to group by material ticker
# From that we will have information on each good, BID/ASK in each station
# And from that we can calculate which routes are the best

exchange_data = requests.get("https://rest.fnar.net/exchange/all").json()

ex_df = pd.DataFrame(exchange_data)
#remove unecessary clutter
ex_df.drop(["MMBuy", "MMSell"], axis=1, inplace=True)
#Find out where we can buy goods the cheapest
min_asks_idx = ex_df.groupby("MaterialTicker")["Ask"].idxmin().dropna()
#Find out where we can sell goods with the largest price
max_bids_idx = ex_df.groupby("MaterialTicker")["Bid"].idxmax().dropna()

min_asks_df = ex_df.loc[min_asks_idx]
max_bids_df = ex_df.loc[max_bids_idx]

profits_df = pd.merge(left=min_asks_df,
                      right=max_bids_df,
                      on="MaterialTicker",
                      suffixes=("_ASK", "_BID"))

profits_df['profit'] = profits_df["Bid_BID"] -  profits_df["Ask_ASK"]
profits_df.sort_values(by='profit', inplace=True, ascending=False)
profits_df.to_csv("trade_finder.csv")
