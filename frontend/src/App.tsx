import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  ResponsiveContainer,
} from 'recharts'
import Button from 'react-bootstrap/Button'
import ListGroup from 'react-bootstrap/ListGroup'
import Image from 'react-bootstrap/Image'
import Form from 'react-bootstrap/Form'
import Container from 'react-bootstrap/Container'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import Modal from 'react-bootstrap/Modal'
import Spinner from 'react-bootstrap/Spinner'
import { LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { DatePicker } from '@mui/x-date-pickers/DatePicker'
import { useQuery, gql } from '@apollo/client'
import { useState } from 'react'
import dayjs from 'dayjs'

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
      Name
      Location
      Url
      AvatarUrl
    }
  }
`

function Filter({
  defaultFromDate,
  defaultToDate,
  defaultLocation,
  onFilterChange,
  onNotifyClick,
}: any) {
  const [fromDate, setFromDate] = useState(defaultFromDate)
  const [toDate, setToDate] = useState(defaultToDate)
  const [location, setLocation] = useState("")
  return (
    <>
      <div className="shadow rounded p-3" style={{ width: '1000px' }}>
        <Container>
          <Row>
            <Col>
              <Form.Control
                value={location}
                onChange={(e) => setLocation(e.target.value)}
                placeholder={defaultLocation}
              />
            </Col>
            <Col>
              <LocalizationProvider dateAdapter={AdapterDayjs}>
                <Container fluid="sm">
                  <Row>
                    <Col>
                      <DatePicker
                        value={fromDate}
                        onChange={(newValue) =>
                          newValue && setFromDate(newValue)
                        }
                        slotProps={{ textField: { size: 'small' } }}
                        format="DD.MM.YY"
                        maxDate={toDate}
                      />
                    </Col>
                    <Col>
                      <DatePicker
                        value={toDate}
                        onChange={(newValue) => newValue && setToDate(newValue)}
                        slotProps={{ textField: { size: 'small' } }}
                        format="DD.MM.YY"
                        minDate={fromDate}
                      />
                    </Col>
                  </Row>
                </Container>
              </LocalizationProvider>
            </Col>
            <Col md="auto">
              <Button
                variant="primary"
                className="me-3"
                onClick={() =>
                  onFilterChange({
                    fromDate: fromDate,
                    toDate: toDate,
                    location: location,
                  })
                }
              >
                Explore
              </Button>
              <Button variant="outline-primary" onClick={onNotifyClick}>
                Notify
              </Button>
            </Col>
          </Row>
        </Container>
      </div>
    </>
  )
}

function Organization({ githubId, stars, color }: any) {
  const { loading, data } = useQuery(GET_ORGANIZATION, {
    variables: { githubId: githubId },
  })
  return (
    <>
      <ListGroup.Item>
        {loading || (
          <>
            <Container>
              <Row>
                <span
                  style={{
                    background: color,
                    width: '10px',
                    height: '10px',
                    borderRadius: '50%',
                    display: 'inline-block',
                    position: 'absolute',
                    top: '3px',
                    left: '3px',
                    padding: '0',
                  }}
                ></span>
                <Col>
                  <a href={data.organization.Url}>
                    <Image
                      src={data.organization.AvatarUrl}
                      rounded
                      width="50px"
                    />
                  </a>
                </Col>
                <Col>
                  <p>{data.organization.Name}</p>
                </Col>
                <Col>
                  <p>{data.organization.Location}</p>
                </Col>
                <Col>
                  <p>{`${stars} \u2197`}</p>
                </Col>
              </Row>
            </Container>
          </>
        )}
      </ListGroup.Item>
    </>
  )
}

function Chart({ fromDate, toDate, location }: any) {
  const { loading, error, data } = useQuery(GET_TIME_POINTS, {
    variables: {
      fromDate: fromDate.unix(),
      toDate: toDate.unix(),
      location: location,
    },
  })
  let colors: any
  let organizations
  let chartData
  if (error) return <p>Something went wrong..</p>
  if (!loading) {
    colors = [
      'blue',
      'red',
      'green',
      'orange',
      'purple',
      'brown',
      'pink',
      'darkgreen',
      'black',
      'gray',
    ]
    organizations = data.timePoints
      .reduce(
        (acc: any, timePoint: any) =>
          timePoint.Values.reduce((acc: any, value: any) => {
            acc = acc.filter(
              (organization: any) => organization.name != value.Name
            )
            acc.push({ name: value.Name, stars: value.Value })
            return acc
          }, acc),
        []
      )
      .toSorted((a: any, b: any) => b.stars - a.stars)
    organizations.forEach(
      (organization: any, i: any) => (organization['color'] = colors[i])
    )
    chartData = data.timePoints.map((element: any) =>
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
  }
  return (
    <>
      <Container>
        <Row
          className="align-items-center justify-content-center"
          style={{ height: '500px' }}
        >
          {loading ? (
            <span>
              <Spinner
                animation="border"
                role="status"
                variant="primary"
                style={{ width: '50px', height: '50px', marginBottom: '15px' }}
              />
              <p className="text-primary">
                Travelling through the GitHub Universe...
              </p>
            </span>
          ) : (
            <>
              <Col className="h-100">
                <p style={{ color: 'rgb(100,100,100)' }}>
                  Accumulated Received Stars of all Repositories
                </p>
                <ResponsiveContainer width="100%" height="100%">
                  <LineChart data={chartData}>
                    <CartesianGrid strokeDasharray="3 3" />
                    <XAxis
                      dataKey="time"
                      type="number"
                      domain={['dataMin', 'dataMax']}
                      tickFormatter={(unixTime) =>
                        dayjs.unix(unixTime).format('DD.MM.YY')
                      }
                    />
                    <YAxis type="number" />
                    {organizations.map((organization: any) => {
                      return (
                        <Line
                          type="basis"
                          dataKey={organization.name}
                          strokeWidth={2}
                          stroke={organization.color}
                          dot={false}
                          connectNulls
                        />
                      )
                    })}
                  </LineChart>
                </ResponsiveContainer>
              </Col>
              <Col className="h-100">
                <ListGroup className="h-100 overflow-scroll">
                  {organizations.map((organization: any) => {
                    return (
                      <Organization
                        key={organization.name}
                        githubId={parseInt(organization.name)}
                        stars={organization.stars}
                        color={organization.color}
                      />
                    )
                  })}
                </ListGroup>
              </Col>
            </>
          )}
        </Row>
      </Container>
    </>
  )
}

function NotificationModal({ show, onClose }: any) {
  const [notificationType, setNotificationType] = useState('trigger')
  return (
    <>
      <Modal show={show} onHide={onClose}>
        <Modal.Header closeButton>
          <Modal.Title>Notifications</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            <Form.Group className="mb-3" controlId="type">
              <Form.Label>Type</Form.Label>
              <Form.Select
                onChange={(e) => setNotificationType(e.target.value)}
              >
                <option value="trigger">Trigger</option>
                <option value="newsletter">Newsletter</option>
              </Form.Select>
            </Form.Group>
            {notificationType == 'trigger' && (
              <>
                <Form.Group className="mb-3" controlId="organization">
                  <Form.Label>Organization</Form.Label>
                  <Form.Control type="text" placeholder="Google" />
                </Form.Group>
                <Form.Group className="mb-3" controlId="location">
                  <Form.Label>Location</Form.Label>
                  <Form.Control type="text" placeholder="Munich" />
                </Form.Group>
                <Row>
                  <Col>
                    <Form.Group className="mb-3" controlId="timeInterval">
                      <Form.Label>Time Interval Days</Form.Label>
                      <Form.Control type="number" placeholder="1" />
                    </Form.Group>
                  </Col>
                  <Col>
                    <Form.Group className="mb-3" controlId="threshold">
                      <Form.Label>Threshold</Form.Label>
                      <Form.Control type="number" placeholder="50" />
                    </Form.Group>
                  </Col>
                  <Col>
                    <Form.Group>
                      <Form.Label>Metric</Form.Label>
                      <Form.Select>
                        <option value="1">Stars Accumulated</option>
                        <option value="2">Members</option>
                        <option value="3">Followers</option>
                      </Form.Select>
                    </Form.Group>
                  </Col>
                </Row>
              </>
            )}
            {notificationType == 'newsletter' && (
              <>
                <Form.Group className="mb-3" controlId="organization">
                  <Form.Label>Organization</Form.Label>
                  <Form.Control type="text" placeholder="Google" />
                </Form.Group>
                <Form.Group className="mb-3" controlId="location">
                  <Form.Label>Location</Form.Label>
                  <Form.Control type="text" placeholder="Munich" />
                </Form.Group>
                <Form.Group className="mb-3" controlId="timeInterval">
                  <Form.Label>Time Schedule Days</Form.Label>
                  <Form.Control type="number" placeholder="1" />
                </Form.Group>
              </>
            )}
            <Form.Group className="mb-3" controlId="communication">
              <Form.Label>Communication</Form.Label>
              <Form.Select>
                <option value="1">Slack</option>
                <option value="2">E-Mail</option>
              </Form.Select>
            </Form.Group>
            <Form.Group className="mb-3" controlId="email">
              <Form.Label>E-Mail</Form.Label>
              <Form.Control type="email" placeholder="john.doe@tum.de" />
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="primary" onClick={onClose}>
            Save
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  )
}

function Header() {
  return (
    <div
      style={{
        position: 'fixed',
        top: '0',
        left: '0px',
        width: '100%',
        marginTop: '65px',
      }}
    >
      <svg
        height="70px"
        viewBox="0 0 512 512"
        style={{ position: 'relative', top: '-12px' }}
      >
        <g>
          <path
            d="M452.425,202.575l-38.269-23.11c-1.266-10.321-5.924-18.596-13.711-21.947l-86.843-52.444l-0.275,0.598
		c-3.571-7.653-9.014-13.553-16.212-16.668L166.929,10.412l-0.236,0.543v-0.016c-3.453-2.856-7.347-5.239-11.594-7.08
		c-32.315-13.923-74.124,11.013-93.38,55.716c-19.241,44.624-8.7,92.215,23.622,106.154c4.256,1.826,8.669,3.005,13.106,3.556
		l-0.19,0.464l146.548,40.669c7.19,3.107,15.206,3.004,23.229,0.37l-0.236,0.566L365.55,238.5
		c7.819,3.366,17.094,1.125,25.502-5.082l42.957,11.909c7.67,3.312,18.014-3.548,23.104-15.362
		C462.202,218.158,460.11,205.894,452.425,202.575z M154.516,99.56c-11.792,27.374-31.402,43.783-47.19,49.132
		c-6.962,2.281-13.176,2.556-17.605,0.637c-14.536-6.254-25.235-41.856-8.252-81.243c16.976-39.378,50.186-56.055,64.723-49.785
		c4.429,1.904,8.519,6.592,11.626,13.246C164.774,46.699,166.3,72.216,154.516,99.56z"
          />
          <path
            d="M297.068,325.878c-1.959-2.706-2.25-6.269-0.724-9.25c1.518-2.981,4.562-4.846,7.913-4.846h4.468
		c4.909,0,8.889-3.972,8.889-8.897v-7.74c0-4.909-3.98-8.897-8.889-8.897h-85.789c-4.908,0-8.897,3.988-8.897,8.897v7.74
		c0,4.925,3.989,8.897,8.897,8.897h4.492c3.344,0,6.388,1.865,7.914,4.846c1.518,2.981,1.235,6.544-0.732,9.25L128.715,459.116
		c-3.225,4.287-2.352,10.36,1.927,13.569c4.295,3.225,10.368,2.344,13.578-1.943l107.884-122.17l4.036,153.738
		c0,5.333,4.342,9.691,9.691,9.691c5.358,0,9.692-4.358,9.692-9.691l4.043-153.738l107.885,122.17
		c3.209,4.287,9.282,5.168,13.568,1.943c4.288-3.209,5.145-9.282,1.951-13.569L297.068,325.878z"
          />
          <path
            d="M287.227,250.81c0-11.807-9.573-21.388-21.396-21.388c-11.807,0-21.38,9.582-21.38,21.388
		c0,11.831,9.574,21.428,21.38,21.428C277.654,272.238,287.227,262.642,287.227,250.81z"
          />
        </g>
      </svg>
      <p
        className="h1"
        style={{
          display: 'inline-block',
          marginLeft: '10px',
        }}
      >
        CodeLens
      </p>
    </div>
  )
}

function App() {
  const [filter, setFilter] = useState({
    fromDate: dayjs('2023-05-01'),
    toDate: dayjs('2023-05-06'),
    location: 'Germany',
  })
  const [showNotificationModal, setShowNotificationModal] = useState(false)
  return (
    <>
      <Header />
      <Container>
        <Row
          className="justify-content-md-center"
          style={{ marginBottom: '80px' }}
        >
          <Filter
            defaultFromDate={filter.fromDate}
            defaultToDate={filter.toDate}
            defaultLocation={filter.location}
            onFilterChange={setFilter}
            onNotifyClick={() => setShowNotificationModal(true)}
          />
        </Row>
        <Row>
          <Chart
            fromDate={filter.fromDate}
            toDate={filter.toDate}
            location={filter.location}
          />
        </Row>
      </Container>
      <NotificationModal
        show={showNotificationModal}
        onClose={() => setShowNotificationModal(false)}
      />
    </>
  )
}

export default App
