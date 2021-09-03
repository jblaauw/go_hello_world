Feature: Get spots
    In order to book a spot
    As an employee
    I need to see which spots are available and those that are unavailable

    Scenario: Get all booking information from now until one week from now and there is atleast 1 booking
        Given User is verified as employee or admin
        And there is one booking
        When Employee visits bookings overview
        Then Show a list of all available spots for each day for upcoming week
