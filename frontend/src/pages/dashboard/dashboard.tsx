import React, { useState } from "react";
import {
  Layout,
  Card,
  Input,
  Button,
  Row,
  Col,
  Typography,
  Switch,
  notification,
} from "antd";
import { Bar } from "@ant-design/plots"; // For chart
import { BellOutlined, UserOutlined, DollarCircleOutlined, CoffeeOutlined } from "@ant-design/icons";
import "./Dashboard.css"; // Importing the CSS file

const { Header, Content } = Layout;
const { Title, Text } = Typography;

const Dashboard = () => {
  const [darkMode, setDarkMode] = useState(false);

  const stats = [
    {
      number: "61",
      label: "AVAILABLE TO SELL",
      icon: <BellOutlined style={{ fontSize: "32px", color: "#1890ff" }} />,
    },
    {
      number: "70K",
      label: "TOTAL OCC",
      icon: <DollarCircleOutlined style={{ fontSize: "32px", color: "#9b4d96" }} />,
    },
    {
      number: "1",
      label: "DIRTY",
      icon: <CoffeeOutlined style={{ fontSize: "32px", color: "#faad14" }} />,
    },
    {
      number: "60",
      label: "VACANT",
      icon: <BellOutlined style={{ fontSize: "32px", color: "#52c41a" }} />,
    },
    {
      number: "0/0",
      label: "EXP ARR",
      icon: <UserOutlined style={{ fontSize: "32px", color: "#1890ff" }} />,
    },
    {
      number: "0/0",
      label: "EXP DEP",
      icon: <UserOutlined style={{ fontSize: "32px", color: "#ff4d4f" }} />,
    },
  ];

  // Dummy bar chart data
  const chartData = [
    { type: "Available", value: 20 },
    { type: "Occupied", value: 20 },
    { type: "Dirty", value: 1 },
    { type: "Vacant", value: 20 },
  ];

  const toggleDarkMode = () => {
    setDarkMode(!darkMode);
  };

  const handleNotification = () => {
    notification.info({
      message: "New Booking",
      description: "You have a new booking request.",
    });
  };

  return (
    <Layout className="dashboard-container">
      {/* Header */}
      <Header className="dashboard-header">
        <div>
          <Title level={3} className="dashboard-title">
            Dashboard
          </Title>
        </div>
        <div className="dashboard-header-right">
          <Text>EGP - EGP</Text>
          <Text>ADMIN</Text>
          <Switch checked={darkMode} onChange={toggleDarkMode} />
        </div>
      </Header>

      {/* Main Content */}
      <Content className="dashboard-stats-grid">
        {/* Stats Grid */}
        <Row gutter={16}>
          {stats.map((stat, index) => (
            <Col xs={24} sm={12} md={8} lg={6} key={index} className="dashboard-col-md-8">
              <Card
                hoverable
                title={stat.number}
                extra={stat.icon}
                bordered={false}
                className="dashboard-stat-card"
              >
                <Text className="dashboard-stat-card-text">{stat.label}</Text>
              </Card>
            </Col>
          ))}
        </Row>

        {/* Interactive Chart */}
        <Row gutter={16} className="dashboard-row-margin">
          <Col xs={24}>
            <Card title="Room Statistics" className="dashboard-chart-card">
              <Bar
                data={chartData}
                xField="type"
                yField="value"
                colorField="type"
                seriesField="type"
                label={{
                  position: "middle",
                  style: { fill: "#fff", fontSize: 16 },
                }}
              />
            </Card>
          </Col>
        </Row>

        {/* Expected Arrival/Departure Section */}
        <Row gutter={16} className="dashboard-row-margin">
          <Col xs={24} sm={12} md={12} lg={8}>
            <Card
              title="Expected Arrival"
              extra={
                <Button type="primary" onClick={handleNotification}>
                  Search
                </Button>
              }
            >
              <Input placeholder="Search..." style={{ marginBottom: "10px" }} />
              <div className="dashboard-empty-state">
                No Expected Arrival To Display
              </div>
            </Card>
          </Col>
          <Col xs={24} sm={12} md={12} lg={8}>
            <Card
              title="Expected Departure"
              extra={
                <Button type="primary" onClick={handleNotification}>
                  Search
                </Button>
              }
            >
              <Input placeholder="Search..." style={{ marginBottom: "10px" }} />
              <div className="dashboard-empty-state">
                No Expected Departure To Display
              </div>
            </Card>
          </Col>
        </Row>
      </Content>
    </Layout>
  );
};

export default Dashboard;
