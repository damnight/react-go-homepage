import openmeteo_requests

import pandas as pd
import requests_cache
from retry_requests import retry

# Setup the Open-Meteo API client with cache and retry on error
cache_session = requests_cache.CachedSession(".cache", expire_after=3600)
retry_session = retry(cache_session, retries=5, backoff_factor=0.2)
openmeteo = openmeteo_requests.Client(session=retry_session)

# Make sure all required weather variables are listed here
# The order of variables in hourly or daily is important to assign them correctly below
url = "https://api.open-meteo.com/v1/forecast"
params = {
    "latitude": 47.7241,
    "longitude": 9.4698,
    "hourly": [
        "temperature_2m",
        "precipitation_probability",
        "precipitation",
        "cloud_cover",
        "wind_direction_10m",
        "wind_direction_180m",
        "temperature_180m",
        "uv_index",
        "temperature_950hPa",
        "surface_pressure",
    ],
    "current": [
        "temperature_2m",
        "precipitation",
        "surface_pressure",
        "wind_speed_10m",
        "wind_direction_10m",
        "cloud_cover",
    ],
    "forecast_days": 3,
}
responses = openmeteo.weather_api(url, params=params)

# Process first location. Add a for-loop for multiple locations or weather models
response = responses[0]
print(f"Coordinates: {response.Latitude()}°N {response.Longitude()}°E")
print(f"Elevation: {response.Elevation()} m asl")
print(f"Timezone difference to GMT+0: {response.UtcOffsetSeconds()}s")

# Process current data. The order of variables needs to be the same as requested.
current = response.Current()
current_temperature_2m = current.Variables(0).Value()
current_precipitation = current.Variables(1).Value()
current_surface_pressure = current.Variables(2).Value()
current_wind_speed_10m = current.Variables(3).Value()
current_wind_direction_10m = current.Variables(4).Value()
current_cloud_cover = current.Variables(5).Value()

print(f"\nCurrent time: {current.Time()}")
print(f"Current temperature_2m: {current_temperature_2m}")
print(f"Current precipitation: {current_precipitation}")
print(f"Current surface_pressure: {current_surface_pressure}")
print(f"Current wind_speed_10m: {current_wind_speed_10m}")
print(f"Current wind_direction_10m: {current_wind_direction_10m}")
print(f"Current cloud_cover: {current_cloud_cover}")

# Process hourly data. The order of variables needs to be the same as requested.
hourly = response.Hourly()
hourly_temperature_2m = hourly.Variables(0).ValuesAsNumpy()
hourly_precipitation_probability = hourly.Variables(1).ValuesAsNumpy()
hourly_precipitation = hourly.Variables(2).ValuesAsNumpy()
hourly_cloud_cover = hourly.Variables(3).ValuesAsNumpy()
hourly_wind_direction_10m = hourly.Variables(4).ValuesAsNumpy()
hourly_wind_direction_180m = hourly.Variables(5).ValuesAsNumpy()
hourly_temperature_180m = hourly.Variables(6).ValuesAsNumpy()
hourly_uv_index = hourly.Variables(7).ValuesAsNumpy()
hourly_temperature_950hPa = hourly.Variables(8).ValuesAsNumpy()
hourly_surface_pressure = hourly.Variables(9).ValuesAsNumpy()

hourly_data = {
    "date": pd.date_range(
        start=pd.to_datetime(hourly.Time(), unit="s", utc=True),
        end=pd.to_datetime(hourly.TimeEnd(), unit="s", utc=True),
        freq=pd.Timedelta(seconds=hourly.Interval()),
        inclusive="left",
    )
}

hourly_data["temperature_2m"] = hourly_temperature_2m
hourly_data["precipitation_probability"] = hourly_precipitation_probability
hourly_data["precipitation"] = hourly_precipitation
hourly_data["cloud_cover"] = hourly_cloud_cover
hourly_data["wind_direction_10m"] = hourly_wind_direction_10m
hourly_data["wind_direction_180m"] = hourly_wind_direction_180m
hourly_data["temperature_180m"] = hourly_temperature_180m
hourly_data["uv_index"] = hourly_uv_index
hourly_data["temperature_950hPa"] = hourly_temperature_950hPa
hourly_data["surface_pressure"] = hourly_surface_pressure

hourly_dataframe = pd.DataFrame(data=hourly_data)
print("\nHourly data\n", hourly_dataframe)
