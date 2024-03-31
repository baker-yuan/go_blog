use go_blog_test;

truncate table baker_auth_tree_node;
insert into baker_auth_tree_node (id, parent_id, `name`, biz_type)
values  (1, 0, 'baker-blog-user', 1),
        (2, 1, '用户信息模块', 1),
        (100, 0, 'baker-blog-tdp', 1),
        (101, 100, 'OSS模块', 1),
        (200, 0, 'baker-blog-blog', 1),
        (201, 200, '标签管理', 1),
        (210, 200, '分类管理', 1),
        (230, 200, '管理员管理文章', 1),
        (250, 200, '用户查看文章', 1),
        (301, 200, '友链模块', 1),
        (311, 200, '博客信息管理模块', 1),
        (321, 200, '页面管理模块', 1),
        (400, 0, 'baker-blog-pem', 1),
        (401, 400, '菜单权限模块', 1),
        (420, 400, '角色模块', 1),
        (430, 400, '资源模块', 1),
        (450, 400, '用户角色资源维护', 1),
        (500, 0, 'baker-blog-ics', 1),
        (501, 500, '用户评论模块', 1),
        (520, 500, '评论管理', 1);

use go_blog_test;
truncate table baker_auth_resource;
insert into baker_auth_resource(name, status, url, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('图形验证码登陆', 1, '/user/login/byCaptcha', 0, '', false, false, 'trpc.blog.auth.AuthApiHttp', '');

insert into baker_auth_resource(name, status, url, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('获取指定业务类型的树形结构', 1, '/auth/admin/getTreeByBizType', 0, '', false, false, 'trpc.blog.auth.AuthApiHttp', '');

insert into baker_auth_resource(name, status, url, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('树形结构表搜索', 1, '/auth/admin/searchTreeNode', 0, '', false, false, 'trpc.blog.auth.AuthApiHttp', '');

insert into baker_auth_resource(name, status, url, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('添加修改树形结构表', 1, '/auth/admin/addOrUpdateTreeNode', 0, '', false, false, 'trpc.blog.auth.AuthApiHttp', '');


insert into baker_auth_resource(name, status, url, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('删除树形结构表', 1, '/auth/admin/deleteTreeNode', 0, '', false, false, 'trpc.blog.auth.AuthApiHttp', '');


insert into baker_auth_resource(name, status, url, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('树形结构表详情', 1, '/auth/admin/treeNodeDetail', 0, '', false, false, 'trpc.blog.auth.AuthApiHttp', '')

insert into baker_auth_resource(name, status, url, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('菜单树', 1, '/auth/admin/menuTree', 0, '', false, false, 'trpc.blog.auth.AuthApiHttp', '')


insert into baker_auth_resource(name, status, url, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('接口搜索', 1, '/auth/admin/searchResource', 0, '', false, false, 'trpc.blog.auth.AuthApiHttp', '')

insert into baker_auth_resource(name, status, url, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('添加修改接口', 1, '/auth/admin/addOrUpdateResource', 0, '', false, false, 'trpc.blog.auth.AuthApiHttp', '')

insert into baker_auth_resource(name, status, url, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('删除接口', 1, '/auth/admin/deleteResource', 0, '', false, false, 'trpc.blog.auth.AuthApiHttp', '')

insert into baker_auth_resource(name, status, url, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('接口详情', 1, '/auth/admin/resourceDetail', 0, '', false, false, 'trpc.blog.auth.AuthApiHttp', '')

update baker_auth_resource set http_method = 2 where 1 = 1;
