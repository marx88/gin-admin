package businesstype

// Enum 枚举
type Enum int

const (
	// OTHER 其它
	OTHER Enum = iota

	// INSERT 新增
	INSERT

	// UPDATE 修改
	UPDATE

	// DELETE 删除
	DELETE

	// GRANT 授权
	GRANT

	// EXPORT 导出
	EXPORT

	// IMPORT 导入
	IMPORT

	// FORCE 强退
	FORCE

	// GENCODE 生成代码
	GENCODE

	// CLEAN 清空数据
	CLEAN
)

func (e Enum) String() string {
	switch e {
	case OTHER:
		return "其它"
	case INSERT:
		return "新增"
	case UPDATE:
		return "修改"
	case DELETE:
		return "删除"
	case GRANT:
		return "授权"
	case EXPORT:
		return "导出"
	case IMPORT:
		return "导入"
	case FORCE:
		return "强退"
	case GENCODE:
		return "生成代码"
	case CLEAN:
		return "清空数据"
	default:
		return "未知类型"
	}
}
