package cmd

import (
	"github.com/spf13/cobra"
)

var (
	LoginSwitch    = false
	RegisterSwitch = false
	UploadSwitch   = false
	FilesSwitch    = false
	DownloadSwitch = false
	LgUserName     string
	RgUserName     string
	LgPassWord     string
	RgPassWord     string
	FileId         string
	FileAbsPath    string
	PageNum        string
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "this is login command",
	Long: `this is login sub command :
	-u   username string
	-p   password string`,
	Run: func(cmd *cobra.Command, args []string) {
		LoginSwitch = true
	},
}
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "this is register command",
	Long: `this is register sub command :
	-u   username string
	-p   password string`,
	Run: func(cmd *cobra.Command, args []string) {
		RegisterSwitch = true
	},
}
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "this is upload command",
	Long: `this is upload sub command :
	-p     this is file abs path`,
	Run: func(cmd *cobra.Command, args []string) {
		UploadSwitch = true
	},
}
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "this is download command",
	Long: `this is download sub command :
	-i  file id`,
	Run: func(cmd *cobra.Command, args []string) {
		DownloadSwitch = true
	},
}
var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "this is file command",
	Long: `this is file sub command :
	-p   pageNum string`,
	Run: func(cmd *cobra.Command, args []string) {
		FilesSwitch = true
	},
}

func init() {
	loginCmd.Flags().StringVarP(&LgUserName, "user", "u", "", "this is login user_name flag")
	loginCmd.Flags().StringVarP(&LgPassWord, "password", "p", "", "this is login  password flag")
	registerCmd.Flags().StringVarP(&RgUserName, "user", "u", "", "this is register user_name flag")
	registerCmd.Flags().StringVarP(&RgPassWord, "password", "p", "", "this is register password flag")
	uploadCmd.Flags().StringVarP(&FileAbsPath, "path", "p", "", "this is upload path flag")
	filesCmd.Flags().StringVarP(&PageNum, "page", "p", "", "this is files list page num flag")
	downloadCmd.Flags().StringVarP(&FileId, "id", "i", "", "this is download file id flag")

	rootCmd.AddCommand(loginCmd, uploadCmd, registerCmd, downloadCmd, filesCmd)
}
