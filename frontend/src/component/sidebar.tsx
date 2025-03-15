import React, { useState } from "react";
import { Layout, Menu } from "antd";
import {
  HomeOutlined,
  CheckCircleOutlined,
  LogoutOutlined,
  SettingOutlined,
  AppstoreOutlined,
  ProfileOutlined,
  ShopOutlined,
  QuestionCircleOutlined,
} from "@ant-design/icons";
import { Link } from "react-router-dom";
import logo1 from "../assets/logo.png"; // ตรวจสอบ path ให้ถูกต้อง
import "./sidebar.css";

const { Sider } = Layout;

const Sidebar = () => {
  const [collapsed, setCollapsed] = useState(false);

  return (
    <Sider collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)} className="sidebar">
      <div className="sidebar-header">
        <img
          src={logo1}
          alt="Logo"
          className="sidebar-logo"
          style={{ width: collapsed ? "40px" : "120px", transition: "width 0.3s" }} // ขนาดโลโก้เปลี่ยนตาม sidebar
        />
      </div>
      <Menu theme="dark" mode="vertical" className="sidebar-menu">
        <Menu.Item key="dashboard" icon={<HomeOutlined />}>
          <Link to="/">Dashboard</Link>
        </Menu.Item>
        <Menu.Item key="room-management" icon={<AppstoreOutlined />}>
          <Link to="/roommanage">จัดการห้องพัก</Link>
        </Menu.Item>
        <Menu.Item key="booking" icon={<CheckCircleOutlined />}>
          <Link to="/bookmange">การจองห้องพัก</Link>
        </Menu.Item>
        <Menu.Item key="checkout" icon={<LogoutOutlined />}>
          <Link to="/checkout">Check-in / Check-out</Link>
        </Menu.Item>
        <Menu.Item key="reservation" icon={<ProfileOutlined />}>
          <Link to="/reservation">การจองห้องพัก</Link>
        </Menu.Item>
        <Menu.Item key="pos" icon={<ShopOutlined />}>
          <Link to="/pos">ระบบขายสินค้า</Link>
        </Menu.Item>
        <Menu.Item key="settings" icon={<SettingOutlined />}>
          <Link to="/settings">การทำความสะอาด</Link>
        </Menu.Item>
        <Menu.Item key="help" icon={<QuestionCircleOutlined />}>
          <Link to="/help">การจัดการพนักงาน</Link>
        </Menu.Item>
        <Menu.Item key="logout" icon={<LogoutOutlined />}>
          <Link to="/logout">ออกจากระบบ</Link>
        </Menu.Item>
      </Menu>
    </Sider>
  );
};

export default Sidebar;
