import React, { useEffect, useState } from 'react';
import { Card, Button, Form, Input, message, Space, Descriptions, Alert } from 'antd';
import { LoginOutlined, LogoutOutlined } from '@ant-design/icons';
import axios from 'axios';

function Accounts() {
  const [currentAccount, setCurrentAccount] = useState(null);
  const [loading, setLoading] = useState(false);
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [form] = Form.useForm();

  useEffect(() => {
    checkAuth();
  }, []);

  const checkAuth = async () => {
    try {
      const response = await axios.get('/api/accounts/current');
      if (response.data.success && response.data.result) {
        setCurrentAccount(response.data.result);
        setIsAuthenticated(true);
      }
    } catch (error) {
      setIsAuthenticated(false);
    }
  };

  const handleLogin = async (values) => {
    setLoading(true);
    try {
      const response = await axios.post('/api/auth', {
        email: values.email,
        api_key: values.api_key
      });
      
      if (response.data.success) {
        message.success('认证成功！');
        setCurrentAccount(response.data.account);
        setIsAuthenticated(true);
        form.resetFields();
      }
    } catch (error) {
      message.error(error.response?.data?.detail || '认证失败');
    } finally {
      setLoading(false);
    }
  };

  const handleLogout = () => {
    setCurrentAccount(null);
    setIsAuthenticated(false);
    message.info('已清除认证信息');
  };

  return (
    <div style={{ maxWidth: 800, margin: '0 auto' }}>
      <h1>账户管理</h1>
      
      <Alert
        message="认证方式说明"
        description="使用 Cloudflare Email 和 API Key 进行认证。您可以在 Cloudflare Dashboard > My Profile > API Tokens > API Keys 中找到 Global API Key。"
        type="info"
        showIcon
        style={{ marginBottom: 24 }}
      />

      {!isAuthenticated ? (
        <Card title="登录 Cloudflare 账户" style={{ marginTop: 16 }}>
          <Form form={form} onFinish={handleLogin} layout="vertical">
            <Form.Item
              name="email"
              label="Email"
              rules={[
                { required: true, message: '请输入邮箱!' },
                { type: 'email', message: '请输入有效的邮箱地址!' }
              ]}
            >
              <Input placeholder="your@email.com" size="large" />
            </Form.Item>
            
            <Form.Item
              name="api_key"
              label="API Key"
              rules={[{ required: true, message: '请输入 API Key!' }]}
            >
              <Input.Password placeholder="您的 Cloudflare Global API Key" size="large" />
            </Form.Item>
            
            <Form.Item>
              <Button 
                type="primary" 
                htmlType="submit" 
                icon={<LoginOutlined />}
                loading={loading}
                size="large"
                block
              >
                登录
              </Button>
            </Form.Item>
          </Form>
        </Card>
      ) : (
        <Card 
          title="当前账户信息" 
          style={{ marginTop: 16 }}
          extra={
            <Button 
              danger 
              icon={<LogoutOutlined />}
              onClick={handleLogout}
            >
              退出登录
            </Button>
          }
        >
          {currentAccount && (
            <Descriptions column={1} bordered>
              <Descriptions.Item label="账户名称">
                {currentAccount.name}
              </Descriptions.Item>
              <Descriptions.Item label="账户 ID">
                {currentAccount.id}
              </Descriptions.Item>
              <Descriptions.Item label="类型">
                {currentAccount.type}
              </Descriptions.Item>
              {currentAccount.created_on && (
                <Descriptions.Item label="创建时间">
                  {new Date(currentAccount.created_on).toLocaleString()}
                </Descriptions.Item>
              )}
            </Descriptions>
          )}
        </Card>
      )}
    </div>
  );
}

export default Accounts;
