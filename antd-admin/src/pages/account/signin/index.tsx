import React, { useState } from 'react';
import { message } from 'antd';
import { useIntl, Link, history, SelectLang, useModel } from 'umi';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import { createForm } from '@formily/core';
import { Field } from '@formily/react';
import { Form, FormItem, Input, Password, Submit } from '@formily/antd';
import Footer from '@/components/Footer';
import api from '@/services/account';

import styles from './index.less';

/** 此方法会跳转到 redirect 参数所在的位置 */
const goto = () => {
  if (!history) return;
  setTimeout(() => {
    const { query } = history.location;
    const { redirect } = query as { redirect: string };
    history.push(redirect || '/');
  }, 10);
};

const Signin: React.FC = () => {
  // const [userSigninState, setUserSigninState] = useState<API.SigninReply>({});
  const { initialState, setInitialState } = useModel('@@initialState');

  const intl = useIntl();

  const fetchUserInfo = async () => {
    const userInfo = await initialState?.fetchUserInfo?.();
    if (userInfo) {
      setInitialState({
        ...initialState,
        currentUser: userInfo,
      });
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.lang}>{SelectLang && <SelectLang />}</div>
      <div className={styles.content}>
        <div className={styles.top}>
          <div className={styles.header}>
            <Link to="/">
              <img alt="logo" className={styles.logo} src="/logo.svg" />
              <span className={styles.title}>Ant Design</span>
            </Link>
          </div>
          <div className={styles.desc}>
            {intl.formatMessage({ id: 'pages.layouts.userLayout.title' })}
          </div>
        </div>

        <div className={styles.main}>
          <Form
            form={createForm()}
            layout="vertical"
            size="large"
          >
            <Field
              name="account"
              title="账号"
              required
              decorator={[FormItem]}
              component={[
                Input,
                {
                  prefix: <UserOutlined />,
                },
              ]}
            />
            <Field
              name="password"
              title="密码"
              required
              decorator={[FormItem]}
              component={[
                Password,
                {
                  prefix: <LockOutlined />,
                },
              ]}
            />
            <Submit
              block
              size="large"
              onSubmit={async (values) => {
                try {
                  // 登录
                  await api.Signin({ ...values });
                  const defaultsigninSuccessMessage = intl.formatMessage({
                    id: 'pages.signin.success',
                    defaultMessage: '登录成功！',
                  });
                  message.success(defaultsigninSuccessMessage);
                  await fetchUserInfo();
                  goto();
                  return;
                } catch (error) {
                  message.error(error.data.message);
                }
              }}
            >
              登录
            </Submit>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default Signin;
