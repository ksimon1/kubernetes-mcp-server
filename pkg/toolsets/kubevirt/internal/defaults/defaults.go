package defaults

const (
	DefaultToolsetName        = "kubevirt"
	DefaultToolsetDescription = "KubeVirt virtual machine management tools, check the [KubeVirt documentation](https://github.com/containers/kubernetes-mcp-server/blob/main/docs/kubevirt.md) for more details."
)

func ToolsetName() string {
	overrideName := ToolsetNameOverride()
	if overrideName != "" {
		return overrideName
	}
	return DefaultToolsetName
}

func ToolsetDescription() string {
	overrideDescription := ToolsetDescriptionOverride()
	if overrideDescription != "" {
		return overrideDescription
	}
	return DefaultToolsetDescription
}
