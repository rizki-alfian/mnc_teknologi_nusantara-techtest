package routes

import (
	"github.com/labstack/echo/v4"
	"mnc-users/apps/test1"
)

func Test1Routes(g *echo.Group, test1Handler *test1.Test1Handler) {
	test1 := g.Group("/test1")
	test1.GET("/find_match_string", test1Handler.FindMatchString)
	test1.GET("/calculate_change", test1Handler.CalculateChange)
	test1.GET("/is_valid_bracket_sequence", test1Handler.IsValidBracketSequence)
	test1.GET("/check_leave", test1Handler.CheckLeave)
}