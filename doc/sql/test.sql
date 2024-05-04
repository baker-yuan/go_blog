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
        (331, 200, '动态配置', 1),
        (400, 0, 'baker-blog-pem', 1),
        (401, 400, '菜单权限模块', 1),
        (420, 400, '角色模块', 1),
        (430, 400, '资源模块', 1),
        (450, 400, '用户角色资源维护', 1),
        (460, 400, '目录树', 1),
        (500, 0, 'baker-blog-ics', 1),
        (501, 500, '用户评论模块', 1),
        (520, 500, '评论管理', 1);


select id, tree_node_id, name, `describe`, status, url, http_method, is_need_login, is_need_permission, service, method
from baker_auth_resource;

truncate table go_blog_test.baker_auth_resource;
insert into go_blog_test.baker_auth_resource (id, tree_node_id, name, describe, status, url, http_method, is_need_login, is_need_permission, service, method)
values  (1, 2, '图形验证码登陆', '图形验证码登陆', 1, '/user/login/byCaptcha', 2, 0, 0, 'trpc.blog.user.UserApiHttp', '/go_blog.user.LoginApi/Login'),
        (2, 460, '获取指定业务类型的树形结构', '获取指定业务类型的树形结构', 1, '/auth/admin/getTreeByBizType', 2, 1, 1, 'trpc.blog.auth.AuthApiHttp', '/auth/admin/getTreeByBizType'),
        (3, 460, '树形结构表搜索', '树形结构表搜索', 1, '/auth/admin/searchTreeNode', 2, 1, 1, 'trpc.blog.auth.AuthApiHttp', '/go_blog.auth.AuthApi/SearchTreeNode'),
        (4, 460, '添加修改树形结构表', '添加修改树形结构表', 1, '/auth/admin/addOrUpdateTreeNode', 2, 1, 1, 'trpc.blog.auth.AuthApiHttp', '/go_blog.auth.AuthApi/AddOrUpdateTreeNode'),
        (5, 460, '删除树形结构', '删除树形结构', 1, '/auth/admin/deleteTreeNode', 2, 1, 1, 'trpc.blog.auth.AuthApiHttp', '/go_blog.auth.AuthApi/DeleteTreeNode'),
        (6, 460, '树形结构详情', '树形结构详情', 1, '/auth/admin/treeNodeDetail', 2, 1, 1, 'trpc.blog.auth.AuthApiHttp', '/go_blog.auth.AuthApi/TreeNodeDetail'),
        (7, 401, '菜单树', '菜单树', 1, '/auth/admin/getMenuTree', 2, 1, 1, 'trpc.blog.auth.AuthApiHttp', '/go_blog.auth.AuthApi/GetMenuTree'),
        (8, 430, '接口搜索', '接口搜索', 1, '/auth/admin/searchResource', 2, 1, 1, 'trpc.blog.auth.AuthApiHttp', '/go_blog.auth.AuthApi/SearchResource'),
        (9, 430, '添加修改接口', '添加修改接口', 1, '/auth/admin/addOrUpdateResource', 2, 1, 1, 'trpc.blog.auth.AuthApiHttp', '/go_blog.auth.AuthApi/AddOrUpdateResource'),
        (10, 430, '删除接口', '删除接口', 1, '/auth/admin/deleteResource', 2, 1, 1, 'trpc.blog.auth.AuthApiHttp', '/go_blog.auth.AuthApi/DeleteResource'),
        (11, 430, '接口详情', '接口详情', 1, '/auth/admin/resourceDetail', 2, 1, 1, 'trpc.blog.auth.AuthApiHttp', '/go_blog.auth.AuthApi/ResourceDetail'),
        (12, 401, '菜单搜索', '菜单搜索', 1, '/auth/admin/searchMenu', 2, 1, 1, 'trpc.blog.auth.AuthApiHttp', '/go_blog.auth.AuthApi/SearchMenu'),
        (13, 401, '添加修改菜单', '添加修改菜单', 1, '/auth/admin/addOrUpdateMenu', 2, 1, 1, 'trpc.blog.auth.AuthApiHttp', '/go_blog.auth.AuthApi/AddOrUpdateMenu'),
        (14, 401, '删除菜单', '删除菜单', 1, '/auth/admin/deleteMenu', 2, 1, 1, 'trpc.blog.auth.AuthApiHttp', '/go_blog.auth.AuthApi/DeleteMenu'),
        (15, 401, '菜单详情', '菜单详情', 1, '/auth/admin/menuDetail', 2, 1, 1, 'trpc.blog.auth.AuthApiHttp', '/go_blog.auth.AuthApi/MenuDetail'),
        (17, 430, '网关获取有效的接口信息', '网关获取有效的接口信息', 1, '/auth/effectiveResource', 2, 0, 0, 'trpc.blog.auth.AuthApiHttp', '/go_blog.auth.AuthApi/GetEffectiveResource');


truncate table baker_auth_menu;
insert into baker_auth_menu(id, parent_id, name, icon, `describe`, sort_value, path, component, is_hidden)
values (1, 0, '首页', 'HomeFilled', '', '1', '/home/index', '/home/index', false),
       (100, 0, '权限管理', 'Lock', '', '2', '/auth', '', false),
       (101, 100, '菜单权限', 'Menu', '', '2', '/auth/menu', '/auth/menu/index', false),
       (102, 100, '接口权限', 'Menu', '', '2', '/auth/resource', '/auth/resource/index', false);




INSERT INTO oss_mode_cfg (mode, path, `desc`) VALUES ('avatar', 'avatar/', '头像路径');
INSERT INTO oss_mode_cfg (mode, path, `desc`) VALUES ('article', 'article/', '文章图片路径');
INSERT INTO oss_mode_cfg (mode, path, `desc`) VALUES ('voice', 'voice/', '音频路径');
INSERT INTO oss_mode_cfg (mode, path, `desc`) VALUES ('photos', 'photos/', '相册路径');
INSERT INTO oss_mode_cfg (mode, path, `desc`) VALUES ('config', 'config/', '配置图片路径');
INSERT INTO oss_mode_cfg (mode, path, `desc`) VALUES ('upload', 'upload/', '手动上传');