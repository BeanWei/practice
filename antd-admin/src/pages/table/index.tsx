import React, { useRef, useState } from 'react';
import {
  FormDrawer,
  FormItem,
  FormGrid,
  Input,
  Submit,
  Reset,
  Select,
  ArrayItems,
  FormButtonGroup,
} from '@formily/antd';
import { createSchemaField } from '@formily/react'
import { PlusOutlined } from '@ant-design/icons';
import { Button, Tag, Space } from 'antd';
import { PageContainer } from '@ant-design/pro-layout';
import type { ActionType, ProTableProps } from '@ant-design/pro-table';
import ProTable from '@ant-design/pro-table';
import request from 'umi-request';

const SchemaField = createSchemaField({
  components: {
    FormItem,
    Input,
    Select,
    ArrayItems
  },
});

const jsonschema = {
  type: 'object',
  properties: {
    url: {
      type: 'string',
      title: '链接',
      required: false,
      'x-decorator': 'FormItem',
      'x-component': 'Input',
    },
    title: {
      type: 'string',
      title: '标题',
      required: true,
      'x-decorator': 'FormItem',
      'x-component': 'Input',
      'x-list-props': {
        copyable: true,
        ellipsis: true,
        tip: '标题过长会自动收缩',
      },
    },
    state: {
      type: 'string',
      title: '状态',
      required: false,
      'x-decorator': 'FormItem',
      'x-component': 'Select',
      enum: [
        { label: '未解决', value: 'open' },
        { label: '已解决', value: 'closed' },
        { label: '处理中', value: 'processing' },
      ],
      'x-list-props': {
        filters: true,
        onFilter: true,
      },
    },
    labels: {
      type: 'array',
      title: '标签',
      'x-component': 'ArrayItems',
      'x-decorator': 'FormItem',
      'x-list-props': {
        search: false,
        renderFormItem: (_, { defaultRender }) => {
          return defaultRender(_);
        },
        render: (_, record) => (
          <Space>
            {record.labels.map(({ name, color }) => (
              <Tag color={color} key={name}>
                {name}
              </Tag>
            ))}
          </Space>
        ),
      },
      items: {
        type: 'object',
        properties: {
          sort: {
            type: 'void',
            'x-decorator': 'FormItem',
            'x-component': 'ArrayItems.SortHandle',
          },
          name: {
            type: 'string',
            title: '名字',
            'x-decorator': 'FormItem',
            'x-component': 'Input',
          },
          color: {
            type: 'string',
            title: '颜色',
            'x-decorator': 'FormItem',
            'x-component': 'Input',
          },
          remove: {
            type: 'void',
            'x-decorator': 'FormItem',
            'x-component': 'ArrayItems.Remove',
          },
        },
      },
      properties: {
        add: {
          type: 'void',
          title: '添加条目',
          'x-component': 'ArrayItems.Addition',
        },
      },
    },
  },
}

const getColumns = (data: Record<string, unknown>) => {
  const valMap = {
    'Select': 'select',
    'Input': 'text',
  }

  const cols = [];
  // eslint-disable-next-line no-restricted-syntax
  for (const [name, field] of Object.entries(data.properties)) {
    const props = field['x-list-props'];
    if (props) {
      const col = {
        dataIndex: name,
        title: field.title,
        valueType: valMap[field['x-component']],
        valueEnum: field.enum ? field.enum.reduce((obj, item) => ({...obj, [item.value]: item.label}), {}) : undefined,
        ...props,
      }
      cols.push(col);
    }
  }
  console.log(cols)
  return cols;
}

const renderForm = (title: string, schema: Record<string, unknown>, callback: any, initialValues?: Record<string, unknown>) => {
  FormDrawer({
    title, width: '60%',
  }, (resolve) => {
    return (
      <FormGrid maxColumns={3}>
        <SchemaField schema={schema} />
        <FormDrawer.Footer>
          <FormButtonGroup align="right">
            <Submit onClick={resolve}>提交</Submit>
            <Reset>重置</Reset>
          </FormButtonGroup>
        </FormDrawer.Footer>
      </FormGrid>
    )
  })
    .open({
      initialValues
    })
    .then(callback)
}

export type FormTableProps<RecordType> = Omit<ProTableProps<RecordType>> & {
  schema: Record<string, unknown>;
  createRequest?: (record: Record<string, unknown>) => void | undefined;
  updateRequest?: (record: Record<string, unknown>) => void | undefined;
  deleteRequest?: (record: Record<string, unknown>) => void | undefined;
}

function FormTable(props: FormTableProps) {
  const columns = getColumns(props.schema);

  if (props.updateRequest || props.deleteRequest) {
    columns.push({
      title: '操作',
      dataIndex: '__option',
      valueType: 'option',
      render: (_, record) => {
        const actions = [];
        if (props.updateRequest) {
          actions.push(
            <a
              key="edit"
              onClick={() => {
                renderForm('更新', props.schema, props.updateRequest, record)
              }}
            >
              编辑
            </a>,
          )
        }
        if (props.deleteRequest) {
          actions.push(
            <a
              key="delete"
              onClick={() => {
                props.deleteRequest(record)
              }}
            >
              删除
            </a>,
          )
        }
        return actions;
      }
    })
  }

  const toolBars = props.toolBarRender ? toolBarRender() : [];
  if (props.createRequest) {
    toolBars.push(
      <Button key="button" icon={<PlusOutlined />} type="primary" onClick={() => {
        renderForm("新建", props.schema, props.createRequest)
      }}>
        新建
      </Button>,
    )
  }

  return (
    <ProTable
      {...props}
      columns={columns}
      toolBarRender={() => toolBars}
    />
  )
}

export default () => {
  const actionRef = useRef<ActionType>();
  return (
    <PageContainer>
      <FormTable
        schema={jsonschema}
        actionRef={actionRef}
        request={async (params = {}, sort, filter) => {
          console.log(sort, filter);
          return request<{
            data: GithubIssueItem[];
          }>('https://proapi.azurewebsites.net/github/issues', {
            params,
          });
        }}
        editable={{
          type: 'multiple',
        }}
        rowKey="id"
        search={{
          labelWidth: 'auto',
        }}
        form={{
          // 由于配置了 transform，提交的参与与定义的不同这里需要转化一下
          syncToUrl: (values, type) => {
            if (type === 'get') {
              return {
                ...values,
                created_at: [values.startTime, values.endTime],
              };
            }
            return values;
          },
        }}
        pagination={{
          pageSize: 5,
        }}
        dateFormatter="string"
        headerTitle="高级表格"
        createRequest={
          (record) => {
            console.log("create", record)
          }
        }
        updateRequest={
          (record) => {
            console.log("update", record)
          }
        }
        deleteRequest={
          (record) => {
            console.log("delete", record)
          }
        }
      />
    </PageContainer>
  );
};
