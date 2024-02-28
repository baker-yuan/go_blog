package menu_handles

import (
	middleware "github.com/baker-yuan/go-blog/all_packaged_library/middleware/gin"
	"github.com/baker-yuan/go-blog/auth/application/dto"
	"github.com/baker-yuan/go-blog/auth/application/service"
	"github.com/gin-gonic/gin"
)

// MenuUI 菜单控制器
type MenuUI interface {
	ListUserMenus(ginCtx *gin.Context)
}
type MenuUIImpl struct {
	menuService *service.MenuService
}

// ListUserMenus godoc
//
//	@Summary		查看当前用户菜单
//	@Description	查看当前用户菜单
//	@Tags			文章模块
//	@Accept			json
//	@Produce		json
//	@Param			articleId	path		int	true	"文章ID"
//	@Success		200			{object}	middleware.Response{data=dto.ArticleDTO}
//	@Router			/admin/user/menus [get]
func (m MenuUIImpl) ListUserMenus(ginCtx *gin.Context) {
	var (
		userMenus []*dto.UserMenuDTO
		err       error
	)
	userMenus, err = m.menuService.ListUserMenus(1)
	middleware.SendResult(ginCtx, userMenus, err)
}

//  /**
//     * 查询菜单列表
//     *
//     * @param conditionVO 条件
//     * @return {@link Result<MenuDTO>} 菜单列表
//     */
//    @ApiOperation(value = "查看菜单列表")
//    @GetMapping("/admin/menus")
//    public Result<List<MenuDTO>> listMenus(ConditionVO conditionVO) {
//        return Result.ok(menuService.listMenus(conditionVO));
//    }
//
//    /**
//     * 新增或修改菜单
//     *
//     * @param menuVO 菜单
//     * @return {@link Result<>}
//     */
//    @ApiOperation(value = "新增或修改菜单")
//    @PostMapping("/admin/menus")
//    public Result<?> saveOrUpdateMenu(@Valid @RequestBody MenuVO menuVO) {
//        menuService.saveOrUpdateMenu(menuVO);
//        return Result.ok();
//    }
//
//    /**
//     * 删除菜单
//     *
//     * @param menuId 菜单id
//     * @return {@link Result<>}
//     */
//    @ApiOperation(value = "删除菜单")
//    @DeleteMapping("/admin/menus/{menuId}")
//    public Result<?> deleteMenu(@PathVariable("menuId") Integer menuId){
//        menuService.deleteMenu(menuId);
//        return Result.ok();
//    }
//
//    /**
//     * 查看角色菜单选项
//     *
//     * @return {@link Result<LabelOptionDTO>} 查看角色菜单选项
//     */
//    @ApiOperation(value = "查看角色菜单选项")
//    @GetMapping("/admin/role/menus")
//    public Result<List<LabelOptionDTO>> listMenuOptions() {
//        return Result.ok(menuService.listMenuOptions());
//    }
//
