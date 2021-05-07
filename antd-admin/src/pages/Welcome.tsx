import React from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { Card, Typography } from 'antd';

const CodePreview: React.FC = ({ children }) => (
  <pre>
    <code>
      <Typography.Text copyable>{children}</Typography.Text>
    </code>
  </pre>
);

export default (): React.ReactNode => {
  return (
    <PageContainer>
      <Card>
        <CodePreview>Hello World!</CodePreview>
      </Card>
    </PageContainer>
  );
};
