#coding:utf8
import time

from pymysql import *


def main():
    # 创建connection连接
    conn = connect(host='127.0.0.1', port=3306, database='ffly', user='root',
                   password='root', charset='utf8')
    # 获取cursor对象
    cs1 = conn.cursor()
    # 执行sql语句
    now = int(time.time())
    values = []
    # password3
    passwd = "$2a$12$zRMOB..dDYKyilDBi70r2e.wvTX14Ru/pq/GRWAYax3sE5Lgucu5u"
    for i in range(10,1000020):
       
        query = 'insert into users(created_at, updated_at, deleted_at, \
        is_delete, user_name, password_digest, nickname, status, avatar) values(%s, %s, %s, %s, %s, %s, \
            %s, %s, %s)'
     
       
        value = ("{}".format(now), "{}".format(now), 0, 0,\
            "user_name_{}".format(i), passwd,  \
                "nickname_{}".format(i), "active", "active")
        values.append(value)
        if len(values) > 500:
            print i
            cs1.executemany(query, values)
            values = []
            conn.commit()
    if values:
         cs1.executemany(query, values)
    # 提交之前的操作，如果之前已经执行多次的execute，那么就都进行提交
    conn.commit()

    # 关闭cursor对象
    cs1.close()
    # 关闭connection对象
    conn.close()


if __name__ == '__main__':
    main()