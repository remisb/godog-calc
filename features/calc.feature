Feature: calculator

    You should put a description of the feature under test here
    In this test we'll check our Calculator can add two numbers

    Scenario: Test calculator operations
        Given I have entered <number1> into the calculator
        And I have entered <number2> into the calculator
        When I press "<operation>"
        Then the result should be <result> on the screen

        Examples:
            | number1 | number2 | operation | result |
            | 50      | 80      | add       | 130    |
            | 2       | 3       | multiply  | 6      |
            | 45      | 14      | subtract  | 31     |
            | 45      | 15      | divide    | 3      |