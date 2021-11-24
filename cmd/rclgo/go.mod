module github.com/mehmetkillioglu/rclgo/cmd/rclgo

go 1.16

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/mehmetkillioglu/rclgo v0.0.0-20210608082116-2a9c702d8fdc
	github.com/mehmetkillioglu/rclgo-msgs v0.0.0-20210608113630-1716fa950d7b
)

replace github.com/mehmetkillioglu/rclgo => ../..
