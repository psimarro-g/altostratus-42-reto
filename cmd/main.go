/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: psimarro <psimarro@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/10 17:54:51 by psimarro          #+#    #+#             */
/*   Updated: 2024/06/10 18:05:19 by psimarro         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"altostratus-42-reto/configs"
	"altostratus-42-reto/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	//run database
	configs.ConnectDB()

	//routes
	routes.AsteroidRoute(e)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, &echo.Map{"data": "Hello, World!"})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
