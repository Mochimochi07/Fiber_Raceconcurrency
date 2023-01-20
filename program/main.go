package main

import (
    "fmt"
    "sync"

    "github.com/gofiber/fiber"
    "github.com/gofiber/fiber/middleware"
)

func main() {
    app := fiber.New()

    app.Use(middleware.Logger())

    app.Get("/download", func(c *fiber.Ctx) {
        var wg sync.WaitGroup
        games := []string{"Naruto", "Pokemon", "Megaman", "Dragonballz", "Digimon", "Bakugan", "Princemackaroo", "makibao", "slamdunk", "popeye"}

        for _, game := range games {
            wg.Add(1)
            go func(g string) {
                defer wg.Done()
                fmt.Printf("Downloading %s...\n", g)
               
                fmt.Printf("%s Downloaded\n", g)
            }(game)
        }

        wg.Wait()
        c.Send("All games downloaded")
    })

    app.Listen(3000)
}
