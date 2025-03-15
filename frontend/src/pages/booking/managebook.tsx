import React, { useState } from 'react';
import { 
  Table, 
  Card, 
  Button, 
  Input, 
  Select, 
  DatePicker, 
  Tag, 
  Space, 
  Pagination,
  Dropdown,
  Menu
} from 'antd';
import { 
  PlusOutlined, 
  SearchOutlined, 
  FilterOutlined, 
  EditOutlined, 
  DeleteOutlined 
} from '@ant-design/icons';
import { 
  Calendar,
  User,
  Filter as FilterIcon,
  ChevronDown,
  ChevronUp
} from 'lucide-react';
import './BookingManagement.css';

const { RangePicker } = DatePicker;
const { Option } = Select;

const BookingManagement: React.FC = () => {
  const [activeTab, setActiveTab] = useState('upcoming');
  const [filterOpen, setFilterOpen] = useState(false);

  // Sample booking data
  const bookings = [
    {
      id: 'BK-7842',
      guestName: 'Ahmed Mohamed',
      roomType: 'Deluxe Suite',
      checkIn: '2025-03-05',
      checkOut: '2025-03-08',
      status: 'confirmed',
      payment: 'paid',
      amount: 'EGP 3,600',
      email: 'ahmed.m@example.com'
    },
    // ... (previous booking data remains the same)
  ];

  // Status color mapping
  const statusColorMap = {
    confirmed: 'green',
    pending: 'orange',
    cancelled: 'red'
  };

  // Payment color mapping
  const paymentColorMap = {
    paid: 'blue',
    unpaid: 'red',
    'partially paid': 'purple',
    refunded: 'gray'
  };

  // Table columns configuration
  const columns = [
    {
      title: 'Booking ID',
      dataIndex: 'id',
      key: 'id',
      sorter: (a, b) => a.id.localeCompare(b.id)
    },
    {
      title: 'Guest',
      dataIndex: 'guestName',
      key: 'guestName',
      render: (text, record) => (
        <Space>
          <div className="guest-avatar">
            <User size={20} />
          </div>
          <div>
            <div>{text}</div>
            <div className="text-muted">{record.email}</div>
          </div>
        </Space>
      ),
      sorter: (a, b) => a.guestName.localeCompare(b.guestName)
    },
    {
      title: 'Room Type',
      dataIndex: 'roomType',
      key: 'roomType',
      sorter: (a, b) => a.roomType.localeCompare(b.roomType)
    },
    {
      title: 'Check-in',
      dataIndex: 'checkIn',
      key: 'checkIn',
      render: (text) => new Date(text).toLocaleDateString(),
      sorter: (a, b) => new Date(a.checkIn).getTime() - new Date(b.checkIn).getTime()
    },
    {
      title: 'Check-out',
      dataIndex: 'checkOut',
      key: 'checkOut',
      render: (text) => new Date(text).toLocaleDateString(),
      sorter: (a, b) => new Date(a.checkOut).getTime() - new Date(b.checkOut).getTime()
    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status',
      render: (status) => (
        <Tag color={statusColorMap[status]}>
          {status.charAt(0).toUpperCase() + status.slice(1)}
        </Tag>
      ),
      sorter: (a, b) => a.status.localeCompare(b.status)
    },
    {
      title: 'Payment',
      dataIndex: 'payment',
      key: 'payment',
      render: (payment, record) => (
        <div>
          <Tag color={paymentColorMap[payment]}>
            {payment.charAt(0).toUpperCase() + payment.slice(1)}
          </Tag>
          <div className="text-muted">{record.amount}</div>
        </div>
      ),
      sorter: (a, b) => a.payment.localeCompare(b.payment)
    },
    {
      title: 'Actions',
      key: 'actions',
      render: () => (
        <Space>
          <Button type="text" icon={<EditOutlined />} />
          <Button type="text" danger icon={<DeleteOutlined />} />
        </Space>
      )
    }
  ];

  // Tab menu items
  const tabItems = [
    { key: 'upcoming', label: 'Upcoming' },
    { key: 'past', label: 'Past' },
    { key: 'cancelled', label: 'Cancelled' },
    { key: 'all', label: 'All Bookings' }
  ];

  // Filter dropdown menu
  const filterMenu = (
    <Menu>
      <Menu.Item key="date">
        <RangePicker />
      </Menu.Item>
      <Menu.Item key="status">
        <Select 
          placeholder="Select Status" 
          style={{ width: '100%' }}
        >
          <Option value="confirmed">Confirmed</Option>
          <Option value="pending">Pending</Option>
          <Option value="cancelled">Cancelled</Option>
        </Select>
      </Menu.Item>
      <Menu.Item key="roomType">
        <Select 
          placeholder="Select Room Type" 
          style={{ width: '100%' }}
        >
          <Option value="standard">Standard</Option>
          <Option value="deluxe">Deluxe</Option>
          <Option value="suite">Suite</Option>
        </Select>
      </Menu.Item>
    </Menu>
  );

  return (
    <div className="booking-management">
      <Card 
        title="Booking Management" 
        extra={
          <Button type="primary" icon={<PlusOutlined />}>
            New Booking
          </Button>
        }
      >
        <div className="booking-filters">
          <Menu 
            mode="horizontal" 
            selectedKeys={[activeTab]} 
            onSelect={({ key }) => setActiveTab(key as string)}
          >
            {tabItems.map(tab => (
              <Menu.Item key={tab.key}>{tab.label}</Menu.Item>
            ))}
          </Menu>

          <Space className="search-filter-area">
            <Input 
              prefix={<SearchOutlined />} 
              placeholder="Search bookings" 
            />
            <Dropdown 
              overlay={filterMenu} 
              trigger={['click']}
            >
              <Button>
                <FilterIcon /> Filter 
                {filterOpen ? <ChevronUp /> : <ChevronDown />}
              </Button>
            </Dropdown>
          </Space>
        </div>

        <Table 
          columns={columns} 
          dataSource={bookings} 
          pagination={{
            total: bookings.length,
            showSizeChanger: true,
            showQuickJumper: true
          }}
        />
      </Card>
    </div>
  );
};

export default BookingManagement;