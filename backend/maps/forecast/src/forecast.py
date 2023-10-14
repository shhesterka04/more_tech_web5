from prophet import Prophet
import matplotlib.pyplot as plt
import pandas as pd
import json
from datetime import datetime, timedelta


with open ('loadperweek.txt', 'r') as file:
  json_data = json.load(file)

# Преобразование в DataFrame
df = pd.DataFrame(json_data)

start_date = pd.to_datetime('2023-10-14')

def convert_to_datetime(df):
    day_of_week = df['day']
    hour = df['hour']
    day_mapping = {
        'Monday': 0,
        'Tuesday': 1,
        'Wednesday': 2,
        'Thursday': 3,
        'Friday': 4
    }
    return start_date + pd.DateOffset(days=day_mapping[day_of_week], hours=hour)


df['datetime'] = df.apply(convert_to_datetime, axis=1)
df = df.drop(['hour', 'day'], axis=1)

df1 = df[['datetime','numIndividuals']]

df2 = df[['datetime','numLegalEntities']]


df1.reset_index()
df2.reset_index()

df1.columns = ['ds', 'y']
df2.columns = ['ds', 'y']


m1 = Prophet()
m2 = Prophet()
m1.fit(df1)
m2.fit(df2)


future1 = m1.make_future_dataframe(periods=7)
forecast1 = m1.predict(future1)

future2 = m2.make_future_dataframe(periods=7)
forecast2 = m2.predict(future1)



forecast = pd.concat([forecast1[['ds', 'yhat']], forecast2[['yhat']]], axis = 1)
forecast.columns = ['ds', 'numIndividualClients', 'numLegalEntities']
forecast = forecast[['numIndividualClients','numLegalEntities']].astype(int)

with open('forecast_data.txt', 'w') as outfile:
    json.dump(json_data, outfile)

json_data = forecast.to_json(orient='records')