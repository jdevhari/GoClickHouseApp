package main

import "github.com/gofiber/fiber/v2"
import "github.com/gofiber/fiber/v2/middleware/cors"
import "github.com/jdevhari/GoClickHouseApp/pkg/controller"


func main() {
    app := fiber.New()
    
    app.Use(cors.New())

	app.Static("/", "../ui/dist/ui") 

    app.Get("/getTickData", func(c *fiber.Ctx) error {
        queryParams := c.Queries()
        response, err := controller.GetData(queryParams)
        if err != nil {
            return fiber.NewError(fiber.StatusInternalServerError,err.Error())
        }
        return c.JSON(response)
    })

    app.Get("/insertData", func(c *fiber.Ctx) error {
        controller.InsertData()
        return c.SendString("Data Populated Successfully")
    })

    app.Get("/dropTables", func(c *fiber.Ctx) error {
        controller.DropTables()
        return c.SendString("Tables dropped Successfully")
    })
    
    app.Get("/createTables", func(c *fiber.Ctx) error {
        controller.CreateTables()
        return c.SendString("Tables created Successfully")
    })    

    app.Listen(":3000")
}