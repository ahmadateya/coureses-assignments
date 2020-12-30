#!/usr/bin/env python
# coding: utf-8

# # Statistics from Stock Data

# We will start by reading in any of the above CSV files into a DataFrame and see what the data looks like.

# We import pandas into Python
import pandas as pd

# We read in a stock data data file into a data frame and see what it looks like
df = pd.read_csv('./GOOG.csv')

# We display the first 5 rows of the DataFrame
df.head()


# We load the Google stock data into a DataFrame
google_stock = pd.read_csv('./GOOG.csv' , index_col=['Date'], parse_dates=True, usecols=['Date', 'Adj Close'])

# We load the Apple stock data into a DataFrame
apple_stock = pd.read_csv('./AAPL.csv' , index_col=['Date'], parse_dates=True, usecols=['Date', 'Adj Close']) 

# We load the Amazon stock data into a DataFrame
amazon_stock =  pd.read_csv('./AMZN.csv',  index_col=['Date'], parse_dates=True, usecols=['Date', 'Adj Close'])


# We display the google_stock DataFrame
google_stock.head()


# We create calendar dates between '2000-01-01' and  '2016-12-31'
dates = pd.date_range('2000-01-01', '2016-12-31')

# We create and empty DataFrame that uses the above dates as indices
all_stocks = pd.DataFrame(index = dates)


# Change the Adj Close column label to Google
google_stock.rename(columns ={'Adj Close': 'Google'}, inplace=True)

# Change the Adj Close column label to Apple
apple_stock.rename(columns = {'Adj Close': 'Apple'}, inplace=True)

# Change the Adj Close column label to Amazon
amazon_stock.rename(columns = {'Adj Close': 'Amazon'}, inplace=True)


# We display the google_stock DataFrame
google_stock.head()


# We join the Google stock to all_stocks
all_stocks = all_stocks.join(google_stock)

# We join the Apple stock to all_stocks
all_stocks = all_stocks.join(apple_stock)

# We join the Amazon stock to all_stocks
all_stocks = all_stocks.join(amazon_stock)


# We display the google_stock DataFrame
all_stocks.head()


# Check if there are any NaN values in the all_stocks dataframe
all_stocks.isnull()

# Remove any rows that contain NaN values
all_stocks.dropna(axis=0)



# Print the average stock price for each stock
print("Mean:\n" , all_stocks.mean())
# Print the median stock price for each stock
print("Median:\n" ,all_stocks.median())
# Print the standard deviation of the stock price for each stock  
print("Standard Deviation:\n" ,all_stocks.std())
# Print the correlation between stocks
print("Coorelation:\n" ,all_stocks.corr())


# We compute the rolling mean using a 150-Day window for Google stock
rollingMean = all_stocks['Google'].rolling(150).mean()
print(rollingMean)


# this allows plots to be rendered in the notebook
get_ipython().run_line_magic('matplotlib', 'inline')

# We import matplotlib into Python
import matplotlib.pyplot as plt


# We plot the Google stock data
plt.plot(all_stocks['Google'])

# We plot the rolling mean ontop of our Google stock data
plt.plot(rollingMean)
plt.legend(['Google Stock Price', 'Rolling Mean'])
plt.show()

