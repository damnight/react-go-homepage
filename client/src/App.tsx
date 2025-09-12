import Container from 'react-bootstrap/Container'
import WeatherReportCard from './components/WeatherReport.tsx'




function App() {
  const selectedWeatherReport = null

  return (
    <div>
      <h1 className='text-center mt-4'>Weather Dashboard</h1>
      <Container className='mt-4'>
        <WeatherReportCard wr={selectedWeatherReport}/>
      </Container>
      </div>
  );
}

export default App;
