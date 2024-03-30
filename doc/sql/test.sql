insert into baker_auth_resource(name, resource_type, status, url, parent_id, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('图形验证码登陆', 2, 1, '/user/login/byCaptcha', 0, '2', '', false, false, '', '');

insert into baker_auth_resource(name, resource_type, status, url, parent_id, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('获取指定业务类型的树形结构', 2, 1, '/auth/admin/getTreeByBizType', 0, '2', '', false, false, '', '');

insert into baker_auth_resource(name, resource_type, status, url, parent_id, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('树形结构表搜索', 2, 1, '/auth/admin/searchTreeNode', 0, '2', '', false, false, 'trpc.blog.auth.AuthApiHttp', '');

insert into baker_auth_resource(name, resource_type, status, url, parent_id, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('添加修改树形结构表', 2, 1, '/auth/admin/addOrUpdateTreeNode', 0, '2', '', false, false, 'trpc.blog.auth.AuthApiHttp', '');


insert into baker_auth_resource(name, resource_type, status, url, parent_id, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('删除树形结构表', 2, 1, '/auth/admin/deleteTreeNode', 0, '2', '', false, false, 'trpc.blog.auth.AuthApiHttp', '');


insert into baker_auth_resource(name, resource_type, status, url, parent_id, http_method, `describe`, is_need_login, is_need_permission, service, method)
values ('树形结构表详情', 2, 1, '/auth/admin/treeNodeDetail', 0, '2', '', false, false, 'trpc.blog.auth.AuthApiHttp', '')

