package main

import (
    "github.com/gofiber/fiber/v2"
    "fmt"
    "log"
//    "encoding/json"
)

func main() {
    app := fiber.New()

    // GET /api/register
    app.Get("/api/*", func(c *fiber.Ctx) error {
        msg := fmt.Sprintf("✋ %s", c.Params("*"))
        return c.SendString(msg) // => ✋ register
    })

    app.Post("/avg", func(c *fiber.Ctx) error {
    payload := struct {
        Data  []float32 `json:"data"`
    }{}
    if err := c.BodyParser(&payload); err != nil {
        return err
    }
    var avg float32 =0
    n:=0
    for _, v := range payload.Data {
        avg+=v
        n+=1
    }
    if n==0 {
        avg=0
    } else {
        avg/=float32(n)
    }
    return c.JSON(avg)
    })

    log.Fatal(app.Listen(":3000"))
}


