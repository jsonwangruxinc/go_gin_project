package grom

func init() {
	DB.Migrator().DropTable(Techer{})
	DB.AutoMigrate(Techer{})
}
