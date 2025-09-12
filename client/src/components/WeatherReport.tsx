import Card from 'react-bootstrap/Card';
import Placeholder from 'react-bootstrap/Placeholder';
import type { WeatherReport } from '../api/WeatherReport';
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';

interface WeatherReportDetails {
  wr: WeatherReport | null;
}

async function fetchWeatherReports(): Promise<WeatherReport[]> {
  try {
  const { data } = await axios.get<WeatherReport[]>('http://localhost:5000/weatherreports')
  return data
  } catch (error) {
    console.log(error.message)
    return null
  }
}


export default function WeatherReportCard() {
  const data = fetchWeatherReports()
  wr = data[0];
  return (
    <Card>
      <Card.Body>
        <Card.Title>
          { wr ? (`${wr.city.name}`) : ('Test Title City') }
        </Card.Title>
              <Card.Text>{wr ? (`Temperature: ${wr.temperature}`) : `Temperature: 0`}</Card.Text>
      </Card.Body>
    </Card>
  );
}
