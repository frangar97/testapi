# How to run it

- First create the database in postgresql due to gorm works with databases that are already created in order to run the migrations.
- Change .env variables for your postgresql configuration.
- If you have Make in your computer run `make start_api` else run `go run ./cmd/api`.

# Considerations

- The project was made using the latest golang version at the moment which is `1.21.2`.
- When sending the release_date atrribute for Firmware Entity it has to be in the format `yyyy-MM-dd` due to golang formatting.
- When deleting a device the firmwares associated to the device will be deleted due to the cascade strategy.
