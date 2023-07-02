import { LineChart, Line, XAxis, YAxis, CartesianGrid } from 'recharts'
import { LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { DatePicker } from '@mui/x-date-pickers/DatePicker'
import { useQuery, gql } from '@apollo/client'
import { useState } from 'react'
import Button from 'react-bootstrap/Button'
import ListGroup from 'react-bootstrap/ListGroup'
import Image from 'react-bootstrap/Image'
import Form from 'react-bootstrap/Form';
import * as dayjs from 'dayjs'

import './App.css'

const GET_TIME_POINTS = gql`
  query timePoints($fromDate: Int!, $toDate: Int!, $location: String!) {
    timePoints(fromDate: $fromDate, toDate: $toDate, location: $location) {
      Time
      Values {
        Name
        Value
      }
    }
  }
`

const GET_ORGANIZATION = gql`
  query organization($githubId: Int!) {
    organization(githubId: $githubId) {
      Login
      AvatarUrl
    }
  }
`

function Filter({ defaultFromDate, defaultToDate, defaultLocation, onFilterChange }: any) {
  const [fromDate, setFromDate] = useState(defaultFromDate)
  const [toDate, setToDate] = useState(defaultToDate)
  const [location, setLocation] = useState(defaultLocation)
  return (
    <>
      <Form.Control value={location} onChange={setLocation}/>
      <LocalizationProvider dateAdapter={AdapterDayjs}>
        <DatePicker
          value={fromDate}
          onChange={(newValue) => newValue && setFromDate(newValue)}
        />
        <DatePicker
          value={toDate}
          onChange={(newValue) => newValue && setToDate(newValue)}
        />
      </LocalizationProvider>
      <Button
        variant="primary"
        onClick={() => onFilterChange({ fromDate: fromDate, toDate: toDate, location: location })}
      >
        Start
      </Button>
    </>
  )
}

function Organization({ githubId }: any) {
  const { loading, error, data } = useQuery(GET_ORGANIZATION, {
    variables: { githubId: githubId },
  })
  return (
    <>
      <ListGroup.Item>
        {loading || (
          <>
            <Image src={data.organization.AvatarUrl} rounded width="50px" />
            <p>{data.organization.Login}</p>
          </>
        )}
      </ListGroup.Item>
    </>
  )
}

function Chart({ fromDate, toDate, location }: any) {
  const { loading, error, data } = useQuery(GET_TIME_POINTS, {
    variables: { fromDate: fromDate.unix(), toDate: toDate.unix(), location: location },
  })
  if (loading) return <p>Loading...</p>
  if (error) return <p>Error : {error.message}</p>
  const chartData = data.timePoints.map((element: any) =>
    Object.assign(
      {
        time: element.Time,
      },
      ...element.Values.map((value: any) => ({ [value.Name]: value.Value }))
    )
  )
  const lastTimePoint = chartData.reduce(
    (lastTimePoint: any, timePoint: any) => ({
      ...lastTimePoint,
      ...timePoint,
    }),
    {}
  )
  const firstTimePoint = { ...lastTimePoint }
  Object.keys(firstTimePoint).forEach((key) => {
    firstTimePoint[key] = 0
  })
  firstTimePoint['time'] = fromDate.unix()
  chartData.unshift(firstTimePoint)
  chartData.push(lastTimePoint)
  const organizations = Object.keys(firstTimePoint).filter(
    (key) => key != 'time'
  )
  const colors = [
    'red',
    'darkblue',
    'blue',
    'green',
    'purple',
    'brown',
    'orange',
    'pink',
    'black',
    'gray',
  ]
  return (
    <>
      <LineChart
        width={500}
        height={300}
        data={chartData}
        margin={{
          top: 5,
          right: 30,
          left: 20,
          bottom: 5,
        }}
      >
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis
          dataKey="time"
          type="number"
          domain={['dataMin', 'dataMax']}
          tickFormatter={(unixTime) => dayjs.unix(unixTime).format('DD/MM/YY')}
        />
        <YAxis />
        {organizations.map((key: any, index: any) => {
          return (
            <Line
              type="basis"
              dataKey={key}
              stroke={colors[index]}
              connectNulls
            />
          )
        })}
      </LineChart>
      <ListGroup>
        {organizations.map((key: any, index: any) => {
          return <Organization key={key} githubId={parseInt(key)} />
        })}
      </ListGroup>
    </>
  )
}

function App() {
  const [filter, setFilter] = useState({
    fromDate: dayjs().subtract(7, 'day'),
    toDate: dayjs(),
    location: "",
  })
  return (
    <>
      <Filter
        defaultFromDate={filter.fromDate}
        defaultToDate={filter.toDate}
        defaultLocation={filter.location}
        onFilterChange={setFilter}
      />
      <Chart fromDate={filter.fromDate} toDate={filter.toDate} location={location} />
    </>
  )
}

export default App
