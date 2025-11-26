import React, { useEffect, useState } from 'react';
import { Table, Button, Modal, Form, Input, message, Space, Popconfirm } from 'antd';
import { PlusOutlined, DeleteOutlined, SwapOutlined } from '@ant-design/icons';
import axios from 'axios';

function Accounts() {
  const [accounts, setAccounts] = useState([]);
  const [loading, setLoading] = useState(false);
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [form] = Form.useForm();

  useEffect(() => {
    fetchAccounts();
  }, []);

  const fetchAccounts = async () => {
    setLoading(true);
    try {
      const response = await axios.get('/api/accounts');
      // 解析命令输出（简化处理）
      setAccounts([]);
    } catch (error) {
      message.error('Failed to fetch accounts');
    } finally {
      setLoading(false);
    }
  };

  const handleAdd = async (values) => {
    try {
      await axios.post('/api/accounts', values);
      message.success('Account added successfully');
      setIsModalVisible(false);
      form.resetFields();
      fetchAccounts();
    } catch (error) {
      message.error('Failed to add account');
    }
  };

  const handleSwitch = async (name) => {
    try {
      await axios.post(`/api/accounts/${name}/switch`);
      message.success(`Switched to account: ${name}`);
      fetchAccounts();
    } catch (error) {
      message.error('Failed to switch account');
    }
  };

  const handleDelete = async (name) => {
    try {
      await axios.delete(`/api/accounts/${name}`);
      message.success('Account deleted successfully');
      fetchAccounts();
    } catch (error) {
      message.error('Failed to delete account');
    }
  };

  const columns = [
    {
      title: 'Current',
      dataIndex: 'current',
      key: 'current',
      width: 80,
      render: (current) => current ? '⭐' : '',
    },
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: 'Email',
      dataIndex: 'email',
      key: 'email',
    },
    {
      title: 'Account ID',
      dataIndex: 'account_id',
      key: 'account_id',
    },
    {
      title: 'Actions',
      key: 'actions',
      render: (_, record) => (
        <Space>
          <Button
            type="primary"
            icon={<SwapOutlined />}
            onClick={() => handleSwitch(record.name)}
            disabled={record.current}
          >
            Switch
          </Button>
          <Popconfirm
            title="Delete this account?"
            onConfirm={() => handleDelete(record.name)}
            okText="Yes"
            cancelText="No"
          >
            <Button danger icon={<DeleteOutlined />}>
              Delete
            </Button>
          </Popconfirm>
        </Space>
      ),
    },
  ];

  return (
    <div>
      <div style={{ marginBottom: 16, display: 'flex', justifyContent: 'space-between' }}>
        <h1>Accounts Management</h1>
        <Button type="primary" icon={<PlusOutlined />} onClick={() => setIsModalVisible(true)}>
          Add Account
        </Button>
      </div>

      <Table
        columns={columns}
        dataSource={accounts}
        loading={loading}
        rowKey="name"
      />

      <Modal
        title="Add New Account"
        open={isModalVisible}
        onCancel={() => setIsModalVisible(false)}
        footer={null}
      >
        <Form form={form} onFinish={handleAdd} layout="vertical">
          <Form.Item
            name="name"
            label="Account Name"
            rules={[{ required: true, message: 'Please input account name!' }]}
          >
            <Input placeholder="e.g., production" />
          </Form.Item>
          <Form.Item
            name="api_token"
            label="API Token"
            rules={[{ required: true, message: 'Please input API token!' }]}
          >
            <Input.Password placeholder="Your Cloudflare API token" />
          </Form.Item>
          <Form.Item name="email" label="Email (Optional)">
            <Input placeholder="your@email.com" />
          </Form.Item>
          <Form.Item>
            <Space>
              <Button type="primary" htmlType="submit">
                Add Account
              </Button>
              <Button onClick={() => setIsModalVisible(false)}>Cancel</Button>
            </Space>
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
}

export default Accounts;
