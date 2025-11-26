import React, { useEffect, useState } from 'react';
import { Table, Button, Modal, Form, Input, message, Space, Popconfirm, Upload, Card } from 'antd';
import { PlusOutlined, DeleteOutlined, RocketOutlined, UploadOutlined, ReloadOutlined } from '@ant-design/icons';
import axios from 'axios';

function Pages() {
  const [projects, setProjects] = useState([]);
  const [loading, setLoading] = useState(false);
  const [isCreateModalVisible, setIsCreateModalVisible] = useState(false);
  const [isDeployModalVisible, setIsDeployModalVisible] = useState(false);
  const [selectedProject, setSelectedProject] = useState(null);
  const [createForm] = Form.useForm();
  const [deployForm] = Form.useForm();
  const [fileList, setFileList] = useState([]);

  useEffect(() => {
    fetchProjects();
  }, []);

  const fetchProjects = async () => {
    setLoading(true);
    try {
      const response = await axios.get('/api/pages');
      if (response.data.success && response.data.result) {
        setProjects(response.data.result);
      }
    } catch (error) {
      message.error(error.response?.data?.detail || '获取项目列表失败');
    } finally {
      setLoading(false);
    }
  };

  const handleCreate = async (values) => {
    try {
      const response = await axios.post('/api/pages', {
        name: values.name,
        production_branch: values.production_branch || 'main'
      });
      
      if (response.data.success) {
        message.success('项目创建成功！');
        setIsCreateModalVisible(false);
        createForm.resetFields();
        fetchProjects();
      }
    } catch (error) {
      message.error(error.response?.data?.detail || '创建项目失败');
    }
  };

  const handleDeploy = async (values) => {
    if (!selectedProject) return;
    
    if (fileList.length === 0) {
      message.error('请选择要部署的 Worker 文件');
      return;
    }

    try {
      const formData = new FormData();
      formData.append('branch', values.branch || 'main');
      formData.append('worker_file', fileList[0].originFileObj);

      const response = await axios.post(
        `/api/pages/${selectedProject}/deployments`,
        formData,
        {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        }
      );

      if (response.data.success) {
        message.success('部署成功！');
        const deployment = response.data.result;
        Modal.success({
          title: '部署完成',
          content: (
            <div>
              <p><strong>URL:</strong> <a href={deployment.url} target="_blank" rel="noopener noreferrer">{deployment.url}</a></p>
              <p><strong>ID:</strong> {deployment.id}</p>
            </div>
          ),
        });
        setIsDeployModalVisible(false);
        deployForm.resetFields();
        setFileList([]);
      }
    } catch (error) {
      message.error(error.response?.data?.detail || '部署失败');
    }
  };

  const handleDelete = async (projectName) => {
    try {
      await axios.delete(`/api/pages/${projectName}`);
      message.success('项目删除成功');
      fetchProjects();
    } catch (error) {
      message.error(error.response?.data?.detail || '删除项目失败');
    }
  };

  const uploadProps = {
    accept: '.js',
    beforeUpload: (file) => {
      const isJS = file.name.endsWith('.js');
      if (!isJS) {
        message.error('只能上传 .js 文件！');
        return false;
      }
      setFileList([file]);
      return false;
    },
    onRemove: () => {
      setFileList([]);
    },
    fileList,
  };

  const columns = [
    {
      title: '项目名称',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: '生产分支',
      dataIndex: 'production_branch',
      key: 'production_branch',
    },
    {
      title: '创建时间',
      dataIndex: 'created_on',
      key: 'created_on',
      render: (date) => date ? new Date(date).toLocaleString() : '-',
    },
    {
      title: '域名',
      dataIndex: 'subdomain',
      key: 'subdomain',
      render: (subdomain) => subdomain ? (
        <a href={`https://${subdomain}.pages.dev`} target="_blank" rel="noopener noreferrer">
          {subdomain}.pages.dev
        </a>
      ) : '-',
    },
    {
      title: '操作',
      key: 'actions',
      render: (_, record) => (
        <Space>
          <Button
            type="primary"
            icon={<RocketOutlined />}
            onClick={() => {
              setSelectedProject(record.name);
              setIsDeployModalVisible(true);
            }}
          >
            部署
          </Button>
          <Popconfirm
            title="确定要删除这个项目吗？"
            onConfirm={() => handleDelete(record.name)}
            okText="确定"
            cancelText="取消"
          >
            <Button danger icon={<DeleteOutlined />}>
              删除
            </Button>
          </Popconfirm>
        </Space>
      ),
    },
  ];

  return (
    <div>
      <Card>
        <div style={{ marginBottom: 16, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <h1 style={{ margin: 0 }}>Cloudflare Pages</h1>
          <Space>
            <Button icon={<ReloadOutlined />} onClick={fetchProjects}>
              刷新
            </Button>
            <Button type="primary" icon={<PlusOutlined />} onClick={() => setIsCreateModalVisible(true)}>
              创建项目
            </Button>
          </Space>
        </div>

        <Table
          columns={columns}
          dataSource={projects}
          loading={loading}
          rowKey="name"
          pagination={{ pageSize: 10 }}
        />
      </Card>

      <Modal
        title="创建 Pages 项目"
        open={isCreateModalVisible}
        onCancel={() => setIsCreateModalVisible(false)}
        footer={null}
      >
        <Form form={createForm} onFinish={handleCreate} layout="vertical">
          <Form.Item
            name="name"
            label="项目名称"
            rules={[
              { required: true, message: '请输入项目名称!' },
              { pattern: /^[a-z0-9-]+$/, message: '只能包含小写字母、数字和连字符' }
            ]}
          >
            <Input placeholder="my-project" />
          </Form.Item>
          
          <Form.Item
            name="production_branch"
            label="生产分支"
            initialValue="main"
          >
            <Input placeholder="main" />
          </Form.Item>
          
          <Form.Item>
            <Space>
              <Button type="primary" htmlType="submit">
                创建
              </Button>
              <Button onClick={() => setIsCreateModalVisible(false)}>取消</Button>
            </Space>
          </Form.Item>
        </Form>
      </Modal>

      <Modal
        title={`部署到 ${selectedProject}`}
        open={isDeployModalVisible}
        onCancel={() => {
          setIsDeployModalVisible(false);
          setFileList([]);
          deployForm.resetFields();
        }}
        footer={null}
        width={600}
      >
        <Form form={deployForm} onFinish={handleDeploy} layout="vertical">
          <Form.Item
            name="branch"
            label="分支"
            initialValue="main"
          >
            <Input placeholder="main" />
          </Form.Item>
          
          <Form.Item
            label="Worker 文件"
            required
          >
            <Upload {...uploadProps}>
              <Button icon={<UploadOutlined />}>选择 _worker.js 文件</Button>
            </Upload>
            <div style={{ marginTop: 8, color: '#666', fontSize: 12 }}>
              上传您的 Worker 脚本文件（_worker.js）
            </div>
          </Form.Item>
          
          <Form.Item>
            <Space>
              <Button type="primary" htmlType="submit" icon={<RocketOutlined />}>
                开始部署
              </Button>
              <Button onClick={() => {
                setIsDeployModalVisible(false);
                setFileList([]);
                deployForm.resetFields();
              }}>
                取消
              </Button>
            </Space>
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
}

export default Pages;
