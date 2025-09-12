import Card from 'react-bootstrap/Card';
import Placeholder from 'react-bootstrap/Placeholder';

import type { WeatherReport } from '../api/WeatherReport';

interface WeatherReportDetails {
  wr: WeatherReport | null;
}

export default function WeatherReportCard({ wr }: WeatherReportDetails) {
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
