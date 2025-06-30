-- insert语句不能换行
-- 但是为什么还是插入失败？ 算了先跳过吧，自己手动插入一些数据

-- -- 插入顶级菜单
-- INSERT INTO `menu` (`id`, `rule_id`, `status`, `name`, `desc`, `menu`, `order`, `icon`, `method`) VALUES (1, 0, 1, '后台面板', 'index', 1, 1, 'help', 'GET'),(2, 0, 1, '商品管理', 'goods', 1, 2, 'goods', 'GET');

-- -- 插入子菜单
-- INSERT INTO `menu` (`id`, `rule_id`, `status`, `name`, `desc`, `frontpath`, `front_file_path`, `menu`, `order`, `icon`, `method`) VALUES (10, 1, 1, '主控台', 'index', '/', '../page/dashboard/index.vue', 1, 20, 'home-filled', 'GET');

-- INSERT INTO `menu` (`id`, `rule_id`, `status`, `name`, `desc`, `frontpath`, `front_file_path`, `menu`, `order`, `icon`, `method`) VALUES (11, 1, 1, '个人设置', 'settings', '/settings', '../page/settings/index.vue', 1, 21, 'setting', 'GET');

-- INSERT INTO `menu` (`id`, `rule_id`, `status`, `name`, `desc`, `frontpath`, `front_file_path`, `menu`, `order`, `icon`, `method`) VALUES (20, 2, 1, '商品列表', 'goods-list', '/goods/list', '../page/goods/list.vue', 1, 30, 'list', 'GET');

-- INSERT INTO `menu` (`id`, `rule_id`, `status`, `name`, `desc`, `frontpath`, `front_file_path`, `menu`, `order`, `icon`, `method`) VALUES (21, 2, 1, '商品分类', 'goods-category', '/goods/category', '../page/goods/category.vue', 1, 31, 'menu', 'GET');

-- 新增二级子菜单：商品规格管理菜单
-- INSERT INTO `menu` (`id`, `rule_id`, `status`, `name`, `desc`, `frontpath`, `front_file_path`, `menu`, `order`, `icon`, `method`) VALUES (22, 2, 1, '规格管理', 'goods-specs', '/goods/specs', '../page/goods/specs.vue', 1, 32, 'files', 'GET');

-- 新增二级子菜单：商品优惠券管理菜单
-- INSERT INTO `menu` (`id`, `rule_id`, `status`, `name`, `desc`, `frontpath`, `front_file_path`, `menu`, `order`, `icon`, `method`) VALUES (23, 2, 1, '优惠券管理', 'goods-coupons', '/goods/coupons', '../page/goods/coupons.vue', 1, 33, 'ticket', 'GET');