import React, { useState } from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { Layout, Menu, theme } from 'antd';
import {
  CloudServerOutlined,
  GlobalOutlined,
  ApiOutlined,
  DatabaseOutlined,
  HddOutlined,
  FileTextOutlined,
  SettingOutlined,
  UserOutlined
} from '@ant-design/icons';
import './App.css';

// Pages
import Dashboard from './pages/Dashboard';
import Accounts from './pages/Accounts';
import Zones from './pages/Zones';
import DNS from './pages/DNS';
import Workers from './pages/Workers';
import Pages from './pages/Pages';
import KV from './pages/KV';
import R2 from './pages/R2';

const { Header, Content, Footer, Sider } = Layout;

function App() {
  const [collapsed, setCollapsed] = useState(false);
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  const menuItems = [
    {
      key: '/',
      icon: <CloudServerOutlined />,
      label: <Link to="/">Dashboard</Link>,
    },
    {
      key: '/accounts',
      icon: <UserOutlined />,
      label: <Link to="/accounts">Accounts</Link>,
    },
    {
      key: '/zones',
      icon: <GlobalOutlined />,
      label: <Link to="/zones">Zones</Link>,
    },
    {
      key: '/dns',
      icon: <ApiOutlined />,
      label: <Link to="/dns">DNS Records</Link>,
    },
    {
      key: '/workers',
      icon: <SettingOutlined />,
      label: <Link to="/workers">Workers</Link>,
    },
    {
      key: '/pages',
      icon: <FileTextOutlined />,
      label: <Link to="/pages">Pages</Link>,
    },
    {
      key: '/kv',
      icon: <DatabaseOutlined />,
      label: <Link to="/kv">KV Storage</Link>,
    },
    {
      key: '/r2',
      icon: <HddOutlined />,
      label: <Link to="/r2">R2 Storage</Link>,
    },
  ];

  return (
    <Router>
      <Layout style={{ minHeight: '100vh' }}>
        <Sider collapsible collapsed={collapsed} onCollapse={setCollapsed}>
          <div style={{ height: 32, margin: 16, background: 'rgba(255, 255, 255, 0.2)', borderRadius: 6, display: 'flex', alignItems: 'center', justifyContent: 'center', color: 'white', fontWeight: 'bold' }}>
            {collapsed ? 'CF' : 'Cloudflare Manager'}
          </div>
          <Menu theme="dark" defaultSelectedKeys={['/']} mode="inline" items={menuItems} />
        </Sider>
        <Layout>
          <Header style={{ padding: 0, background: colorBgContainer }}>
            <div style={{ padding: '0 24px', fontSize: 20, fontWeight: 'bold' }}>
              🚀 Cloudflare Multi-Account Manager
            </div>
          </Header>
          <Content style={{ margin: '16px' }}>
            <div
              style={{
                padding: 24,
                minHeight: 360,
                background: colorBgContainer,
                borderRadius: borderRadiusLG,
              }}
            >
              <Routes>
                <Route path="/" element={<Dashboard />} />
                <Route path="/accounts" element={<Accounts />} />
                <Route path="/zones" element={<Zones />} />
                <Route path="/dns" element={<DNS />} />
                <Route path="/workers" element={<Workers />} />
                <Route path="/pages" element={<Pages />} />
                <Route path="/kv" element={<KV />} />
                <Route path="/r2" element={<R2 />} />
              </Routes>
            </div>
          </Content>
          <Footer style={{ textAlign: 'center' }}>
            Cloudflare Manager ©{new Date().getFullYear()} - Powered by React + FastAPI + Go
          </Footer>
        </Layout>
      </Layout>
    </Router>
  );
}

export default App;
