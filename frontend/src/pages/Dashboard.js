import React, { useEffect, useState } from 'react';
import { Card, Row, Col, Statistic, Alert } from 'antd';
import {
  CloudServerOutlined,
  GlobalOutlined,
  ApiOutlined,
  DatabaseOutlined,
  HddOutlined,
  FileTextOutlined,
} from '@ant-design/icons';
import axios from 'axios';

function Dashboard() {
  const [stats, setStats] = useState({
    zones: 0,
    workers: 0,
    pages: 0,
    kv: 0,
    r2: 0
  });
  const [loading, setLoading] = useState(false);
  const [currentAccount, setCurrentAccount] = useState(null);

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    setLoading(true);
    try {
      // 获取当前账号信息
      const accountRes = await axios.get('/api/accounts/current');
      setCurrentAccount(accountRes.data);

      // 获取统计数据（这里简化处理）
      // 实际应用中应该并行请求所有API
      setStats({
        zones: 0,
        workers: 0,
        pages: 0,
        kv: 0,
        r2: 0
      });
    } catch (error) {
      console.error('Failed to fetch data:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <h1>Dashboard</h1>
      
      {currentAccount && (
        <Alert
          message="Current Account"
          description={`You are currently managing: ${currentAccount.name || 'Unknown'}`}
          type="info"
          showIcon
          style={{ marginBottom: 24 }}
        />
      )}

      <Row gutter={[16, 16]}>
        <Col xs={24} sm={12} lg={8}>
          <Card>
            <Statistic
              title="Zones"
              value={stats.zones}
              prefix={<GlobalOutlined />}
              valueStyle={{ color: '#3f8600' }}
            />
          </Card>
        </Col>
        <Col xs={24} sm={12} lg={8}>
          <Card>
            <Statistic
              title="Workers"
              value={stats.workers}
              prefix={<CloudServerOutlined />}
              valueStyle={{ color: '#cf1322' }}
            />
          </Card>
        </Col>
        <Col xs={24} sm={12} lg={8}>
          <Card>
            <Statistic
              title="Pages Projects"
              value={stats.pages}
              prefix={<FileTextOutlined />}
              valueStyle={{ color: '#1890ff' }}
            />
          </Card>
        </Col>
        <Col xs={24} sm={12} lg={8}>
          <Card>
            <Statistic
              title="KV Namespaces"
              value={stats.kv}
              prefix={<DatabaseOutlined />}
              valueStyle={{ color: '#722ed1' }}
            />
          </Card>
        </Col>
        <Col xs={24} sm={12} lg={8}>
          <Card>
            <Statistic
              title="R2 Buckets"
              value={stats.r2}
              prefix={<HddOutlined />}
              valueStyle={{ color: '#fa8c16' }}
            />
          </Card>
        </Col>
        <Col xs={24} sm={12} lg={8}>
          <Card>
            <Statistic
              title="DNS Records"
              value={0}
              prefix={<ApiOutlined />}
              valueStyle={{ color: '#52c41a' }}
            />
          </Card>
        </Col>
      </Row>

      <Card title="Quick Actions" style={{ marginTop: 24 }}>
        <p>Welcome to Cloudflare Manager! Use the sidebar to navigate between different services.</p>
        <p>This dashboard provides an overview of your Cloudflare resources across all managed accounts.</p>
      </Card>
    </div>
  );
}

export default Dashboard;
