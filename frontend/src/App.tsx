import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
} from 'recharts'
import { LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { DatePicker } from '@mui/x-date-pickers/DatePicker'
import { useQuery, gql } from '@apollo/client'
import { useState } from 'react'
import Button from 'react-bootstrap/Button'

import './App.css'
import * as dayjs from 'dayjs'

const GET_TIME_POINTS = gql`
  query timePoints($fromDate: Int!, $toDate: Int!) {
    timePoints(fromDate: $fromDate, toDate: $toDate) {
      Time
      Values {
        Name
        Value
      }
    }
  }
`

interface FilterProps {
  defaultFromDate: dayjs.Dayjs
  defaultToDate: dayjs.Dayjs
  onFilterChange: (filter: {
    fromDate: dayjs.Dayjs
    toDate: dayjs.Dayjs
  }) => void
}

interface ChartProps {
  fromDate: dayjs.Dayjs
  toDate: dayjs.Dayjs
}

function Filter({
  defaultFromDate,
  defaultToDate,
  onFilterChange,
}: FilterProps) {
  const [fromDate, setFromDate] = useState(defaultFromDate)
  const [toDate, setToDate] = useState(defaultToDate)
  return (
    <>
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
        onClick={() => onFilterChange({ fromDate: fromDate, toDate: toDate })}
      >
        Start
      </Button>
    </>
  )
}

function Chart({ fromDate, toDate }: ChartProps) {
  const { loading, error, data } = useQuery(GET_TIME_POINTS, {
    variables: { fromDate: fromDate.unix(), toDate: toDate.unix() },
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
  const lastTimePoint = chartData.reduce((lastTimePoint: any, timePoint: any) => ({...lastTimePoint, ...timePoint}), {})
  const firstTimePoint = {...lastTimePoint}
  Object.keys(firstTimePoint).forEach(key => {
    firstTimePoint[key] = 0
  });
  firstTimePoint["time"] = fromDate.unix()
  chartData.unshift(firstTimePoint)
  chartData.push(lastTimePoint)
  const colors = ["red", "darkblue", "blue", "green", "purple", "brown", "orange", "pink", "black", "gray"]
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
        <XAxis dataKey="time" type="number" domain={['dataMin', 'dataMax']} tickFormatter={unixTime => dayjs.unix(unixTime).format('DD/MM/YY')}/>
        <YAxis />
        <Legend />
        {Object.keys(firstTimePoint).filter(key => key != "time").map((key: any, index: any) => {
          return <Line type="basis" dataKey={key} stroke={colors[index]} connectNulls />
        })}
      </LineChart>
    </>
  )
}

function App() {
  const [filter, setFilter] = useState({
    fromDate: dayjs().subtract(7, 'day'),
    toDate: dayjs(),
  })
  return (
    <>
      <Filter
        defaultFromDate={filter.fromDate}
        defaultToDate={filter.toDate}
        onFilterChange={setFilter}
      />
      <Chart fromDate={filter.fromDate} toDate={filter.toDate} />
    </>
  )
}

export default App
