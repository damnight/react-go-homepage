import type { City } from '../api/City'

export interface WeatherReport {
	time: string;
	city: City;
	forecast_days: number;
	temperature: string;
	precipitation_probability: number;
	precipitation: number;
	cloud_cover: number;
	wind_direction: number;
	uv_index: number;
	surface_pressure: number;
}
