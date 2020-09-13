package main

import (
	"fmt"

	"github.com/cucumber/godog"
)

func iHaveEnteredIntoTheCalculator(n int) error {
	calc.enterNumber(n)
	return nil
}

func iPress(operation string) error {
	op := ParseOperator(operation)
	calc.press(op)
	return nil
}

func theResultShouldBeOnTheScreen(expectedResult int) error {
	if calc.result != expectedResult {
		return fmt.Errorf("expected %d, but result is %d", expectedResult, calc.result)
	}
	return nil
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() { calc = Calculator{} })
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
		calc = Calculator{} // clean the state before every scenario
	})

	ctx.Step(`^I have entered (\d+) into the calculator$`, iHaveEnteredIntoTheCalculator)
	ctx.Step(`^I press "([^"]*)"$`, iPress)
	ctx.Step(`^the result should be (\d+) on the screen$`, theResultShouldBeOnTheScreen)
}
